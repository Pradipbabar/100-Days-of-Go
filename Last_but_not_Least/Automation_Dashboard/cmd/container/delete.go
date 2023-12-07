package container

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Pradipbabar/autodash/pkg/container/docker"
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
			err := docker.DeleteImage(deleteimageNameFlag)
			if err != nil {
				fmt.Println("Error ", err)
	
			} else {
				fmt.Println("Deleting Image...")
				
			}
		} else {
			fmt.Println("Please provide an image name to delete a container.")
		}

		if deletecontainerNameFlag != "" {
			fmt.Printf("Deleting container with name: %s\n", deletecontainerNameFlag)
			err := docker.DeleteContainer(deletecontainerNameFlag)
			if err != nil {
				fmt.Println("Error ", err)
	
			} else {
				fmt.Println("Deleting Container...")
				
			}
		} else {
			fmt.Println("Please provide a container name to delete.")
		}

	},
}
func init() {
	// Add flags to the delete command
	deleteCmd.Flags().StringVarP(&deleteimageNameFlag, "image", "I", "", "Specify image name for deletion")
	deleteCmd.Flags().StringVarP(&deletecontainerNameFlag, "container", "c", "", "Specify container name for deletion")
}
