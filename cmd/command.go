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

	fmt.Println("Zoity successfully configured.")
}
