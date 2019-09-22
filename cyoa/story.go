package cyoa

import (
	"encoding/json"
	"io"
)

// JSONStory parses the JSON file into a Story
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story contains all chapters
type Story map[string]Chapter

// Chapter of our story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option of our chapter
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
