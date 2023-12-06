// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/Pradipbabar/autodash/cmd/container"
	"github.com/Pradipbabar/autodash/cmd/infra"

	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

func init() {
	RootCmd = &cobra.Command{
		Use:   "autodes",
		Short: "DevOps Pipeline Orchestrator CLI",
		Long:  "A CLI tool for managing various DevOps tasks.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to DevOps Pipeline Orchestrator CLI!")
		},
	}

	if err := configureCommands(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureCommands() error {
	// Add additional configuration logic here
	RootCmd.AddCommand(infra.InfraCmd)
	RootCmd.AddCommand(container.ContainerCmd)

	return nil
}
