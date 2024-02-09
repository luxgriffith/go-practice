package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

}

func checkAnswer(question string, answer string) (bool, bool) {
	return false, false
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
