package terraform

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func ApplyTerraform() error {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	workingdir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}

	// Initialize Terraform in the specified directory
	tf, err := tfexec.NewTerraform(workingdir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}
	err = tf.Apply(context.Background())
	if err != nil {
		log.Fatalf("error running deploy: %s", err)
	}
	fmt.Println("Terraform deployed successfully.")
	return nil
}
