package adventure_model

import (
	"testing"
)

var expectedOptionString string
var expectedArcString string
var expectedStoryString string
var expectedArcTitleList []string
var testStory *Story
var testArc *Arc
var testOption *Option

func setupTestVars() {
	testStory = &Story{
		arcs: map[string]*Arc{
			"intro": {
				text: []string{"intro"},
				options: []*Option{
					{
						text:     "OptionFoo",
						arcTitle: "Foo",
					},
					{
						text:     "OptionBarr",
						arcTitle: "Bar",
					},
				},
			},
			"Foo": {
				text: []string{"foo, foo2"},
				options: []*Option{
					{
						text:     "OptionBar",
						arcTitle: "Bar",
					},
				},
			},
			"Bar": {
				text:    nil,
				options: nil,
			},
		},
	}

	testArc = &Arc{
		text: []string{"foo, foo2"},
		options: []*Option{
			{
				text:     "OptionBar",
				arcTitle: "Bar",
			},
		},
	}
	testOption = &Option{
		text:     "OptionFoo",
		arcTitle: "Foo",
	}
	expectedArcString = "foo, foo2\tOptionBar, Bar"
	expectedOptionString = "OptionFoo, Foo"
	expectedArcTitleList = []string{"intro", "Foo", "Bar"}
	expectedStoryString = "intro:intro\tOptionFoo, Foo; OptionBar, Bar\nFoo:foo, foo2\tOptionBar, Bar\nBar:\t"
}

func TestOptionToString(t *testing.T) {
	setupTestVars()
	resultOptionString := testOption.toString()
	if resultOptionString != expectedOptionString {
		t.Fail()
		t.Fatalf("testOptionToString Failed, expected %v got %v", expectedOptionString, resultOptionString)
	}
	t.Logf("TestOptionToString passed")
}

func TestArcToString(t *testing.T) {
	setupTestVars()
	resultArcString := testArc.toString()
	if resultArcString != expectedArcString {
		t.Fail()
		t.Fatalf("TestArcString Failed, expected %v got %v", expectedArcString, resultArcString)
	}
	t.Logf("TestArcString passed")
}

func TestGetArcTitles(t *testing.T) {
	setupTestVars()
	resultArcTitleList := testStory.getArcTitles()
	if len(resultArcTitleList) != len(expectedArcTitleList) {
		t.Fail()
		t.Fatalf("TestGetArcTitles failed, expected %v got %v", expectedArcTitleList, resultArcTitleList)
	}
	for idx := range resultArcTitleList {
		if expectedArcTitleList[idx] != resultArcTitleList[idx] {
			t.Fail()
			t.Fatalf("TestGetArcTitles failed, expected %v got %v", expectedArcTitleList, resultArcTitleList)
		}
	}
	t.Logf("TestGetArcTitles passed")
}

func TestStoryToString(t *testing.T) {
	setupTestVars()
	resultStoryString := testStory.toString()
	if resultStoryString != expectedStoryString {
		t.Fail()
		t.Fatalf("TestStoryToString failed, expected %v got %v", expectedStoryString, resultStoryString)
	}
	t.Logf("TestStoryToString passed")
}
