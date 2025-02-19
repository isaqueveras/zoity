package internal

import (
	"fmt"
	"os"

	"github.com/isaqueveras/zoity/types"
	"gopkg.in/yaml.v3"
)

var (
	services   []types.Service
	configFile string

	totalServiceRunning int
)

func init() {
	if configFile = os.Getenv("ZOITY_CONFIG"); configFile == "" {
		fmt.Println("[ERROR] Indicate where the service configuration file is located in the ZOITY_CONFIG variable")
		return
	}

	raw, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = yaml.Unmarshal(raw, &services); err != nil {
		fmt.Println(err)
		return
	}
}
