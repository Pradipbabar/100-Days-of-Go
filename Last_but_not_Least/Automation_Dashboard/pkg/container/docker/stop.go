// docker/docker.go
package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopContainer(containerName string) {
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
			fmt.Print("Stopping container ", container.ID[:10], "... ")
			noWaitTimeout := 0 // to not wait for the container to exit gracefully
			if err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
				panic(err)
			}
			fmt.Println("Success")
			return // Exit the loop after stopping the container
		}
	}

	// If the loop completes without finding the container, print a message
	fmt.Println("Container not found:", containerName)
}
