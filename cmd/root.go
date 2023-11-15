package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	var cli = &cobra.Command{
		Use:   "zoity",
		Short: "Zoity is an orchestrator for configuring and running services locally.",
	}

	cli.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Use to initialize Zoity configuration",
		Run:   commandInit,
	})

	cli.AddCommand(&cobra.Command{
		Use: "version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Zoity version v0.0.0")
		},
	})

	commandAddService := &cobra.Command{
		Use:     "add",
		Short:   "Add a service",
		Run:     add,
		Example: `zoity --name myservice --port 8987 --command "go run *.go" --path ~/path-your-service`,
	}

	commandAddService.Flags().String("name", "", "Name of your service")
	commandAddService.Flags().String("port", "", "Port of your service")
	commandAddService.Flags().String("path", "", "Path of your service")
	commandAddService.Flags().String("command", "", "Command of your service")

	cli.AddCommand(commandAddService)

	cli.AddCommand(&cobra.Command{
		Use:   "services",
		Short: "get services",
		Run:   get,
	})

	cli.AddCommand(&cobra.Command{
		Use:     "run",
		Short:   "run service",
		Example: `zoity run powersso powersso-ui`,
		Run:     run,
	})

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
