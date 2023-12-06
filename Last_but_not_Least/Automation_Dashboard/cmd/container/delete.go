package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete Container",
	Long:  "delete Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("deleting Container...")
		// Implement initialization logic here

	},
}
