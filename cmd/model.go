package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Root struct {
	PathZoity string     `json:"path_zoity"`
	Services  []Service  `json:"services"`
	Sequences []Sequence `json:"sequences"`
}

type Service struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Path    string `json:"path"`
	Port    string `json:"port"`
}

type Sequence struct {
	Name     string   `json:"name"`
	Services []string `json:"services"`
}

func getConfig() (config *Root) {
	raw, err := ioutil.ReadFile(pathRoot + "/config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func addConfig() {}
