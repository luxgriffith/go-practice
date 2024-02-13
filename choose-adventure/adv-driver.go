package main

import (
	"encoding/json"
	"log"
	"net/http"

	adventure_model "github.com/griffithscg/go-practice/choose-adventure/adventure-model"
)

func main() {
	story, err := adventure_model.ReadStory("story.json")
	if err != nil {
		panic(err)
	}
	adventure_model.RunServer(story, loadPage)
}

func loadPage(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("adventure-view/base-page/"))
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Loading base page"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
