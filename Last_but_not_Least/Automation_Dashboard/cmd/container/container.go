// cmd/container/container.go
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage Container",
	Long:  "Commands for managing Container.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify a subcommand for Container management.")
	},
}

func init() {

	ContainerCmd.AddCommand(runCmd)
	ContainerCmd.AddCommand(showCmd)
	ContainerCmd.AddCommand(deleteCmd)
	ContainerCmd.AddCommand(stopCmd)

}
