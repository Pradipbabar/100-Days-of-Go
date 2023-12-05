package infra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show Infrastructure",
	Long:  "show infrastructure for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showing infrastructure...")
		// Implement initialization logic here
	},
}
