package infra

import (
	"fmt"

	"github.com/Pradipbabar/autodash/pkg/infrastructure/terraform"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy Infrastructure",
	Long:  "destroy infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		err := terraform.DestoryTerraform()
		if err != nil {
			fmt.Println("Error getting working directory:", err)

		} else {
			fmt.Println("Destoying infrastructure...")
			// Implement initialization logic here
		}
		fmt.Println("destroy infrastructure...")
		// Implement initialization logic here
	},
}
