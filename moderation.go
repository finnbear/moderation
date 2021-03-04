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
	Profane Type = 1 << iota
	Offensive
	Sexual
	Mean
	Spam
	Inappropriate = Profane | Offensive | Sexual
	Any           = Profane | Offensive | Sexual | Spam | Mean
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

const (
	minMatchable rune = 0x0020
	maxMatchable rune = 0x007E
)

// IsInappropriate returns whether a phrase contains enough inappropriate words
// to meet or exceed InappropriateThreshold
//
// Equivalent to
//  moderation.Is(text, moderation.Profane|moderation.Offensive|moderation.Sexual)
//
func IsInappropriate(text string) bool {
	return Is(text, Profane|Offensive|Sexual)
}

// Is returns whether a phrase contains enough words matching the
// types flag to meet or exceed InappropriateThreshold
func Is(text string, types Type) bool {
	// Sanitize input, if needed
	needsSanitize := false
	for _, textRune := range text {
		if textRune < minMatchable || maxMatchable < textRune {
			needsSanitize = true
			break
		}
	}
	if needsSanitize {
		buf := make([]byte, 0, len(text))
		_, n, _ := transform.Append(removeAccentsTransform, buf, []byte(text))
		text = string(buf[:n])
	}

	separate := true

	// Scan status
	var matches radix.Queue
	inappropriateLevel := 0
	var lastMatchable byte

	// For spam detection purposes
	var upperCount int
	var repetitionCount int

	for _, textRune := range text {
		textByte := byte(textRune)
		var textBytes string

		matchable := false
		replaced := false
		skippable := false

		var replacement string
		if int(textByte) < len(replacements) {
			replacement = replacements[textByte]
		}

		switch {
		case textRune < minMatchable || maxMatchable < textRune:
			// Unhandled runes (not printable, not representable as byte, etc.)
			// matchable = false
			switch textRune {
			case '\n', '\r', '\t':
				skippable = true
			}
		case textByte >= 'A' && textByte <= 'Z':
			upperCount++
			textByte += 'a' - 'A'
			fallthrough
		case textByte >= 'a' && textByte <= 'z':
			matchable = true
		case replacement != "":
			textByte = replacement[0]
			textBytes = replacement
			matchable = true
		default:
			switch textByte {
			case '*': // these count as replacements
				replaced = true
				fallthrough
			case ' ', '~', '-', '_', '.', ',': // false positives may contain these
				skippable = true
			}
		}

		if matchable {
			if textByte == lastMatchable {
				repetitionCount++
			}

			// Add a new blank match to assume the new byte
			//println(string([]byte{textByte}), "\t", separate)
			matches.AppendUnique(radix.Match{Node: tree.Root(), Length: 0, Replaced: false, Separate: separate})
			//println("+", "root", separate, replaced)
			originalLength := matches.Len()
			for m := 0; m < originalLength; m++ {
				match := matches.Remove()

				// Technically should compare to previous byte of given match,
				// but this would be slower and give similar results for the
				// given replacements
				if (skippable || textByte == lastMatchable) && match.Length > 0 {
					// Undo remove (and add one to length)
					matches.AppendUnique(radix.Match{Node: match.Node, Length: match.Length + 1, Replaced: replaced || match.Replaced, Separate: match.Separate})
					//println("=", match.Node.Depth(), match.Separate, match.Replaced)
				} else {
					//println("-", match.Node.Depth(), match.Separate, match.Replaced)
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

					next := match.Node.Next(loopTextByte)

					if next == nil {
						continue
					}

					if next.Word() {
						if next.Depth() > 4 || (next.Depth() > 3 && next.Start() != 's') || match.Separate {
							data := next.Data()
							for i := 0; i < 4; i++ {
								if types&Type(1<<i) == 0 {
									continue
								}

								level := int(int8(data >> (i * 8)))

								// False positives that contain replacements are not matched
								if level > 0 || !(match.Replaced || replaced) {
									inappropriateLevel += level
								}
							}
						}
					}

					matches.Append(radix.Match{Node: next, Length: match.Length + 1, Replaced: replaced || match.Replaced, Separate: match.Separate})
					//println("+", next.Depth(), match.Separate, match.Replaced)
				}
			}

			lastMatchable = textByte
		} else if !skippable {
			matches.Clear()
		}

		separate = skippable || !matchable
	}

	if types&Spam != 0 && len(text) > 5 {
		spamPercent := (100 / 2) * (upperCount + repetitionCount) / len(text)
		inappropriateLevel += spamPercent / 30
	}

	return inappropriateLevel >= InappropriateThreshold
}
