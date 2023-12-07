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
			err := docker.ListContainer()
			if err != nil {
				fmt.Println("Error ", err)
	
			}
			err := docker.ListImage()
			if err != nil {
				fmt.Println("Error ", err)
	
			}
		}
		if containerFlag {
			fmt.Printf("Showing details for container: \n" )
			err := docker.ListContainer()
			if err != nil {
				fmt.Println("Error ", err)
	
			}
		}
		if imageFlag {
			fmt.Printf("Showing details for image: \n")
			err := docker.ListImage()
			if err != nil {
				fmt.Println("Error ", err)
	
			}
		} else {
			fmt.Println ("error")
		}
		// Implement initialization logic here

	},
}

func init() {
	// Add bool flags to the show command
	showCmd.Flags().BoolVarP(&showAllFlag, "all", "a", false, "Show details for all containers")
	showCmd.Flags().BoolVarP(&containerFlag, "container", "c", false, "Show details for a specific container")
	showCmd.Flags().BoolVarP(&imageFlag, "image", "i", false, "Show details for a specific image")

}
