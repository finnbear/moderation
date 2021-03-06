package main

import (
	"fmt"
	"github.com/finnbear/moderation"
)

func main() {
	printResult("hello world")
	printResult("$#1t")
	printResult("a$$")
	printResult("assassin")

	/*
		Expected results:
		"hello world" is appropriate.
		"$#1t" is NOT appropriate.
		"a$$" is NOT appropriate.
		"assassin" is appropriate.
	*/
}

func printResult(phrase string) {
	description := "is appropriate"
	if moderation.IsInappropriate(phrase) {
		description = "is NOT appropriate"
	}
	fmt.Printf("\"%s\" %s.\n", phrase, description)
}
