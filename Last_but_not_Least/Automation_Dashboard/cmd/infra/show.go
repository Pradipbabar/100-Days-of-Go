package infra

import (
	"fmt"

	"github.com/Pradipbabar/autodash/pkg/infrastructure/terraform"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show Infrastructure",
	Long:  "show infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		err := terraform.ShowTerraform()
		if err != nil {
			fmt.Println("Error getting working directory:", err)

		}

	},
}
