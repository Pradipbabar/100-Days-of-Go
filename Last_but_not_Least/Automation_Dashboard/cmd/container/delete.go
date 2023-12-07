package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deleteimageNameFlag     string
	deletecontainerNameFlag string
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete Container",
	Long:  "delete Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
       // Use the flag values in your logic
		if deleteimageNameFlag != "" {
			fmt.Printf("Deleting container with image name: %s\n", deleteimageNameFlag)
		} else {
			fmt.Println("Please provide an image name to delete a container.")
		}

		if deletecontainerNameFlag != "" {
			fmt.Printf("Deleting container with name: %s\n", deletecontainerNameFlag)
		} else {
			fmt.Println("Please provide a container name to delete.")
		}
		fmt.Println("deleting Container...")
		// Implement initialization logic here

	},
}
func init() {
	// Add flags to the delete command
	deleteCmd.Flags().StringVarP(&deleteimageNameFlag, "image", "I", "", "Specify image name for deletion")
	deleteCmd.Flags().StringVarP(&deletecontainerNameFlag, "container", "c", "", "Specify container name for deletion")
}
