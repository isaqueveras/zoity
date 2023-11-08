package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "zoity",
		Short: "Zoity is an orchestrator for configuring and running services locally.",
	}

	versionCommand = &cobra.Command{
		Use: "version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Zoity version v0.0.0")
		},
	}

	addServiceCommand = &cobra.Command{
		Use:   "add",
		Short: "Configure a service",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func Execute() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Use to initialize Zoity configuration",
		Run:   commandInit,
	}, versionCommand, addServiceCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
