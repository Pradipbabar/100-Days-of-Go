package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	runportFlag     int
	runcontainerNameFlag string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Container",
	Long:  "Run Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
        // Use the flag values in your logic
		if runportFlag > 0 && runcontainerNameFlag != ""{
			fmt.Printf("Running container %s with port: %d\n", runcontainerNameFlag, runportFlag)
		} else {
			fmt.Println("Please provide a valid port number and Container Name to run a container.")
		}

		fmt.Println("Running Container...")
		// Implement initialization logic here

	},
}
func init() {
	// Add flags to the run command
	runCmd.Flags().IntVarP(&runportFlag, "port", "p", 0, "Specify port number for running the container")
	runCmd.Flags().StringVarP(&runcontainerNameFlag, "name", "n", "", "Specify name for the running container")
}
