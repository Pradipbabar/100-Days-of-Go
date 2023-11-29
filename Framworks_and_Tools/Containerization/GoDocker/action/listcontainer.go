package action

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/client"
)

func listAllContainer(cli *client.Client) error{

	// List containers
	containers, err := cli.ContainerList(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
    return err

	}

	// Print container information
	for _, container := range containers {
		fmt.Printf("Container ID: %s\n", container.ID)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Command: %s\n", container.Command)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Println("---------------------------")
	}
  return nil
}
