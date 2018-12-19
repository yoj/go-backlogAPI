package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Projects struct
type Projects []struct {
	ID                                int    `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	UseWikiTreeView                   bool   `json:"useWikiTreeView"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
	DisplayOrder                      int    `json:"displayOrder"`
}

// getJSON return array
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	Projects := Projects{}
	err := getJSON("https://xxx.backlog.com/api/v2/projects?apiKey=uiyatsu", &Projects)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, project := range Projects {
		fmt.Println(project)
	}
}
