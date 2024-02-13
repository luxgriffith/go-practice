package adventure_model

import (
	"errors"
	"fmt"
)

// Defines the thread that runs the actual http server
func runServer() {

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
