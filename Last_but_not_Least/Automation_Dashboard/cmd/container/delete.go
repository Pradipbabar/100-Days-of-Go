package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	imageNameFlag     string
	containerNameFlag string
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete Container",
	Long:  "delete Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
       // Use the flag values in your logic
		if imageNameFlag != "" {
			fmt.Printf("Deleting container with image name: %s\n", imageNameFlag)
		} else {
			fmt.Println("Please provide an image name to delete a container.")
		}

		if containerNameFlag != "" {
			fmt.Printf("Deleting container with name: %s\n", containerNameFlag)
		} else {
			fmt.Println("Please provide a container name to delete.")
		}
		fmt.Println("deleting Container...")
		// Implement initialization logic here

	},
}
func init() {
	// Add flags to the delete command
	deleteCmd.Flags().StringVarP(&imageNameFlag, "image", "I", "", "Specify image name for deletion")
	deleteCmd.Flags().StringVarP(&containerNameFlag, "container", "c", "", "Specify container name for deletion")
}
