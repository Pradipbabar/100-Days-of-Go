package infra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy Infrastructure",
	Long:  "destroy infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destroy infrastructure...")
		// Implement initialization logic here
	},
}
