# moderation

[![godocs](https://godoc.org/github.com/schollz/progressbar?status.svg)](https://godoc.org/github.com/finnbear/moderation)

`moderation` is a profanity filter for `Go`.

## Goals

1. Easy to use
2. Minimum possible allocations, processing time, and binary size
3. Minimum false negatives (including text like `h3110_w0r!d`)
4. Minimum false positives
5. Provide a way to censor text
6. (Future) Other analysis types than inappropriate, profane, offensive, sexual, mean, spam (violence, contact info, etc.)
7. (Future) Basic support for languages other than English

## Example
```go
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
}

func printResult(phrase string) {
	description := "is appropriate"
	if moderation.IsInappropriate(phrase) {
		description = "is NOT appropriate"
	}
	fmt.Printf("\"%s\" %s.\n", phrase, description)
}

```

```console
$ go run hello_world.go
"hello world" is appropriate.
"$#1t" is NOT appropriate.
"a$$" is NOT appropriate.
"assassin" is appropriate.
```

## Comparison
Accuracy was evaluated based on the first 100,000 items from this [dataset of moderated comments](https://raw.githubusercontent.com/vzhou842/profanity-check/master/profanity_check/data/clean_data.csv).

|**Package**|**Time**|**Accuracy**|**Comment**|
|:-----:|:-----:|:-----:|:-----:|
|[finnbear/moderation](https://github.com/finnbear/moderation)|1.49s|90.76%|Current API version is not stable|
|[TwinProduction/go-away](https://github.com/TwinProduction/go-away)|2.23s|82.10%|Many false positives from combined words like "push it"|


## Acknowledgements

1. Radix implementation based on https://gitlab.com/caibear/go-boggle/
2. Some profanities and test cases are based on https://github.com/TwinProduction/go-away
