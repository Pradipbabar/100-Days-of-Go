package container

import (
	"fmt"
	"os"

	"github.com/Pradipbabar/autodash/pkg/container/docker"
	"github.com/spf13/cobra"
)

var (
	runportFlag          int
	runcontainerNameFlag string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Container",
	Long:  "Run Container for your project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Use the flag values in your logic
		if runportFlag > 0 && runcontainerNameFlag != "" {
			fmt.Printf("Running container %s with port: %d\n", runcontainerNameFlag, runportFlag)
			err := docker.RunGoApplication(runportFlag, runcontainerNameFlag)
			if err != nil {
				fmt.Println("Error ", err)

			} else {
				fmt.Println("Running Container...")

			}

		} else {
			 wd, _ := os.Getwd()
			fmt.Println(wd)
			fmt.Println("Please provide a valid port number and Container Name to run a container.")
		}
	},
}

func init() {
	// Add flags to the run command
	runCmd.Flags().IntVarP(&runportFlag, "port", "p", 0, "Specify port number for running the container")
	runCmd.Flags().StringVarP(&runcontainerNameFlag, "name", "n", "", "Specify name for the running container")
}
