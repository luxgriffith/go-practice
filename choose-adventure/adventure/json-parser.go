package adventure

// The Object that defines the story
type Story struct {
	arcs map[string]*Arc
}

// The object that defines an option in an arc
type Option struct {
	text     string
	arcTitle string
}

// The Object that defines a specific arc
type Arc struct {
	text    string
	options []*Option
}

// Method that takes in a file path, and returns a Story object and nil on success or nil and an error on failure
func readStory(path string) (*Story, error) {
	return nil, nil
}
