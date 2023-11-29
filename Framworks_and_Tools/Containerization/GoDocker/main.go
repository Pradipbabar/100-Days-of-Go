package main

import (
	"godocker/action"
	"github.com/docker/docker/client"


)

func main()  {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client")
	}

	// List All Containers
	action.listAllContainer(cli)

	// Stop Container
	action.stopContainer(cli, "ContainerName")

	// Remove Container
	action.RemoveContainer(cli, "ContainerName")


}