package cyoa

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Option struct {
	Text    string `json:"text"`
	ArcDest string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Sections struct {
	Intro     StoryArc `json:"intro"`
	NewYork   StoryArc `json:"new-york"`
	Debate    StoryArc `json:"debate"`
	SeanKelly StoryArc `json:"sean-kelly"`
	MarkBates StoryArc `json:"mark-bates"`
	Denver    StoryArc `json:"denver"`
	Home      StoryArc `json:"home"`
}

func ParseJson(filePath string) (Sections, error) {
	if !strings.HasSuffix(filePath, ".json") {
		return Sections{}, errors.New("the file isn't json file")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Sections{}, nil
	}
	defer file.Close()

	var sections Sections
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&sections); err != nil {
		return Sections{}, err
	}
	return sections, nil
}
