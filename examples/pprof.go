package main

import (
	"encoding/csv"
	"github.com/finnbear/moderation"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("moderation.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	wikiModerationData, err := os.Open("../wikipedia-test.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(wikiModerationData)

	for total := 0; total < 50000; total++ {
		fields, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		_ = moderation.IsInappropriate(fields[1])
	}

	err = wikiModerationData.Close()
	if err != nil {
		log.Fatal(err)
	}
}
