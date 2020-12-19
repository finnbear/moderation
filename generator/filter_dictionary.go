package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/schollz/progressbar"
)

var (
	dictionaryFile         string
	profanityFile          string
	blacklistFile          string
	falsePositiveFile      string
	filteredDictionaryFile string
	goFilename             string

	plural = pluralize.NewClient()
)

const goTemplateSrc = `package moderation

// Generated by generator/filter_dictionary.go; DO NOT EDIT

var profanities = {{.Profanities}}

var falsePositives = {{.FalsePositives}}
`

func init() {
	if len(os.Args) != 7 {
		log.Fatalf("expected 7 args, got %d", len(os.Args))
	}
	dictionaryFile = os.Args[1]
	profanityFile = os.Args[2]
	blacklistFile = os.Args[3]
	falsePositiveFile = os.Args[4]
	filteredDictionaryFile = os.Args[5]
	goFilename = os.Args[6]
}

func fileToStrings(filename string) (lines []string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	str := string(buf)
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)
	return
}

func main() {
	dictionary := fileToStrings(dictionaryFile)
	profanities := fileToStrings(profanityFile)
	blacklistRaw := fileToStrings(blacklistFile)
	falsePositives := fileToStrings(falsePositiveFile)

	for _, falsePositive := range falsePositives {
		dictionary = append(dictionary, falsePositive)
	}

	var blacklistRegexes []*regexp.Regexp
	for _, raw := range blacklistRaw {
		r := regexp.MustCompile(raw)
		blacklistRegexes = append(blacklistRegexes, r)
	}

	filtered := make(map[string]int)

	bar := progressbar.New(len(dictionary))

	for _, word := range dictionary {
		bar.Add(1)

		blacklist := false
		wordSingular := plural.Singular(word)

		count := 0

		for _, profanity := range profanities {
			//profanitySegments := strings.Split(profanityLine, ",")
			//profanity := profanitySegments[0]
			//_, _ = strconv.Atoi(profanitySegments[1])
			if word == profanity || wordSingular == profanity {
				blacklist = true
				break
			}
			if strings.Index(word, profanity) != -1 {
				count++
				break
			}
		}
		if count > 0 && !blacklist {
			for _, r := range blacklistRegexes {
				if RegexMatchAll(r, word) {
					blacklist = true
					break
				}
			}

			if !blacklist {
				filtered[word] = count
			}
		}
	}

	var buffer bytes.Buffer

	var remove []string
	for word1, count1 := range filtered {
		for word2, count2 := range filtered {
			if count1 == count2 && len(word2) > len(word1) && strings.Index(word2, word1) != -1 {
				remove = append(remove, word2)
			}
		}
	}

	for _, r := range remove {
		delete(filtered, r)
	}

	goTemplate, err := template.New("go").Parse(goTemplateSrc)
	if err != nil {
		panic(err)
	}

	var list []string

	for word, count := range filtered {
		if count > 1 {
			// If this happens, may need to store how many profanities are
			// in each false positive
			println("warning: " + word + " contains multiple profanities")
		}
		list = append(list, word)
	}

	sort.Strings(list)

	/*
	   clean := fileToStrings("clean.txt")
	   for _, c := range clean {
	       found := false
	       for word := range filtered {
	           if word == c {
	               found = true
	               break
	           }
	       }
	       if !found {
	           println("missed", c)
	       }
	   }
	*/

	for _, word := range list {
		buffer.WriteString(word)
		buffer.WriteByte('\n')
	}

	err = ioutil.WriteFile(filteredDictionaryFile, buffer.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	goFile, err := os.Create(goFilename)
	if err != nil {
		log.Fatal(err)
	}
	err = goTemplate.Execute(goFile, map[string]interface{}{
		"Profanities":    fmt.Sprintf("%#v", profanities),
		"FalsePositives": fmt.Sprintf("%#v", list),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = goFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	println()
}

func RegexMatchAll(pattern *regexp.Regexp, str string) bool {
	pattern.Longest()
	return len(pattern.FindString(str)) == len(str)
}
