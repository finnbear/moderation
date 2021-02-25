// The package moderation implements a profanity filter.
package moderation

import (
	"github.com/finnbear/moderation/internal/radix"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

// Types of inappropriate (combinable with bitwise OR)
type Type uint32

const (
	Profane Type = 255 << (iota * 8)
	Offensive
	Sexual
	Mean
)

var (
	// The threshold where a phrase is considered inappropriate
	InappropriateThreshold int = 1

	tree *radix.Tree
)

func init() {
	tree = radix.New()
	for _, wv := range wordValues {
		tree.Add(wv.word, wv.value)
	}
}

// Replace the key with any one of the characters in the value
var replacements = [...]string{
	'!': "li",
	'@': "a",
	'4': "a",
	'8': "b",
	'6': "b",
	'(': "c",
	'<': "c",
	'3': "eg",
	'9': "gq",
	'#': "h",
	'1': "li",
	'0': "o",
	'5': "s",
	'$': "s",
	'+': "t",
	'7': "t",
	'2': "z",
}

var removeAccentsTransform = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

// IsInappropriate returns whether a phrase contains enough inappropriate words
// to meet or exceed InappropriateThreshold
//
// Equivalent to
//  moderation.Is(text, moderation.Profane|moderation.Offensive|moderation.Sexual)
//
func IsInappropriate(text string) bool {
	return Is(text, Profane|Offensive|Sexual)
}

// IsInappropriate returns whether a phrase contains enough words matching the
// types flag to meet or exceed InappropriateThreshold
func Is(text string, types Type) bool {
	// Sanitize input
	buf := make([]byte, 0, len(text))
	_, n, _ := transform.Append(removeAccentsTransform, buf, []byte(text))
	text = string(buf[:n])

	// How many characters ago the last separator character was observed,
	// expressed as an upper and a lower bound in relation to the current queue
	// of matches
	lastSepMin := 0
	lastSepMax := 0

	// Same as above, but for replacements
	lastReplacementMin := 0
	lastReplacementMax := 0

	// Scan status
	var matches radix.Queue
	inappropriateLevel := 0
	var lastMatchable byte

	for _, textRune := range text {
		// Unhandled runes (not printable, not representable as byte, etc.)
		if textRune < 0x0020 || textRune > 0x007E {
			continue
		}

		textByte := byte(textRune)
		var textBytes string
		lastSepMin++
		lastSepMax++
		lastReplacementMin++
		lastReplacementMax++

		matchable := false
		skippable := false

		var replacement string
		if int(textByte) < len(replacements) {
			replacement = replacements[textByte]
		}

		switch {
		case textByte >= 'A' && textByte <= 'Z':
			textByte += 'a' - 'A'
			fallthrough
		case textByte >= 'a' && textByte <= 'z':
			matchable = true
		case replacement != "":
			textByte = replacement[0]
			textBytes = replacement
			matchable = true
			lastReplacementMin = 0
			lastReplacementMax = 0
		default:
			switch textByte {
			case '*': // these count as replacements
				lastReplacementMin = 0
				lastReplacementMax = 0
				fallthrough
			case ' ', '~', '-', '_', '.', ',': // false positives may contain these
				skippable = true
				lastSepMin = 0
				lastSepMax = 0
			default:
				continue
			}
		}

		if matchable {
			if textByte == lastMatchable {
				// this character doesn't count so cancel the increments to min
				lastSepMin--
				lastReplacementMin--
			} else {
				// Add a new blank match to assume the new byte
				matches.Append(tree.Root())
			}

			//println(string([]byte{textByte}), textBytes)
			originalLength := matches.Len()
			for m := 0; m < originalLength; m++ {
				match := matches.Remove()

				// Technically should compare to previous byte of given match,
				// but this would be slower and give similar results for the
				// given replacements
				if skippable || textByte == lastMatchable {
					// Undo remove
					matches.AppendUnique(match)
					//println("=", match.Str, match.Depth())
				} else {
					//println("-", match.Str, match.Depth())
				}

				// Process textBytes as multiple textBytes or textByte
				loops := 1
				if len(textBytes) > 1 {
					loops = len(textBytes)
				}

				for l := 0; l < loops; l++ {
					loopTextByte := textByte
					if len(textBytes) > 0 {
						loopTextByte = textBytes[l]
					}
					next := match.Next(loopTextByte)

					if next == nil {
						continue
					}

					if next.Word() {
						if next.Depth() > 4 || (next.Depth() > 3 && next.Start() != 's') || (next.Depth() >= lastSepMin && next.Depth() <= lastSepMax) {
							match := next.Data() & uint32(types)
							for i := 0; i < 4; i++ {
								level := int(int8(match >> (i * 8)))

								// False positives that contain replacements are not matched
								if level > 0 || next.Depth()-1 <= lastReplacementMax {
									inappropriateLevel += level
								}
							}
						}
					}

					matches.Append(next)
					//println("+", next.Str, next.Depth(), " <- ", match.Depth())
				}
			}

			lastMatchable = textByte
		} else if !skippable {
			matches.Clear()
		}
	}

	return inappropriateLevel >= InappropriateThreshold
}
