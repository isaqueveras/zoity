package cmd

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

const ZOITY_PATH_CONFIG string = "ZOITY_PATH_CONFIG"

func initt(_ *cobra.Command, _ []string) {
	root := new()
	root.load()

	if len(root.Services) != 0 {
		fmt.Println("error: settings is not empty, try reset command")
		return
	}

	root.update()
	fmt.Println("Zoity successfully configured.")
}

func get(_ *cobra.Command, _ []string) {
	cfg := new()
	cfg.load()

	fmt.Printf("| %-10s | %-25s | %-22s | %-30s|\n", "ID", "NAME", "CREATED", "COMMAND")
	fmt.Println("|------------|---------------------------|------------------------|-------------------------------|")
	for _, s := range cfg.Services {
		fmt.Printf("| %-10s | %-25s | %-22s | %-30s|\n", s.Id, s.Name, s.CreatedAt.Local().Format(time.RFC822Z), s.Command)
	}
}

func add(cmd *cobra.Command, _ []string) {
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

	id := func(length int) string {
		const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
		bytes := make([]byte, length)
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphabet[b%byte(len(alphabet))]
		}
		return string(bytes)
	}

	cfg := new()
	cfg.load()

	cfg.Services = append(cfg.Services, Service{
		Id:        id(8),
		Name:      name,
		Command:   command,
		Path:      path,
		CreatedAt: time.Now(),
	})

	cfg.update()
	fmt.Println("zoity: service configured successfully")
}

func run(_ *cobra.Command, args []string) {
	cfg := new()
	cfg.update()
	cfg.load()

	for idx := range args {
		service := cfg.searchServiceByName(args[idx])
		if service == nil {
			fmt.Println("zoity:\033[1;31m service " + args[idx] + " not found\033[0m")
			continue
		}

		cfg.kill(service.Id)

		cmd := exec.Command("/bin/bash", "-c", service.Command)
		cmd.Env, cmd.Dir = os.Environ(), service.Path

		if err := cmd.Start(); err != nil {
			fmt.Println("zoity:\033[1;31m error running the "+service.Name+" service\033[0m", err.Error())
			continue
		}

		cfg.add(cmd.Process.Pid, service.Id)
		cfg.update()

		fmt.Sprintln(fmt.Printf("zoity:\033[1;32m pid=%d: the %s service has been initialized\033[0m\n", cmd.Process.Pid, service.Name))
	}
}

func down(_ *cobra.Command, args []string) {
	cfg := new()
	cfg.update()
	cfg.load()

	for _, name := range args {
		service := cfg.searchServiceByName(name)
		if service == nil {
			fmt.Println("zoity:\033[1;31m service " + name + " not found\033[0m")
			continue
		}
		cfg.kill(service.Id)
	}
}
