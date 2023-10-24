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

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter

func ParseJson(filePath string) (Story, error) {
	if !strings.HasSuffix(filePath, ".json") {
		return nil, errors.New("the file isn't json file")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var story Story
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
