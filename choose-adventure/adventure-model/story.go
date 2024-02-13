package adventure_model

// The Object that defines the story
type Story struct {
	arcs map[string]*Arc
}

// Return a string that represents the contents of the story for testing and debugging purposes
func (s *Story) toString() string {
	out := ""
	for idx, arcTitle := range s.getArcTitles() {
		out = out + arcTitle + ":" + s.arcs[arcTitle].toString()
		if idx == len(s.getArcTitles())-1 {
			out += "\n"
		}
	}
	return out
}

func (s *Story) buildFromMap(map[string]interface{}) error {
	return nil
}

// Returns a list of the story's arc titles
func (s *Story) getArcTitles() []string {
	arcs := make([]string, 0, len(s.arcs))
	for arcTitle := range s.arcs {
		arcs = append(arcs, arcTitle)
	}
	return arcs
}

// The Object that defines a specific arc
type Arc struct {
	text    []string
	options []*Option
}

// Return a string that represents the contents of the arc for testing and debugging purposes
func (a *Arc) toString() string {
	out := ""
	if a.text != nil {
		for idx, paragraph := range a.text {
			out += paragraph
			if idx != len(a.text)-1 {
				out += ", "
			}
		}
	}
	out += "\t"
	if a.options != nil {
		for idx, option := range a.options {
			out += option.toString()
			if idx != len(a.options)-1 {
				out += "; "
			}
		}
	}
	return out
}

// The object that defines an option in an arc
type Option struct {
	text     string
	arcTitle string
}

// Return a string that represents the contents of the Option for testing and debugging purposes
func (o *Option) toString() string {
	return o.text + ", " + o.arcTitle
}
