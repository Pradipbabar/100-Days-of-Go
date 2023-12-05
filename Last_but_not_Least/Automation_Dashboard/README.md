
# DevOps Automation CLI

## Overview

This project is a Command Line Interface (CLI) application built with Go and [Cobra](https://github.com/spf13/cobra) to facilitate DevOps automation tasks. The CLI serves as a centralized dashboard for managing various aspects of infrastructure, deployment, containers, monitoring, logging, and more.

## Project Structure

The project is structured as follows:

- **cmd/:** This directory contains the main entry point for the CLI application.
  - **root.go:** The entry point for the CLI application where the root command is set up using Cobra.

- **pkg/:** The packages representing different functional areas of the DevOps Automation Dashboard.
  - **infrastructure/:**
    - **terraform/:** Wrapper for Terraform integration.
    - **packer/:** Wrapper for Packer integration.
  - **deployment/:**
    - **jenkins/:** Wrapper for Jenkins CI/CD integration.
  - **container/:**
    - **docker/:** Wrapper for Docker SDK for Go.
    - **kubernetes/:** Wrapper for Kubernetes integration.
  - **monitoring/:**
    - **prometheus/:** Wrapper for Prometheus metrics integration.
  - **logging/:**
    - **elk/:** Wrapper for ELK integration.

- **scripts/:** Contains logic for executing custom Go scripts.

- **analytics/:**
  - **grafana/:** Wrapper for Grafana integration, allowing visualization of key metrics and trends.

- **security/:** Logic for security checks to ensure the security of infrastructure and application components.

- **extensibility/:** Directory focusing on creating a modular design.
  - **modular_design/:** Example (`module1.go`) of how additional functionality can be added easily.

- **scheduler/:** Logic (`tasks.go`) for scheduling routine automation tasks for maintenance purposes.

- **ui/:** Responsible for user interface design logic.
  - **interfaces/:** (`user_interface.go`) Defines how the user interacts with the CLI application.

- **main.go:** Main application setup and initialization.

- **go.mod** and **go.sum:** Go module files defining dependencies and versions.

## Getting Started

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Pradipbabar/100-Days-of-Go.git
   cd Last_but_not_Least/Automation_Dashboard
   ```

2. **Build and Run:**

   ```bash
   go build -o autodash
   ./autodash
   ```

3. **Usage:**
   - Check available commands: `./autodash --help`
   - Execute specific command: `./autodash command [flags]`

## Contributing

We welcome contributions! Feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
