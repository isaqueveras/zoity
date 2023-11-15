package cmd

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const pathRoot = "/etc/zoity"

func commandInit(_ *cobra.Command, _ []string) {
	if err := os.MkdirAll(pathRoot, os.ModePerm); err != nil {
		if err == os.ErrPermission {
			fmt.Println("zoity: permission denied")
			return
		}
		fmt.Println(err)
		return
	}

	file, err := os.Create(pathRoot + "/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Chmod(os.ModePerm)

	if err = updateConfig(&Root{Services: []Service{}, Sequences: []Sequence{}}); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zoity successfully configured.")
}

func getServices(_ *cobra.Command, _ []string) {
	cfg := getConfig()
	fmt.Printf("\n| %-10s | %-25s | %-10s | %-15s | %-30s|\n", "ID", "NAME", "PORT", "CREATED", "COMMAND")
	fmt.Println("|------------|---------------------------|------------|-----------------|-------------------------------|")
	for _, s := range cfg.Services {
		fmt.Printf("| %-10s | %-25s | %-10s | %-15s | %-30s|\n", s.Id, s.Name, ":"+s.Port, s.CreatedAt.Local().Format(time.DateOnly), s.Command)
	}
}

func addService(cmd *cobra.Command, _ []string) {
	flags := cmd.Flags()

	name, _ := flags.GetString("name")
	if name == "" {
		fmt.Println("zoity: the --name flag is mandatory")
		return
	}

	path, _ := flags.GetString("path")
	if path == "" {
		fmt.Println("zoity: the --path flag is mandatory")
		return
	}

	command, _ := flags.GetString("command")
	if command == "" {
		fmt.Println("zoity: the --command flag is mandatory")
		return
	}

	port, _ := flags.GetString("port")
	if port == "" {
		fmt.Println("zoity: the --port flag is mandatory")
		return
	}

	id := func(length int) string {
		const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
		bytes := make([]byte, length)
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphabet[b%byte(len(alphabet))]
		}
		return string(bytes)
	}

	cfg := getConfig()
	cfg.Services = append(cfg.Services, Service{
		Id:        id(8),
		Name:      name,
		Command:   command,
		Path:      path,
		Port:      port,
		CreatedAt: time.Now(),
	})

	if err := updateConfig(cfg); err != nil {
		fmt.Println("zoity: " + err.Error())
		return
	}

	fmt.Println("zoity: service configured successfully")
}
