package omd

import (
	"encoding/json"
	"io"
	"net/http"
)

type Result struct {
	Search       []Search `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type Search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func SearchResult(apikey, title string) (*Result, error) {
	resp, err := http.Get("https://www.omdbapi.com/?apikey=" + apikey + "&s=" + title)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
