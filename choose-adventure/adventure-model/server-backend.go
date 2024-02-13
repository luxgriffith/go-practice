package adventure_model

import (
	"errors"
	"fmt"
	"net/http"
)

var CurrentArc *Arc

// Defines the thread that runs the actual http server
func RunServer(startingArc *Arc) {
	CurrentArc = startingArc
	mux := http.NewServeMux()
	mux.HandleFunc("/", presentSite)
	mux.HandleFunc("/change-arc", changeArc)
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		panic(err)
	}
}

func presentSite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(CurrentArc.toString()))
}
func changeArc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Change Arc request recieved")
}

// Takes in an option the user picked and the story, and returns the arc that option leads to and its title, as well as an optional error
func getNextArc(option *Option, story *Story) (title string, arc *Arc, err error) {
	if option == nil || story == nil {
		return "", nil, errors.New(fmt.Sprintf("Recieved nil input"))
	}
	nextArcTitle := option.arcTitle
	nextArc := story.arcs[nextArcTitle]
	return nextArcTitle, nextArc, nil
}
