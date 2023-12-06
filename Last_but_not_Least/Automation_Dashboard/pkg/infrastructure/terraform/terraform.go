package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform/exec"
	"github.com/spf13/cobra"
)
func initTerraform() error {
	// Create a new Terraform executor
	executor := exec.NewExecutor()

	// Initialize Terraform in the specified directory
	initArgs := []string{"init"}
	initOpts := &exec.Options{
		WorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
  },
	}

	// Execute the Terraform init command
	initExitCode, err := executor.Run(initArgs, initOpts)
	if err != nil {
		return fmt.Errorf("failed to execute Terraform init: %v", err)
	}

	if initExitCode != 0 {
		return fmt.Errorf("Terraform init exited with non-zero status code: %d", initExitCode)
	}

	fmt.Println("Terraform initialized successfully.")
	return nil
  }
