// docker/docker.go
package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func DeleteImage(imageName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		for _, tag := range image.RepoTags {
			// Check if the image name matches the provided imageName
			if imageName == tag {
				fmt.Print("Deleting image ", tag, "... ")
				_, err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{})
				if err != nil {
					panic(err)
				}
				fmt.Println("Success")
				return // Exit the loop after deleting the image
			}
		}
	}

	// If the loop completes without finding the image, print a message
	fmt.Println("Image not found:", imageName)
}

func DeleteContainer(containerName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		// Check if the container name matches the provided containerName
		if containerName == container.Names[0][1:] {
			fmt.Print("Deleting container ", container.ID[:10], "... ")
			if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{}); err != nil {
				panic(err)
			}
			fmt.Println("Success")
			return // Exit the loop after deleting the container
		}
	}

	// If the loop completes without finding the container, print a message
	fmt.Println("Container not found:", containerName)
}
