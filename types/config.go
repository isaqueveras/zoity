package types

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Services            []Service
	ConfigFile          string
	TotalServiceRunning int
)

func init() {
	InitConfig()
}

func InitConfig() {
	if ConfigFile = os.Getenv("ZOITY_CONFIG"); ConfigFile == "" {
		fmt.Println("[ERROR] environment variable ZOITY_CONFIG not found")
		return
	}

	raw, err := os.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = yaml.Unmarshal(raw, &Services); err != nil {
		fmt.Println(err)
		return
	}
}
