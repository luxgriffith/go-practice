package adventure_model

import (
	"testing"
)

func TestGetNextArc(t *testing.T) {
	setupTestStory()
	expectedTitle := "A"
	expectedArc := ExpectedStory.arcs[expectedTitle]
	inputOption := &Option{
		text:     "Option A",
		arcTitle: "A",
	}
	inputStory, _ := readStory("../story.json")
	resultTitle, resultArc, err := getNextArc(inputOption, inputStory)
	if err != nil {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed: Error triggered, %v", err.Error())
	}
	if expectedTitle != resultTitle {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed, title mismatch: Expected %v, got %v", expectedTitle, resultTitle)
	}
	if expectedArc.toString() != resultArc.toString() {
		t.Fail()
		t.Fatalf("TestGetNextArc Failed, arc mismatch: expected %v, got %v", expectedArc.toString(), resultArc.toString())
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
