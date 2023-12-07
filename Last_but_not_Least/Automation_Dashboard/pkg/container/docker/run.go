// docker/docker.go
package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// RunGoApplication runs a Go application in a Docker container
func RunGoApplication(port int, containerName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	imageName := "golang:latest"

	// Pull the Golang image
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	// Build the Go application
	buildCmd := exec.Command("go", "build", "-o", "app")
	buildCmd.Dir = getAppPath()
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return err
	}

	// Get the working directory
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create a container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"./app"},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(fmt.Sprintf("%d/tcp", port)): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", port),
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		return err
	}

	// Copy the built Go application into the container
	buildFilePath := filepath.Join(workingDir, "app")
	err = copyToContainer(ctx, cli, resp.ID, buildFilePath, "/app")
	if err != nil {
		return err
	}

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	fmt.Printf("Go application running in container '%s' on port %d\n", containerName, port)
	return nil
}

// Helper function to get the path to the Go application
func getAppPath() string {
	exePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}

// Helper function to copy files or directories from the host to a container
func copyToContainer(ctx context.Context, cli *client.Client, containerID, sourcePath, targetPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	return cli.CopyToContainer(ctx, containerID, targetPath, source, types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
	})
}
