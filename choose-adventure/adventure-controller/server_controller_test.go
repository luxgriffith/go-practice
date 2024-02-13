package adventure_controller

import (
	"testing"

	adventure_model "github.com/griffithscg/go-practice/choose-adventure/adventure-model"
)

var expectedStory *adventure_model.Story

func setupTestStory() {
}

func TestGetNextArc(t *testing.T) {
	setupTestStory()
	expectedTitle := "A"
	expectedStory, _ = adventure_model.ReadStory("../story.json")
	expectedArc := expectedStory.GetArcs()[expectedTitle]
	inputOption := adventure_model.NewOption("Option A", "A")
	inputStory, _ := adventure_model.ReadStory("../story.json")
	resultTitle, resultArc, err := getNextArc(inputOption, inputStory)
	if err != nil {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed: Error triggered, %v", err.Error())
	}
	if expectedTitle != resultTitle {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed, title mismatch: Expected %v, got %v", expectedTitle, resultTitle)
	}
	if expectedArc.ToString() != resultArc.ToString() {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed, arc mismatch: expected %v, got %v", expectedArc.ToString(), resultArc.ToString())
	}
	t.Logf("TestGetNextArc Passed")
}

func TestBadInputs(t *testing.T) {
	_, _, err := getNextArc(nil, nil)
	if err == nil {
		t.Fail()
		t.Fatalf("TestBadInputs Failed, no error returned")
	}
	t.Logf("TestBadInputs caused error %v", err.Error())
}
