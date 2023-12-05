// cmd/infra/infra.go
package infra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var InfraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Manage Infrastructure",
	Long:  "Commands for managing infrastructure.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify a subcommand for infrastructure management.")
	},
}

func init() {

	InfraCmd.AddCommand(initCmd)
	InfraCmd.AddCommand(showCmd)
	InfraCmd.AddCommand(destroyCmd)
	InfraCmd.AddCommand(deployCmd)
}
