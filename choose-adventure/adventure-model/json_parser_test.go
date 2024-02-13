package adventure_model

import (
	"maps"
	"testing"
)

var ExpectedStory *Story

func setupTestStory() {
	ExpectedStory = &Story{
		arcs: map[string]*Arc{
			"intro": {
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
				text: []string{},
				options: []*Option{
					{
						text:     "Option D",
						arcTitle: "D",
					},
				},
			},
			"D": {
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
	_, err := readStory("lolololol")
	if err == nil {
		t.Fail()
		t.Fatalf("TestBadPath Failed, err was nil")
	}
	t.Logf("TestBadPath triggered error: %v", err.Error())
}

func TestBadFormat(t *testing.T) {
	_, err := readStory("../badformatstory.json")
	if err == nil {
		t.Fail()
		t.Fatalf("TestBadPath Failed, err was nil")
	}
	t.Logf("TestBadFormat triggered error: %v", err.Error())
}

func TestReadableStory(t *testing.T) {
	resultStory, err := readStory("../story.json")
	setupTestStory()
	if err != nil {
		t.Fail()
		t.Fatalf("TestReadableStory failed: %v", err.Error())
	}
	if !maps.Equal(resultStory.arcs, ExpectedStory.arcs) {
		t.Fail()
		t.Fatalf("TestReadableStory failed: Expected %v got %v", ExpectedStory.arcs, resultStory.arcs)
	}
	t.Logf("TestReadableStory Passed")
}
