package action

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/client"
)

func listAllContainer(client *client.Client) {

	// List containers
	containers, err := client.ContainerList(context.Background())
	if err != nil {
		log.Fatal(err)
		// return err

	}

	// Print container information
	for _, container := range containers {
		fmt.Printf("Container ID: %s\n", container.ID)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Command: %s\n", container.Command)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Println("---------------------------")
	}
	// return nil
}
