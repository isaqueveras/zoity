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
	f.WriteString(`{"path": "` + path + `"}`)

	fmt.Println(f.Name())
	fmt.Println("Zoity successfully configured.")
}
