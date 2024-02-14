package adventure_model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Method that takes in a file path, and returns a Story object and nil on success or nil and an error on failure
func ReadStory(path string) (*Story, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err.Error())
		return nil, err
	}
	story_json := make(map[string]interface{})
	err = json.Unmarshal(fileData, &story_json)
	if err != nil {
		fmt.Printf("Error while unmarshalling json: %v\n", err.Error())
		return nil, err
	}
	story := &Story{make(map[string]*Arc)}
	err = story.buildFromMap(story_json)
	if err != nil {
		fmt.Printf("Error while building story: %v\n", err.Error())
		return nil, err
	}
	return story, nil
}
