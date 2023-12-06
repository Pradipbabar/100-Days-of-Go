package infra

import (
	"fmt"

	"github.com/Pradipbabar/autodash/pkg/infrastructure/terraform"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Infrastructure",
	Long:  "Initialize infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		err := terraform.InitTerraform()
		if err != nil {
			fmt.Println("Error getting working directory:", err)

		} else {
			fmt.Println("Initializing infrastructure...")
			// Implement initialization logic here
		}

	},
}
