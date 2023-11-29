package action

import (
	"context"
	"log"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
)

func StopContainer(client *client.Client, containername string) {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, container.StopOptions{}); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
		// return err
	} else {
		log.Printf("Container %s Stop sucessfully", containername)
	}

	// return nil
}
