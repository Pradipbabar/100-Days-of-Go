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
		workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %s", err)
	}

	// Initialize Terraform instance
	tf, err = tfexec.NewTerraform(workingDir, "terraform")
	if err != nil {
		log.Fatalf("Error running NewTerraform: %s", err)
	}
	},
}

func init() {

	InfraCmd.AddCommand(initCmd)
	InfraCmd.AddCommand(showCmd)
	InfraCmd.AddCommand(destroyCmd)
	InfraCmd.AddCommand(deployCmd)
}
