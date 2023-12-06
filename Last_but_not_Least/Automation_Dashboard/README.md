# Autodash - DevOps Automation CLI

Autodash is a powerful Go-based CLI application designed to simplify DevOps tasks by providing automation for both infrastructure management and container orchestration. This tool is an outcome of the "100 Days of Go" challenge and comes equipped with intuitive commands for seamless workflow integration.

## Installation

1. Ensure you have Go installed on your machine.
2. Clone the repository:

    ```bash
    git clone https://github.com/Pradipbabar/100-Days-of-Go.git
    ```

3. Navigate to the Autodash directory:

    ```bash
    cd 100-Days-of-Go/Last_but_not_Least/Automation_Dashboard
    ```

4. Build the executable:

    ```bash
    go build -o autodash cmd/main.go
    ```

5. Move the executable to a directory in your system's PATH.

## Usage

Autodash provides two main categories of commands: `infra` for managing infrastructure and `container` for handling container-related tasks.

### Infrastructure Management

#### Initialize Infrastructure

```bash
autodash infra init
```

#### Deploy Infrastructure

```bash
autodash infra deploy
```

#### Destroy Infrastructure

```bash
autodash infra destroy
```

#### Show Infrastructure Details

```bash
autodash infra show
```

### Container Orchestration

#### Delete Container

```bash
autodash container delete -c containername
```

#### Delete Image

```bash
autodash container delete -i imagename
```

#### Run Container

```bash
autodash container run -p port -n name
```

#### Show Container Details

```bash
autodash container show -a 
```

## Command Details

- **Infra Commands:**
  - `autodash infra init`: Initialize infrastructure.
  - `autodash infra deploy`: Deploy infrastructure.
  - `autodash infra destroy`: Destroy infrastructure.
  - `autodash infra show`: Show infrastructure details.

- **Container Commands:**
  - `autodash container delete -c containername`: Delete a container by name.
  - `autodash container delete -i imagename`: Delete an image by name.
  - `autodash container run -p port -n name`: Run a container with specified port and name.
  - `autodash container show -a`: Show all container and image details.
  - `autodash container show -i`: Show all  image details.
  - `autodash container show -c`: Show all container  details.

## Configuration

Customize Autodash by providing necessary configurations, such as infrastructure details and container parameters.

## Contributions

Contributions are welcome! Feel free to open issues, suggest improvements, or submit pull requests. Let's collaborate to make Autodash even more robust.

## License

This project is licensed under the [MIT License](LICENSE).

