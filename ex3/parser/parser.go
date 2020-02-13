package parser

import (
	"encoding/json"
	"io/ioutil"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ParseJson(fileName string) (map[string]Arc, error) {
	var arcs map[string]Arc
	dataRaw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataRaw, &arcs)
	if err != nil {
		return nil, err
	}
	return arcs, nil
}
