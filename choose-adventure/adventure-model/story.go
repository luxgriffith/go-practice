package adventure_model

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// The Object that defines the story
type Story struct {
	arcs map[string]*Arc
}

// Return a string that represents the contents of the story for testing and debugging purposes
func (s *Story) toString() string {
	outList := make([]string, len(s.getArcTitles()))
	for _, arcTitle := range s.getArcTitles() {
		outList = append(outList, (arcTitle + ":" + s.arcs[arcTitle].toString()))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(outList)))
	return strings.Join(outList, "\n")
}

func (s *Story) buildFromMap(input map[string]interface{}) error {
	newStory := make(map[string]*Arc)
	for arcTitle := range input {
		arc := input[arcTitle]
		arcMap, isMap := arc.(map[string]interface{})
		if arcMap == nil || !isMap {
			return errors.New(fmt.Sprintf("Arc %v isn't tied to a map", arcTitle))
		}
		if len(arcMap) != 2 {
			return errors.New(fmt.Sprintf("Arc %v's map has to many values", arcTitle))
		}
		var story []string
		var options []*Option
		for arcMapKey := range arcMap {
			switch arcMapKey {
			case "options":
				optionsMapList, isMapList := arcMap[arcMapKey].([]map[string]interface{})
				if !isMapList {
					return errors.New(fmt.Sprintf("Arc %v's options aren't a list of maps", arcTitle))
				}
				for _, optionsMap := range optionsMapList {
					if len(optionsMap) != 2 {
						return errors.New(fmt.Sprintf("Arc %v has an option with the wrong number of elements", arcTitle))
					}
					var text string
					var optionArcTitle string
					for optionsMapKey := range optionsMap {
						switch optionsMapKey {
						case "text":
							textVal, ok := optionsMap[optionsMapKey].(string)
							if !ok {
								return errors.New(fmt.Sprintf("Option in arc %v has text that isn't a string: %v", arcTitle, textVal))
							} else {
								text = textVal
							}
						case "arc":
							textVal, ok := optionsMap[optionsMapKey].(string)
							if !ok {
								return errors.New(fmt.Sprintf("Option in arc %v has an arc that isn't a string: %v", arcTitle, textVal))
							} else {
								optionArcTitle = textVal
							}
						default:
							return errors.New(fmt.Sprintf("Arc %v has an option with an invalid key %v", arcTitle, optionsMapKey))
						}
					}
					options = append(options, &Option{text: text, arcTitle: optionArcTitle})
				}
			case "story":
				storyTextList, ok := arcMap[arcMapKey].([]string)
				if !ok {
					return errors.New(fmt.Sprintf("Arc %v has a story that isn't a list of strings", arcTitle))
				} else {
					story = storyTextList
				}
			default:
				return errors.New(fmt.Sprintf("Arc %v's map has an invalid key %v", arcTitle, arcMapKey))
			}
		}
		newStory[arcTitle] = &Arc{
			text:    story,
			options: options,
		}
	}
	s.arcs = newStory
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
