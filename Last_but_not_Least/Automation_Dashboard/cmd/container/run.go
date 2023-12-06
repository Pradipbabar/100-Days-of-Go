package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Container",
	Long:  "Run Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Running Container...")
		// Implement initialization logic here

	},
}
