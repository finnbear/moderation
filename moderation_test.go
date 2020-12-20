package moderation

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func TestAnalyze(t *testing.T) {
	type TestCase struct {
		phrase    string
		offensive bool
	}
	testCases := []TestCase{
		{"hi", false},
		{"hello", false},
		{"hello my name is Bob.", false},
		{"SHIT", true},
		{"shhhhhiiiiter", true},
		{"shhHhhit", true},
		{"lol fuck this", true},
		{"f*u*c*k", true},
		{"$#1t", true},
		{" fučk", true},
		{"ass", true},
		{"glass", false},
		{"ÄšŚ", true},
		{"sex", true},
		{"hello_world-sex_word", true},
		{"sexy", true},
		{"is extreme", false},
		{"pÓöp", true},
		{"what a bunch of bullsh1t", true},
		{"bitčh", true},
		{"assassin", false},
		{"carcass", false},
		{"I had called upon my friend, Mr. Sherlock Holmes, one day in the autumn of last year and found him in deep conversation with a very stout, florid-faced, elderly gentleman with fiery red hair.", false},
		{"With an apology for my intrusion, I was about to withdraw when Holmes pulled me abruptly into the room and closed the door behind me.", false},
		{"You could not possibly have come at a better time, my dear Watson, he said cordially", false},
		{"I was afraid that you were engaged.", false},
		{"So I am. Very much so.", false},
		{"Then I can wait in the next room.", false},
		{"Not at all. This gentleman, Mr. Wilson, has been my partner and helper in many of my most successful cases, and I have no doubt that he will be of the utmost use to me in yours also.", false},
		{"The stout gentleman half rose from his chair and gave a bob of greeting, with a quick little questioning glance from his small fat-encircled eyes", false},
		{"Try the settee, said Holmes, relapsing into his armchair and putting his fingertips together, as was his custom when in judicial moods.", false},
		{"I know, my dear Watson, that you share my love of all that is bizarre and outside the conventions and humdrum routine of everyday life.", false},
		{"You have shown your relish for it by the enthusiasm which has prompted you to chronicle, and, if you will excuse my saying so, somewhat to embellish so many of my own little adventures.", false},
		{"You did, Doctor, but none the less you must come round to my view, for otherwise I shall keep on piling fact upon fact on you until your reason breaks down under them and acknowledges me to be right.", false},
		{"Now, Mr. Jabez Wilson here has been good enough to call upon me this morning, and to begin a narrative which promises to be one of the most singular which I have listened to for some time.", false},
		{"You have heard me remark that the strangest and most unique things are very often connected not with the larger but with the smaller crimes, and occasionally", false},
		{"indeed, where there is room for doubt whether any positive crime has been committed.", false},
		{"As far as I have heard it is impossible for me to say whether the present case is an instance of crime or not, but the course of events is certainly among the most singular that I have ever listened to.", false},
		{"Perhaps, Mr. Wilson, you would have the great kindness to recommence your narrative.", false},
		{"I ask you not merely because my friend Dr. Watson has not heard the opening part but also because the peculiar nature of the story makes me anxious to have every possible detail from your lips.", false},
		{"As a rule, when I have heard some slight indication of the course of events, I am able to guide myself by the thousands of other similar cases which occur to my memory.", false},
		{"In the present instance I am forced to admit that the facts are, to the best of my belief, unique.", false},
		{"We had reached the same crowded thoroughfare in which we had found ourselves in the morning.", false},
		{"Our cabs were dismissed, and, following the guidance of Mr. Merryweather, we passed down a narrow passage and through a side door, which he opened for us", false},
		{"Within there was a small corridor, which ended in a very massive iron gate.", false},
		{"We were seated at breakfast one morning, my wife and I, when the maid brought in a telegram. It was from Sherlock Holmes and ran in this way", false},
	}
	for _, testCase := range testCases {
		analysis := Analyze(testCase.phrase)
		if analysis.IsInappropriate() != testCase.offensive {
			t.Errorf("phrase=\"%s\" analysis offensive=%v actual offensive=%v", testCase.phrase, analysis.IsInappropriate(), testCase.offensive)
		}
	}
}

func TestAnalyzeWikipedia(t *testing.T) {
	wikiModerationData, err := os.Open("wikipedia-test.csv")
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
		analysis := Analyze(phrase)
		if analysis.IsInappropriate() == offensive {
			correct++
		} else {
			//t.Errorf("phrase=\"%s\" analysis offensive=%v actual offensive=%v", phrase, analysis.IsOffensive(), offensive)
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
