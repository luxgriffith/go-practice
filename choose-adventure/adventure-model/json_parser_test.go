package adventure_model

import (
	"testing"
)

var ExpectedStory *Story

func setupTestStory() {
	ExpectedStory = &Story{
		arcs: map[string]*Arc{
			"intro": {
				title: "Intro",
				text: []string{
					"Intro to story",
					"a paragraph",
					"another paragraph",
					"weeeeeeee"},
				options: []*Option{
					{
						text:     "Option A",
						arcTitle: "A",
					},
					{
						text:     "Option B",
						arcTitle: "B",
					},
					{
						text:     "Option C",
						arcTitle: "C",
					},
				},
			},
			"A": {
				title: "A",
				text: []string{
					"Option A's Story",
					"a paragraph",
					"another paragraph",
					"weeeeeeee",
					"your stuck teehee",
				},
				options: []*Option{
					{
						text:     "Option B",
						arcTitle: "B",
					},
				},
			},
			"B": {
				title: "B",
				text: []string{
					"Option B's Story",
					"a paragraph",
					"another paragraph",
					"weeeeeeee",
					"jail time",
				},
				options: []*Option{
					{
						text:     "Option A",
						arcTitle: "A",
					},
				},
			},
			"C": {
				title: "C",
				text:  []string{},
				options: []*Option{
					{
						text:     "Option D",
						arcTitle: "D",
					},
				},
			},
			"D": {
				title: "D",
				text: []string{
					"YOU DIED",
					"insert dark souls noise",
				},
				options: []*Option{},
			},
		},
	}
}

func TestBadPath(t *testing.T) {
	_, err := ReadStory("lolololol")
	if err == nil {
		t.Fail()
		t.Fatalf("TestBadPath Failed, err was nil")
	}
	t.Logf("TestBadPath triggered error: %v", err.Error())
}

func TestBadFormat(t *testing.T) {
	_, err := ReadStory("../badformatstory.json")
	if err == nil {
		t.Fail()
		t.Fatalf("TestBadPath Failed, err was nil")
	}
	t.Logf("TestBadFormat triggered error: %v", err.Error())
}

func TestReadableStory(t *testing.T) {
	resultStory, err := ReadStory("../story.json")
	setupTestStory()
	if err != nil {
		t.Fail()
		t.Fatalf("TestReadableStory failed: %v", err.Error())
	}
	if resultStory.toString() != ExpectedStory.toString() {
		t.Fail()
		t.Fatalf("TestReadableStory Failed, story mismatch. Expected %v got %v", ExpectedStory.toString(), resultStory.toString())
	}
	t.Logf("TestReadableStory Passed")
}
