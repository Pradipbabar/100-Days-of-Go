package action

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	
)

func ListImagePackages(client client.Client) error {
	// List images
	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}

	// Print package names for each image 
	for _, image := range images {
		for _, tag := range image.RepoTags {
			fmt.Printf("Image: %s\n", tag)
			// You can extract more information about the image if needed
			// For example, image.Labels may contain additional metadata
		}
	}

	return nil
}

