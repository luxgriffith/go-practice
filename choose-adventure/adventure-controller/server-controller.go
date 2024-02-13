package adventure_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	adventure_model "github.com/griffithscg/go-practice/choose-adventure/adventure-model"
	error_page "github.com/griffithscg/go-practice/choose-adventure/adventure-view/error-page"
)

var currentArc *adventure_model.Arc
var story *adventure_model.Story
var errorPage bool

// Defines the thread that runs the actual http server
func RunServer(inStory *adventure_model.Story, workingSite http.Handler, errorSite http.Handler) {
	story = inStory
	currentArc = story.GetArcs()["intro"]
	errorPage = true
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToPage)
	mux.Handle("/story-page/", http.StripPrefix("/story-page/", workingSite))
	mux.Handle("/error-page/", http.StripPrefix("/error-page/", errorSite))
	mux.HandleFunc("/change-arc", changeArc)
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		panic(err)
	}
}

func redirectToPage(w http.ResponseWriter, r *http.Request) {
	if !errorPage {
		r.URL.Path = "/story-page/base-page.html"
		fmt.Printf("Redirecting to %v", r.URL.String())
		http.RedirectHandler(r.URL.String(), 301).ServeHTTP(w, r)
	} else {
		r.URL.Path = "/error-page/error-page.html"
		fmt.Printf("Redirecting to %v", r.URL.String())
		http.RedirectHandler(r.URL.String(), 301).ServeHTTP(w, r)
	}
}

func changeArc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Change Arc request recieved")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
		writeErrorResponse(w, err)
		error_page.BuildErrorPage(err)
		errorPage = true
		redirectToPage(w, r)
		return
	}
	bodyJson := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		writeErrorResponse(w, err)
		error_page.BuildErrorPage(err)
		errorPage = true
		redirectToPage(w, r)
		return
	}
	optionText, ok1 := bodyJson["text"].(string)
	optionArcTitle, ok2 := bodyJson["title"].(string)
	if !(ok1 && ok2) {
		fmt.Printf("Option Values are not strings")
		err := errors.New(fmt.Sprintf("Option Text %v or Option Arc %v are not the correct type (String)", bodyJson["text"], bodyJson["title"]))
		writeErrorResponse(w, err)
		error_page.BuildErrorPage(err)
		errorPage = true
		redirectToPage(w, r)
		return
	}
	option := adventure_model.NewOption(optionText, optionArcTitle)
	title, arc, err := getNextArc(option, story)
	if err != nil {
		fmt.Printf("Error while getting next arc")
		writeErrorResponse(w, err)
		error_page.BuildErrorPage(err)
		errorPage = true
		redirectToPage(w, r)
		return
	}
	currentArc = arc
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	arcMap := arc.ToMap()
	resp := make(map[string]interface{})
	resp["title"] = title
	resp["arc"] = arcMap
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

// Takes in an option the user picked and the story, and returns the arc that option leads to and its title, as well as an optional error
func getNextArc(option *adventure_model.Option, story *adventure_model.Story) (title string, arc *adventure_model.Arc, err error) {
	if option == nil || story == nil {
		return "", nil, errors.New(fmt.Sprintf("Recieved nil input"))
	}
	nextArcTitle := option.GetArcTitle()
	nextArc := story.GetArcs()[nextArcTitle]
	return nextArcTitle, nextArc, nil
}
