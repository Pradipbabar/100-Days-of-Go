package main

import (
	"log"
	"godocker/action"
	"github.com/docker/docker/client"
)

func main()  {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Unable to create docker client")
	}

	// List All Containers
	// action.ListAllContainer(cli)

	// Stop Container
	// action.StopContainer(cli, "ContainerName")

	// // Remove Container
	// action.RemoveContainer(cli, "ContainerName")

	// List Image
	action.ListImagePackages(cli)

	// Build Image 
	// Client, imagename and Dockerfile location
	tags := []string{"this_is_a_imagename"}
	dockerfile := "Dockerfile"
	err = action.BuildImage(cli, tags, dockerfile)
	if err != nil {
		log.Println(err)
	}

	// run contianer
	
	imagename := "imagename"
	containername := "containername"
	portopening := "8080"
	inputEnv := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	err = action.RunContainer(cli, imagename, containername, portopening, inputEnv)
	if err != nil {
		log.Println(err)
	}

	
}
