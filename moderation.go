// The package moderation implements a profanity filter.
package moderation

import (
	"github.com/finnbear/moderation/internal/radix"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

// Types and severities of inappropriateness
//
// For compability, always reference them by name as their value may change
// from version to version.
//
// Use a bitwise OR of multiple profanity classifications, and a bitwise AND to
// specify a severity level (default Mild). The definition of Inappropriate
// (mildly profane, mildly offensive, mildly sexual, or severely mean) serves
// as a good example.
//
// Other operations on Type's are NOT supported.
//
// Severities sould be interpreteted on an "at least" basis, e.g. Mild means
// Mild, Moderate, OR Severe.
type Type uint32

const (
	Profane Type = 0b111 << (iota * 3)
	Offensive
	Sexual
	Mean
	Spam
	Inappropriate = Profane | Offensive | Sexual | (Mean & Severe)
	Any           = Profane | Offensive | Sexual | Spam | Mean

	Mild     Type = 0b111_111_111_111_111
	Moderate      = 0b110_110_110_110_110
	Severe        = 0b100_100_100_100_100

	countableTypes = 4

	// A subset of the ASCII range that requires no sanitization
	minNormal rune = 0x0020
	maxNormal rune = 0x007E
)

var (
	tree radix.Tree = radix.New()

	removeAccentsTransform = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
)

func init() {
	for _, wv := range wordValues {
		tree.Add(wv.word, wv.value)
	}
}

// IsInappropriate returns whether a phrase contains enough inappropriate words
// to meet or exceed InappropriateThreshold
//
// Equivalent to
//  moderation.Is(text, moderation.Inappropriate)
//
// Also, for the time being, equivalent to
//  moderation.Is(text, moderation.Profane|moderation.Offensive|moderation.Sexual)
//
func IsInappropriate(text string) bool {
	return Is(text, Inappropriate)
}

// Is returns whether a phrase contains words matching the types flag, useful if
// checking only one type or set of types is needed
func Is(text string, types Type) bool {
	return Scan(text)&types != 0
}

// Scan returns a bitmask of all types, useful if checking multiple types or
// sets of types is needed, without multiple calls to Is(text, types)
func Scan(text string) (types Type) {
	// Figure out if sanitization is needed, and if so, do it
	for _, textRune := range text {
		if textRune < minNormal || maxNormal < textRune {
			// Sanitize
			buf := make([]byte, 0, len(text))
			_, n, _ := transform.Append(removeAccentsTransform, buf, []byte(text))
			text = string(buf[:n])

			// Done sanitizing, stop scanning
			break
		}
	}

	// Scan status
	var matches radix.Queue
	var countableTypeLevels [countableTypes]int
	separate := true // whether the previous character was a separator
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
		} else if textRune > maxNormal {
			replacement = runeReplacements[textRune]
			if replacement == "" {
				lowerRune := unicode.ToLower(textRune)
				replacement = runeReplacements[lowerRune]
			}
		}

		switch {
		case textByte >= 'a' && textByte <= 'z': // most likely case
			matchable = true
		case textByte >= 'A' && textByte <= 'Z':
			upperCount++
			textByte += 'a' - 'A'
			matchable = true
		case replacement != "": // if there is a valid set of replacements
			textByte = replacement[0]
			textBytes = replacement
			matchable = true
		default:
			// matchable = false implied
			switch textByte {
			case '*': // these count as replacements
				replaced = true
				fallthrough
			case ' ', '~', '-', '_', '.', ',', '\n', '\r', '\t': // false positives may contain these
				skippable = true
			}
		}

		if matchable {
			if textByte == lastMatchable {
				repetitionCount++
			}

			// Add a new blank match to assume the new byte(s)
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
							for i := 0; i < countableTypes; i++ {
								level := int(int8(data >> (i * 8)))

								// False positives that contain replacements are not matched
								if level > 0 || !(match.Replaced || replaced) {
									countableTypeLevels[i] += level
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

	for i, level := range countableTypeLevels {
		var severity Type

		if level >= 3 {
			severity = 0b100 // severe
		} else if level == 2 {
			severity = 0b010 // moderate
		} else if level == 1 {
			severity = 0b001 // mild
		}

		types |= severity << (i * 3)
	}

	// Min length is arbitrary, but must be > 0 to avoid dividing by zero
	if len(text) > 5 {
		spamPercent := (100 / 2) * (upperCount + repetitionCount) / len(text)

		// TODO: Define severe spam

		if spamPercent > 50 {
			types |= Spam & Moderate
		} else if spamPercent > 30 {
			types |= Spam & Mild
		}
	}

	return
}
