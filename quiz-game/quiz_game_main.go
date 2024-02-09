package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pathFlag := flag.String("file", "problems.csv", "The path to the quiz file")
	flag.Parse()
	quiz, err := readFile(*pathFlag)
	if err != nil {
		os.Exit(-1)
	}
	reader := bufio.NewReader(os.Stdin)
	correctCount := 0
	for _, question := range getKeys(quiz) {
		fmt.Println(question)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)
		if checkAnswer(quiz[question], answer) {
			correctCount = correctCount + 1
		}
	}
	fmt.Printf("You got %v answers correct\n", correctCount)
}

func getKeys(in map[string]string) []string {
	r := make([]string, 0, len(in))
	for k := range in {
		r = append(r, k)
	}
	return r
}

func checkAnswer(correct_answer string, user_answer string) bool {
	return strings.Compare(correct_answer, user_answer) == 0
}

func readFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("ERROR:Unable to open file at path %v because of %v\n", path, err.Error())
		return nil, err
	}
	defer file.Close()
	quiz_csv, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Printf("ERROR: Unable to read csv because of %v\n", err.Error())
		return nil, err
	}
	quiz := make(map[string]string)
	for row := range quiz_csv {
		if len(quiz_csv[row]) != 2 {
			fmt.Printf("ERROR: Row %v of the quiz has the incorrect number of elements\n", row)
			return nil, err
		}
		quiz[string(quiz_csv[row][0])] = string(quiz_csv[row][1])
	}
	return quiz, nil
}
