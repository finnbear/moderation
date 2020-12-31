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
	for word, value := range wordValues {
		tree.Add(word, value)
	}
}

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
	buf := make([]byte, 0, len(text))
	_, n, _ := transform.Append(removeAccentsTransform, buf, []byte(text))
	text = string(buf[:n])
	lastSepMin := 0
	lastSepMax := 0

	var matches radix.Queue
	inappropriateLevel := 0

	var lastMatchable byte
	for _, textRune := range text {
		if textRune >= 0x0020 && textRune <= 0x007E {
			textByte := byte(textRune)
			var textBytes string
			lastSepMin++
			lastSepMax++

			ok := true
			matchable := false
			skippable := false

			var replacement string
			if int(textByte) < len(replacements) {
				replacement = replacements[textByte]
			}

			switch {
			case textByte >= 'a' && textByte <= 'z':
				matchable = true
			case textByte >= 'A' && textByte <= 'Z':
				textByte += 'a' - 'A'
				matchable = true
			case replacement != "":
				textByte = replacement[0]
				textBytes = replacement
				matchable = true
			default:
				switch textByte {
				case ' ', '_', '-', '.', ',', '*':
					skippable = true
					lastSepMin = 0
					lastSepMax = 0
				default:
					ok = false
				}
			}

			if textByte == lastMatchable {
				lastSepMin-- // this character doesn't count
			}

			if ok {
				if matchable {
					matches.Append(tree.Root())

					originalLength := matches.Len()
					for i := 0; i < originalLength; i++ {
						match := matches.Remove()

						if textByte == lastMatchable {
							matches.Append(match)
						}

						// Process textBytes as multiple textBytes or textByte
						loops := 1
						if len(textBytes) > 1 {
							loops = len(textBytes)
						}

						for i := 0; i < loops; i++ {
							loopTextByte := textByte
							if len(textBytes) > 0 {
								loopTextByte = textBytes[i]
							}
							next := match.Next(loopTextByte)

							if next != nil {
								if next.Word() {
									if next.Depth() > 4 || (next.Depth() > 3 && next.Start() != 's') || (next.Depth() >= lastSepMin && next.Depth() <= lastSepMax) {
										match := next.Data() & uint32(types)
										for i := 0; i < 4; i++ {
											inappropriateLevel += int(int8(match >> (i * 8)))
										}
									}
								}

								matches.Append(next)
							}
						}
					}

					lastMatchable = textByte
				} else if !skippable {
					matches.Clear()
				}
			}
		}
	}
	return inappropriateLevel >= InappropriateThreshold
}
