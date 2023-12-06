package infra

import (
	"fmt"

	"github.com/Pradipbabar/autodash/pkg/infrastructure/terraform"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy Infrastructure",
	Long:  "deploy infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		err := terraform.ApplyTerraform()
		if err != nil {
			fmt.Println("Error getting working directory:", err)

		} else {
			fmt.Println("Deployinging infrastructure...")
			// Implement initialization logic here
		}
		fmt.Println("deploying infrastructure...")
		// Implement initialization logic here
	},
}
