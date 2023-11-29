package action

import (
	"log"
	"fmt"
	"context"

	"github.com/docker/docker/client"
)

func stopContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
		return err
	} else {
		log.Printf("Container %s Stop sucessfully",containername)
	}

	return nil