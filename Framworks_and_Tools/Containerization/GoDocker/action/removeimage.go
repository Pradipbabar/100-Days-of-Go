package action
import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func RemoveImage(client client.Client, imageName string) error {
	_, err := client.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{
		Force:         true, // Use force to remove the image even if it is in use by containers
		PruneChildren: true, // Prune children (untagged or dangling images) as well
	})
	if err != nil {
		log.Printf("Unable to remove image %s: %s", imageName, err)
		return err
	}

	log.Printf("Image %s removed successfully", imageName)
	return nil
}
