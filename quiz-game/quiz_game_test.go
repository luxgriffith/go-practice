package main

import (
	"maps"
	"sort"
	"testing"
)

var base_answers []string
var incorrect_answers []string
var bad_input []string
var nondefault_answers []string
var default_path string
var nondefault_path string
var bad_path string
var base_quiz map[string]string = make(map[string]string)
var nondefault_quiz map[string]string = make(map[string]string)

func setup() {
	default_path = "problems.csv"
	nondefault_path = "moreproblems.csv"
	bad_path = "lollmao"
	nondefault_quiz = map[string]string{"5+6": "11", "1+3": "4", "8+3": "11", "1+2": "3", "8+8": "16", "3+10": "14", "1+4": "5", "5+1": "6", "2+3": "5", "3+3": "6", "2+4": "6", "5+2": "7"}
	base_quiz = map[string]string{"5+5": "10", "1+1": "2", "8+3": "11", "1+2": "3", "8+6": "14", "3+1": "4", "1+4": "5", "5+1": "6", "2+3": "5", "3+3": "6", "2+4": "6", "5+2": "7"}
	base_answers = []string{"14", "11", "10", "7", "6", "6", "4", "6", "5", "5", "3", "2"}
	incorrect_answers = []string{"110", "125", "111", "131", "114", "14", "15", "64", "15", "16", "16", "17"}
	nondefault_answers = []string{"16", "11", "11", "7", "6", "6", "14", "6", "5", "5", "4", "3"}
	bad_input = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
}

func getSortedValues(input map[string]string) []string {
	sortedKeys := getSortedKeys(input)
	output := make([]string, 0, len(sortedKeys))
	for _, sortedKey := range sortedKeys {
		output = append(output, input[sortedKey])
	}
	return output
}

func getSortedKeys(input map[string]string) []string {
	keys := getKeys(input)
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return keys
}

func TestFileReader(t *testing.T) {
	setup()
	file_value, err := readFile(default_path)
	if err != nil {
		t.Fail()
		t.Fatalf("TestFileReader fail: %v", err.Error())
	}
	if !maps.Equal(file_value, base_quiz) {
		t.Fail()
		t.Fatalf("TestFileReader fail: Expected %v, loaded %v from file instead", base_quiz, file_value)
	}
	t.Logf("TestFileReader passed")
}
func TestBadPath(t *testing.T) {
	setup()
	_, err := readFile(bad_path)
	if err != nil {
		t.Logf("TestBadPath passed, err successfully handled")
	}
}

func TestUniqueFile(t *testing.T) {
	setup()
	file_value, err := readFile(nondefault_path)
	if err != nil {
		t.Fail()
		t.Fatalf("TestUniqueFile fail: %v", err.Error())
	}
	if !maps.Equal(file_value, nondefault_quiz) {
		t.Fail()
		t.Fatalf("TestUniqueFile fail: Expected %v, loaded %v from file instead", nondefault_quiz, file_value)
	}
	t.Logf("TestUniqueFile passed")
}
func TestRightAnswers(t *testing.T) {
	setup()
	quiz, _ := readFile(default_path)
	expected_answers := getSortedValues(quiz)
	for answerIdx := range base_answers {
		answer := checkAnswer(expected_answers[answerIdx], base_answers[answerIdx])
		if !answer {
			t.Fail()
			t.Fatalf("TestRightAnswers failed: Incorrectly interpreted %v as incorrect answer", base_answers[answerIdx])
		}
	}
	t.Logf("TestRightAnswers passed")
}

func TestNondefaultAnswers(t *testing.T) {
	setup()
	quiz, _ := readFile(nondefault_path)
	expected_answers := getSortedValues(quiz)
	for answerIdx := range nondefault_answers {
		answer := checkAnswer(expected_answers[answerIdx], nondefault_answers[answerIdx])
		if !answer {
			t.Fail()
			t.Fatalf("TestNondefaultAnswers failed: Incorrectly interpreted %v as  incorrect answer", nondefault_answers[answerIdx])
		}
	}
	t.Logf("TestNondefaultAnswers passed")
}

func TestWrongAnswers(t *testing.T) {
	setup()
	quiz, _ := readFile(default_path)
	expected_answers := getSortedValues(quiz)
	sort.Sort(sort.Reverse(sort.StringSlice(base_answers)))
	for answerIdx := range incorrect_answers {
		answer := checkAnswer(expected_answers[answerIdx], incorrect_answers[answerIdx])
		if answer {
			t.Fail()
			t.Fatalf("TestWrongAnswers failed: Incorrectly interpreted %v as correct answer", incorrect_answers[answerIdx])
		}
	}
	t.Logf("TestWrongAnswers passed")
}
