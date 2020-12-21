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

var replacements = [...]string{
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
	buf := make([]byte, 0, len(text))
	_, n, _ := transform.Append(removeAccentsTransform, buf, []byte(text))
	text = string(buf[:n])
	lastSepMin := 0
	lastSepMax := 0

	var matchesGet, matchesPut radix.Buffer

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
					matchesGet.Append(tree.Root())

					for i := 0; i < matchesGet.Len(); i++ {
						match := matchesGet.Get(i)

						if textByte == lastMatchable {
							matchesPut.Append(match)
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
										analysis.InappropriateLevel += int(next.Data())
									}
								}

								matchesPut.Append(next)
							}
						}
					}

					lastMatchable = textByte
					matchesGet = matchesPut
					matchesPut.Clear()
				} else if !skippable {
					matchesGet.Clear()
				}
			}
		}
	}
	return
}
