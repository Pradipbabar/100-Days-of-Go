package action
import (
	"log"
	"fmt"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Stop and remove a container
func RemoveContainer(client *client.Client, containername string) error {
	ctx := context.Background()


	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	} 

	return nil
}
