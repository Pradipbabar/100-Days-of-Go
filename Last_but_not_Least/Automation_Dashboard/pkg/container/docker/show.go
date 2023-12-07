package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListContainer() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {

		return err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {

		return err
	}

	for _, container := range containers {
		fmt.Printf("Image Id %s \n", container.ID)

	}
	fmt.Println("no container found")
	return nil
}

func ListImage() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {

		return err
	}
	defer cli.Close()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {

		return err
	}

	for _, image := range images {
		fmt.Printf("Image Id %s \n", image.ID)

	}
	// fmt.Println("No image found")
	return nil
}
