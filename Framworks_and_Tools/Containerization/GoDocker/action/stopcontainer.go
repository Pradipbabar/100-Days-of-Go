package action

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

func stopContainer(client *client.Client, containername string) {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
		// return err
	} else {
		log.Printf("Container %s Stop sucessfully", containername)
	}

	// return nil
}
