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
		if portFlag > 0 && containerNameFlag != ""{
			fmt.Printf("Running container %s with port: %d\n", containerNameFlag, portFlag)
		} else {
			fmt.Println("Please provide a valid port number and Container Name to run a container.")
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
