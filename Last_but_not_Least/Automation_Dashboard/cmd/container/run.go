package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	portFlag     int
	containerNameFlag string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Container",
	Long:  "Run Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
        // Use the flag values in your logic
		if portFlag > 0 {
			fmt.Printf("Running container with port: %d\n", portFlag)
		} else {
			fmt.Println("Please provide a valid port number to run a container.")
		}

		if containerNameFlag != "" {
			fmt.Printf("Running container with name: %s\n", containerNameFlag)
		} else {
			fmt.Println("Please provide a container name.")
		}
		fmt.Println("Running Container...")
		// Implement initialization logic here

	},
}
func init() {
	// Add flags to the run command
	runCmd.Flags().IntVarP(&portFlag, "port", "p", 0, "Specify port number for running the container")
	runCmd.Flags().StringVarP(&containerNameFlag, "name", "n", "", "Specify name for the running container")
}
