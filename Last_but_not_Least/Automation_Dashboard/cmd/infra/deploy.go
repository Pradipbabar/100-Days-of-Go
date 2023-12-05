package infra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy Infrastructure",
	Long:  "deploy infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploying infrastructure...")
		// Implement initialization logic here
	},
}
