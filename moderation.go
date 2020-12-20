// The package moderation implements a profanity filter.
package moderation

import (
	"github.com/finnbear/moderation/internal/radix"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

var tree *radix.Tree

func init() {
	tree = radix.New()
	for _, profanity := range profanities {
		tree.Add(profanity, 1)
	}
	for _, falsePositive := range falsePositives {
		tree.Add(falsePositive, -1)
	}
}

var replacements = map[byte]string{
	'!': "li",
	'@': "a",
	'4': "a",
	'8': "b",
	'6': "b",
	'(': "c",
	'<': "c",
	'3': "e",
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

// Analyze analyzes a given phrase for moderation purposes
func Analyze(text string) (analysis Analysis) {
	text, _, _ = transform.String(removeAccentsTransform, text)
	lastSepMin := 0
	lastSepMax := 0

	matches := make([]*radix.Node, 0, len(text))
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

			replacement, replaceable := replacements[textByte]

			switch {
			case textByte >= 'a' && textByte <= 'z':
				matchable = true
			case textByte >= 'A' && textByte <= 'Z':
				textByte += 'a' - 'A'
				matchable = true
			case replaceable:
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

			if len(textBytes) < 1 {
				textBytes += string(textByte)
			}

			if textByte == lastMatchable {
				lastSepMin-- // this character doesn't count
			}

			if ok {
				matches = append(matches, tree.Root())
				if matchable {
					for matchIndex, match := range matches {
						if match == nil {
							continue
						}

						for i := 0; i < len(textBytes); i++ {
							textByte := textBytes[i]
							next := match.Next(textByte)

							if next != nil {
								if next.Word() {
									if next.Data() == 0 {
										// clear
									} else if next.Depth() > 4 || (next.Depth() > 3 && next.Start() != 's') || (next.Depth() >= lastSepMin && next.Depth() <= lastSepMax) {
										analysis.InappropriateLevel += int(next.Data())
									}
								}
							}

							if textByte == lastMatchable || i > 1 {
								if next != nil {
									matches = append(matches, next)
								}
							} else {
								matches[matchIndex] = next
							}
						}
					}

					lastMatchable = textByte
				} else if !skippable {
					matches = matches[:0]
				}
			}
		}
	}
	return
}
