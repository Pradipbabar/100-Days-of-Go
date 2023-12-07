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
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer cli.Close()

	imageName := "golang:latest"

	// Pull the Golang image
	if err := pullImage(ctx, cli, imageName); err != nil {
		return fmt.Errorf("failed to pull image %s: %v", imageName, err)
	}

	// Build the Go application
	buildPath := getAppPath()
	if err := buildGoApp(buildPath); err != nil {
		return fmt.Errorf("failed to build Go application: %v", err)
	}

	// Create and start the container
	if err := createAndStartContainer(ctx, cli, imageName, port, containerName, buildPath); err != nil {
		return fmt.Errorf("failed to create and start container: %v", err)
	}

	fmt.Printf("Go application running in container '%s' on port %d\n", containerName, port)
	return nil
}

// pullImage pulls a Docker image
func pullImage(ctx context.Context, cli *client.Client, imageName string) error {
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

// buildGoApp builds the Go application
func buildGoApp(buildPath string) error {
	buildCmd := exec.Command("go", "build", "-o", "app")
	buildCmd.Dir = buildPath
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	return buildCmd.Run()
}

// createAndStartContainer creates and starts a Docker container
func createAndStartContainer(ctx context.Context, cli *client.Client, imageName string, port int, containerName string, buildPath string) error {
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
	if err := copyToContainer(ctx, cli, resp.ID, buildFilePath, "/app"); err != nil {
		return fmt.Errorf("failed to copy files to container: %v", err)
	}

	// Start the container
	return cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
}

// getAppPath gets the path to the Go application
func getAppPath() string {
	exePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}

// copyToContainer copies files or directories from the host to a container
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
