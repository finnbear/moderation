package moderation

var (
	// Replace the key with any one of the characters in the value
	replacements = [...]string{
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

	runeReplacements = map[rune]string{
		// Greek letters
		'α': "a",
		'β': "b",
		'γ': "y",
		'∆': "a",
		'δ': "d",
		'ε': "e",
		'ζ': "z",
		'η': "hn",
		'θ': "o",
		'ι': "i",
		'κ': "k",
		'λ': "l",
		'μ': "mu",
		'ν': "nv",
		'ο': "o",
		'ρ': "p",
		'ς': "s",
		'τ': "t",
		'υ': "u",
		'φ': "p",
		'χ': "x",
		'ψ': "t",
		'Ω': "o",
		'ω': "w",

		// Math symbols
		'⊗': "o",
		'⊕': "o",
		'σ': "o",
		'∩': "n",
		'∪': "u",
		'⊂': "c",
		'⊆': "c",
		'⊄': "c",
		'∈': "e",
		'⊖': "o",
		'Ø': "o",
		'∨': "v",
		'∄': "ab",
		'∫': "l",

		// Letterlike
		'ℂ': "c",
		'℃': "c",
		'℄': "c",
		'ℇ': "e",
		'℉': "f",
		'ℊ': "g",
		'ℋ': "h",
		'ℌ': "h",
		'ℍ': "h",
		'ℎ': "h",
		'ℏ': "h",
		'ℐ': "j",
		'ℑ': "j",
		'ℒ': "l",
		'ℓ': "l",
		'℔': "b",
		'ℕ': "n",
		'№': "n",
		'℗': "p",
		'℘': "p",
		'ℙ': "p",
		'ℚ': "q",
		'ℛ': "r",
		'ℜ': "r",
		'ℝ': "r",
		'℟': "r",
		'℣': "v",
		'ℤ': "z",
		'℧': "o",
		'℩': "i",
		'K': "k",
		'Å': "a",
		'ℬ': "b",
		'ℭ': "c",
		'℮': "e",
		'e': "e",
		'ℰ': "e",
		'ℱ': "f",
		'ℳ': "m",
		'ℴ': "o",
		'ℵ': "n",
		'ℹ': "i",
		'℺': "o",
		'ℼ': "n",
		'ℽ': "v",
		'ℿ': "n",
		'⅀': "e",
		'⅁': "g",
		'⅄': "l",
		'ⅅ': "d",
		'ⅆ': "d",
		'ⅇ': "e",
		'ⅈ': "i",
		'ⅉ': "ji",
	}
)
