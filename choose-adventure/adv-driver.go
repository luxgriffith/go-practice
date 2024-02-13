package main

import (
	"net/http"

	adventure_controller "github.com/griffithscg/go-practice/choose-adventure/adventure-controller"
	adventure_model "github.com/griffithscg/go-practice/choose-adventure/adventure-model"
)

func main() {
	story, err := adventure_model.ReadStory("story.json")
	if err != nil {
		panic(err)
	}
	adventure_controller.RunServer(story, http.FileServer(http.Dir("./adventure-view/base-page")), http.FileServer(http.Dir("./adventure-view/error-page")))
}
