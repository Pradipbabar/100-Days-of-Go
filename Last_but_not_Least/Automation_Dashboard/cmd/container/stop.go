package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	
	stopcontainerNameFlag string
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop Container",
	Long:  "stop Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
       // Use the flag values in your logic


		if stopcontainerNameFlag != "" {
			fmt.Printf("Deleting container with name: %s\n", stopcontainerNameFlag)
		} else {
			fmt.Println("Please provide a container name to delete.")
		}
		fmt.Println("stoping Container...")
		// Implement initialization logic here

	},
}
func init() {
	// Add flags to the delete command
	stopCmd.Flags().StringVarP(&stopcontainerNameFlag, "container", "c", "", "Specify container name for deletion")
}
