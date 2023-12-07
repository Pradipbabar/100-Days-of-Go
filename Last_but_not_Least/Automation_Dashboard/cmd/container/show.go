package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showAllFlag bool
	containerFlag bool
	imageFlag bool
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show Container",
	Long:  "show Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
// Check the values of bool flags and execute logic accordingly
		if showAllFlag {
			fmt.Println("Showing details for all containers...")
		}
		if containerFlag {
			fmt.Printf("Showing details for container: %s\n", containerNameFlag)
		}
		if imageFlag {
			fmt.Printf("Showing details for image: %s\n", imageNameFlag)
		} else {
			fmt.Println ()
		}
		// Implement initialization logic here

	},
}
