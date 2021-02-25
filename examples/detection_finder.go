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
	shorter := input
	for moderation.Is(shorter, 0xffffffff) { // satisfies all bitmasks
		input = shorter
		shorter = shorter[:len(shorter)-1]
	}

	shorter = input
	for moderation.IsInappropriate(shorter) {
		input = shorter
		shorter = shorter[1:]
	}

	if moderation.IsInappropriate(input) {
		fmt.Printf("Found inappropriate phrase: %s\n", input)
	} else {
		fmt.Println("No inappropriate phrase found")
	}
}
