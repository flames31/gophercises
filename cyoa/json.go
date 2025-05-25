package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Adventure map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonConverter(r io.Reader) (Adventure, error) {
	adventure := make(map[string]Chapter)
	d := json.NewDecoder(r)
	err := d.Decode(&adventure)
	if err != nil {
		return Adventure{}, fmt.Errorf("error decoding json : %w", err)
	}

	return adventure, nil
}
