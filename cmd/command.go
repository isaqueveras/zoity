package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const pathRoot = "/etc/zoity"

func commandInit(_ *cobra.Command, args []string) {
	path := pathRoot
	if len(args) != 0 {
		path = args[0]
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create(path + "/config.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	f.Chmod(os.ModePerm)
	f.WriteString(`{"path_zoity": "` + path + `"}`)

	fmt.Println(f.Name())
	fmt.Println("Zoity successfully configured.")
}

func getServices(_ *cobra.Command, _ []string) {
	cfg := getConfig()

	fmt.Printf("%-10s %-25s %-10s %-15s %-5s\n", "ID", "NAME", "PORT", "COMMAND", "RUNNING")
	for _, s := range cfg.Services {
		fmt.Printf("%-10s %-25s %-10s %-15s %-5s\n", s.Id, s.Name, ":"+s.Port, s.Command, "false")
	}
}
