package cyoa

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
