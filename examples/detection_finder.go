package main

import (
	"flag"
	"fmt"
	"github.com/finnbear/moderation"
)

var input string

func init() {
	flag.StringVar(&input, "input", "", "the input to find an inappropriate phase in")
	flag.Parse()
}

func main() {
	fmt.Println("Original phrase: " + input)
	fmt.Printf("Inappropriate: %t\n", moderation.IsInappropriate(input))
	censored, numCensored := moderation.Censor(input, moderation.Inappropriate)
	fmt.Printf("Censored phrase: %s (%d characters replaced)\n", censored, numCensored)

	shorter := input
	for moderation.Scan(shorter).Is(moderation.Any) {
		input = shorter
		shorter = shorter[:len(shorter)-1]
	}

	shorter = input
	for moderation.Scan(shorter).Is(moderation.Any) {
		input = shorter
		shorter = shorter[1:]
	}

	if moderation.Scan(input).Is(moderation.Any) {
		fmt.Printf("Found inappropriate phrase: %s\n", input)
	} else {
		fmt.Println("No inappropriate phrase found")
	}
}
