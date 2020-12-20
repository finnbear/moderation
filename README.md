# moderation

[![godocs](https://godoc.org/github.com/schollz/progressbar?status.svg)](https://godoc.org/github.com/finnbear/moderation) 

`moderation` is a profanity filter for `Go`.

## Goals

1. Easy to use
2. Minimum possible allocations, processing time, and binary size
3. Minimum false negatives (including text like `h3110_w0r!d`)
4. Minimum false positives
5. (Future) Other analysis types than inappropriate (spam, violence, contact info, etc.)
6. (Future) Implement a way to censor text
7. (Future) basic support for languages other than English

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
    if isInappropriate(phrase) {
        description = "is NOT appropriate"
    }
    fmt.Printf("\"%s\" %s.\n", phrase, description)
}

func isInappropriate(phrase string) bool {
    return moderation.Analyze(phrase).IsInappropriate()
}
```

Output:
```console
"hello world" is appropriate.
"$#1t" is NOT appropriate.
"a$$" is NOT appropriate.
"assassin" is appropriate.
```

## Acknowledgements

1. Radix implementation based on https://gitlab.com/caibear/go-boggle/
2. Some profanities and test cases are based on https://github.com/TwinProduction/go-away
