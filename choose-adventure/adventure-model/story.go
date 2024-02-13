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

func (s *Story) GetArcs() map[string]*Arc {
	return s.arcs
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
		if len(arcMap) != 3 {
			return errors.New(fmt.Sprintf("Arc %v's map has the wrong number of values", arcTitle))
		}
		var title string
		var story []string
		var options []*Option
		for arcMapKey := range arcMap {
			switch arcMapKey {
			case "title":
				titleContent, ok := arcMap[arcMapKey].(string)
				if !ok {
					return errors.New(fmt.Sprintf("Arc %v's title value isn't a string", arcTitle))
				} else {
					title = titleContent
				}
			case "story":
				if arcMap[arcMapKey] == nil {
					continue
				}
				storyTextList, ok := arcMap[arcMapKey].([]interface{})
				if !ok {
					return errors.New(fmt.Sprintf("Arc %v has a story %v that isn't a list", arcTitle, arcMap[arcMapKey]))
				} else {
					for _, paragraph := range storyTextList {
						paragraphContent, ok := paragraph.(string)
						if !ok {
							return errors.New(fmt.Sprintf("Arc %v has a paragraph in its story %v that isn't a string", arcTitle, paragraph))
						} else {
							story = append(story, paragraphContent)
						}
					}
				}
			case "options":
				if arcMap[arcMapKey] == nil {
					continue
				}
				optionsList, isList := arcMap[arcMapKey].([]interface{})
				if !isList {
					return errors.New(fmt.Sprintf("Arc %v's options aren't a list", arcTitle))
				}
				for _, optionsElement := range optionsList {
					optionsMap, isMap := optionsElement.(map[string]interface{})
					if !isMap {
						return errors.New(fmt.Sprintf("Arc %v's option %v is not a map", arcTitle, optionsElement))
					}
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
			default:
				return errors.New(fmt.Sprintf("Arc %v's map has an invalid key %v", arcTitle, arcMapKey))
			}
		}
		newStory[arcTitle] = &Arc{
			title:   title,
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
	title   string
	text    []string
	options []*Option
}

// Return a string that represents the contents of the arc for testing and debugging purposes
func (a *Arc) toString() string {
	out := a.title + "\t"
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

func (a *Arc) toMap() map[string]interface{} {
	out := make(map[string]interface{})
	out["title"] = a.title
	out["story"] = a.text
	outOptions := make([]interface{}, len(a.options))
	for _, option := range a.options {
		outOptions = append(outOptions, option.toMap())
	}
	out["options"] = outOptions
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

func (o *Option) toMap() map[string]interface{} {
	out := make(map[string]interface{})
	out["text"] = o.text
	out["arcTitle"] = o.arcTitle
	return out
}
