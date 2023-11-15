package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Root struct {
	Services  []Service  `json:"services"`
	Sequences []Sequence `json:"sequences"`
}

type Service struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Command   string    `json:"command"`
	Path      string    `json:"path"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"created_at"`
}

type Sequence struct {
	Name     string   `json:"name"`
	Services []string `json:"services"`
}

func getConfig() (config *Root) {
	raw, err := os.ReadFile(pathRoot + "/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		fmt.Println(err)
		return
	}

	return config
}

func updateConfig(cfg *Root) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(pathRoot+"/config.json", bytes, os.ModePerm)
}
