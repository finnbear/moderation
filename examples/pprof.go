package main

import (
	"github.com/finnbear/moderation"
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

	for i := 0; i < 100000; i++ {
		moderation.Analyze("hello")
		moderation.Analyze("sh1t")
		moderation.Analyze("Hello John Doe, I hope you're feeling well, as I come today bearing shitty news regarding your favorite chocolate chip cookie brand")
	}
}
