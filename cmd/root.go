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

	cli.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Configure a service",
		Run:   func(cmd *cobra.Command, args []string) {},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "services",
		Short: "get services",
		Run:   getServices,
	})

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
