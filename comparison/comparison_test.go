package moderation

import (
	"encoding/csv"
	"io"
	"os"
	"testing"

    "github.com/finnbear/moderation"
    "github.com/TwinProduction/go-away"
)

func moderationIsInappropriate(phrase string) bool {
	return moderation.Analyze(phrase).IsInappropriate()
}

func TestModerationWikipedia(t *testing.T) {
    testWikipedia(t, moderationIsInappropriate)
}

func TestGoAwayWikipedia(t *testing.T) {
    testWikipedia(t, goaway.IsProfane)
}

func testWikipedia(t *testing.T, isInappropriate func(string) bool) {
	wikiModerationData, err := os.Open("../wikipedia-test.csv")
	if err != nil {
		t.Skip()
	}
	reader := csv.NewReader(wikiModerationData)

	correct := 0
	total := 0

	for total < 50000 {
		fields, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Error(err)
		}
		phrase := fields[1]
		offensive := fields[0] == "1"
		inappropriate := isInappropriate(phrase)
		if inappropriate == offensive {
			correct++
		} else {
			//fmt.Printf("phrase=\"%s\" analysis offensive=%v actual offensive=%v", phrase, inappropriate, offensive)
		}

		total++
	}

	accuracy := 100 * float64(correct) / float64(total)

	// Wikipedia takes into account more than whether the text contains
	// bad words
	const requiredAccuracy = 90

	if accuracy >= requiredAccuracy {
		t.Logf("accuracy was %f%% (%d%% required)\n", accuracy, requiredAccuracy)
	} else {
		t.Errorf("accuracy was %f%% (%d%% required)\n", accuracy, requiredAccuracy)
	}

	err = wikiModerationData.Close()
	if err != nil {
		t.Error(err)
	}
}
