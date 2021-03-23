package moderation

import "golang.org/x/exp/utf8string"

var CensorReplacment rune = '*'

// Censor returns a string with all but the first character of any inappropriate
// segment replaced with CensorReplacment
//
// It is currently Experimental and not fully tested
func Censor(text string, types Type) (censoredText string, replaced int) {
	// Fast path
	if len(text) == 0 || !Scan(text).Is(types) {
		return text, 0
	}

	str := utf8string.NewString(text)

	censored := make([]rune, 0, str.RuneCount())

	start := 0

	// TODO: scan ahead for false positives

	for i := start; i <= str.RuneCount(); i++ {
		slice := str.Slice(start, i)
		if /* (i == str.RuneCount() || str.At(i) == ' ') && */ Scan(slice).Is(types) {
			for j := start; j <= i; j++ {
				slice2 := str.Slice(j, i)
				if !Scan(slice2).Is(types) {
					censored = append(censored, []rune(str.Slice(start, j))...)
					replaced += i - j
					for k := 0; k < i-j; k++ {
						censored = append(censored, CensorReplacment)
					}
					break
				}
			}
			start = i
		}
	}
	censored = append(censored, []rune(str.Slice(start, str.RuneCount()))...)
	censoredText = string(censored)
	return
}
