# Day 91-100: Advanced Projects

## [DevOps Automation Dashboard Roadmap](/Last_but_not_Least/Automation_Dashboard/)

### Day 90

**Task:** Set up the project structure.

1. Create a new Go project for your CLI application.
2. Initialize a Go module.
3. Set up the main application file.

### Day 91

**Task:** Implement the basic Cobra CLI framework.

1. Integrate the Cobra library into your project.
2. Define the root command for your CLI application.
3. Add a simple command that prints a welcome message.

### Day 92

**Task:** Implement Terraform initialization command.

1. Add a Cobra subcommand for initializing Terraform.
2. Use the `github.com/hashicorp/terraform-exec/tfexec` library to run Terraform initialization.
3. Test the initialization command.

### Day 93

**Task:** Implement Terraform apply and destroy commands.

1. Add Cobra subcommands for applying and destroying Terraform configurations.
2. Utilize the `tfexec` library to run Terraform apply and destroy.
3. Test the apply and destroy commands.

### Day 94

**Task:** Implement Docker build and run commands.

1. Add Cobra subcommands for building and running Docker images.
2. Utilize the Docker SDK for Go (`github.com/docker/docker/client`) for building and running Docker images.
3. Test the build and run commands.

### Day 95

**Task:** Add options and flags.

1. Enhance your commands by adding options and flags.
2. Allow users to provide input parameters such as Terraform configuration files, Dockerfile paths, etc.
3. Update your commands to use the provided input.

### Day 96

**Task:** Implement a combined task.

1. Create a new Cobra command that orchestrates Terraform and Docker tasks.
2. Allow users to execute Terraform commands followed by Docker commands or vice versa.
3. Test the combined task.

### Day 97

**Task:** Implement error handling and logging.

1. Enhance your CLI application with proper error handling.
2. Implement logging to track the execution flow and errors.
3. Ensure that users receive meaningful error messages.

### Day 98

**Task:** Add support for environment variables.

1. Allow users to configure certain parameters using environment variables.
2. Document the supported environment variables.
3. Update your commands to use environment variables where applicable.

### Day 99

**Task:** Implement command-line prompts.

1. Use a library like `github.com/AlecAivazis/survey` to prompt users for input.
2. Add interactive prompts for critical information, such as confirming destructive actions.
3. Test the interactive prompts.

### Day 100

**Task:** Documentation and Finalization.

1. Create a README.md file explaining how to use your CLI application.
2. Document command usage, flags, environment variables, and examples.
3. Ensure that your code is well-documented.
4. Test your application thoroughly.
5. Celebrate the completion of your CLI application!
