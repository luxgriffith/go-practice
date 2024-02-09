package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	quiz, err := readFile("problems.csv")
	if err != nil {
		os.Exit(-1)
	}
	reader := bufio.NewReader(os.Stdin)
	correctCount := 0
	for _, question := range getKeys(quiz) {
		fmt.Println(question)
		answer, _ := reader.ReadString('\n')
		if checkAnswer(quiz[question], answer) {
			correctCount++
		}
	}
	fmt.Printf("You got %v answers correct", correctCount)
}

func getKeys(in map[string]string) []string {
	r := make([]string, 0, len(in))
	for k := range in {
		r = append(r, k)
	}
	return r
}

func checkAnswer(correct_answer string, answer string) bool {
	return strings.Compare(correct_answer, answer) == 0
}

func readFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("ERROR:Unable to open file at path %v because of %v", path, err.Error())
		return nil, err
	}
	defer file.Close()
	quiz_csv, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Printf("ERROR: Unable to read csv because of %v", err.Error())
		return nil, err
	}
	quiz := make(map[string]string)
	for row := range quiz_csv {
		if len(quiz_csv[row]) != 2 {
			fmt.Printf("ERROR: Row %v of the quiz has the incorrect number of elements", row)
			return nil, err
		}
		quiz[string(quiz_csv[row][0])] = string(quiz_csv[row][1])
	}
	return quiz, nil
}
