package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show Container",
	Long:  "show Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("showing Running Container...")
		// Implement initialization logic here

	},
}
