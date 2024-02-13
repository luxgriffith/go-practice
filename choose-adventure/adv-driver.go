package main

import adventure_model "github.com/griffithscg/go-practice/choose-adventure/adventure-model"

func main() {
	story, err := adventure_model.ReadStory("story.json")
	if err != nil {
		panic(err)
	}
	go adventure_model.RunServer(story.GetArcs()["intro"])
}
