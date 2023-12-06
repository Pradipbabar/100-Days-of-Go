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

func ShowTerraform() error {
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
	state, err = tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running deploy: %s", err)
	}
	fmt.Println(state)
	return nil
}
package main

import (
	"encoding/json"
	"fmt"
)

// FlattenJSON flattens a nested JSON structure into a key-value map.
func FlattenJSON(data map[string]interface{}, parentKey string, flattened map[string]interface{}) {
	for key, value := range data {
		newKey := key
		if parentKey != "" {
			newKey = fmt.Sprintf("%s.%s", parentKey, key)
		}

		switch value := value.(type) {
		case map[string]interface{}:
			FlattenJSON(value, newKey, flattened)
		default:
			flattened[newKey] = value
		}
	}
}

func main() {
	// Example Terraform state JSON data (replace this with your actual data)
	terraformStateJSON := `
	{
		"version": 4,
		"terraform_version": "1.0.6",
		"resources": [
			{
				"name": "example_resource",
				"type": "aws_instance",
				"instances": [
					{
						"index_key": "index_value",
						"attributes": {
							"id": "i-1234567890abcdef0",
							"instance_type": "t2.micro"
						}
					}
				]
			}
		]
	}
	`

	var terraformState map[string]interface{}
	err := json.Unmarshal([]byte(terraformStateJSON), &terraformState)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Flatten Terraform state into key-value pairs
	flattenedState := make(map[string]interface{})
	FlattenJSON(terraformState, "", flattenedState)

	// Print the flattened key-value pairs
	for key, value := range flattenedState {
		fmt.Printf("%s: %v\n", key, value)
	}
}
