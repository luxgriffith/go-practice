package adventure_model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Method that takes in a file path, and returns a Story object and nil on success or nil and an error on failure
func readStory(path string) (*Story, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err.Error())
		return nil, err
	}
	var story *Story
	err = json.Unmarshal(fileData, &story)
	if err != nil {
		fmt.Printf("Error while unmarshalling json: %v\n", err.Error())
		return nil, err
	}
	return story, nil
}
