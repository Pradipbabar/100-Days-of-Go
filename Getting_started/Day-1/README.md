# Day 1

## Setting up Go Environment

### Step 1: Downloading and Installing Go

1. Visit the official [Go downloads page](https://golang.org/dl/) and download the appropriate installer for your operating system.
2. Run the installer and follow the installation instructions provided by the installer.

### Step 2: Setting Up Workspace and Environment Variables

1. **Setting Up Workspace**:
   - Create a directory that will serve as your workspace. This directory should be the root of all your Go code and packages. For example, you might create a directory named `go` in your home directory.
   - Inside the workspace directory, create three subdirectories: `src`, `pkg`, and `bin`. These directories will hold the source code, package objects, and binaries respectively. Your workspace directory structure should look like this:

```plaintext
   ~/go/
       |- src/
       |- pkg/
       |- bin/
```

2. **Setting Up Environment Variables**:
   - Add the Go binary directory to your `PATH` environment variable. This will allow you to run Go commands from the terminal without specifying the full path to the executable.
   - Additionally, set the `GOPATH` environment variable to point to your workspace directory. This is essential for Go to locate your source code, packages, and binaries.

   On Unix-based systems, you can do this by adding the following lines to your `.bashrc` or `.bash_profile` file, which is located in your home directory:

   ```plaintext
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

   After saving the changes, you can either restart your terminal or run the following command to apply the changes immediately:

   ```bash
   source ~/.bashrc
   ```

   On Windows systems, you can set environment variables through the Control Panel:
   - Open the Control Panel and navigate to System and Security > System > Advanced System Settings.
   - Click on the "Environment Variables" button.
   - Under the "System variables" section, click "New" and add a variable named `GOPATH` with the value set to the path of your Go workspace.
   - Edit the `Path` variable to include the path to the Go binary directory.

### Advanced Concepts in Go

1. **Go Modules**: Go has a built-in support for dependency management through modules. Modules allow you to define, maintain, and version dependencies for your Go projects. You can initialize a new module using the `go mod init` command.

2. **Concurrency**: Go has rich support for concurrency. Goroutines are lightweight threads managed by the Go runtime, allowing you to execute functions concurrently. Channels are used for communication and synchronization between Goroutines.

3. **Interfaces and Structs**: Go supports the creation of custom data types using structs and interfaces. Structs are used to define custom data structures, while interfaces define behavior that types can implement.

4. **Standard Library**: Go comes with a powerful standard library that provides support for various functionalities like networking, cryptography, file handling, and more.

5. **Error Handling**: Go encourages explicit error handling. It uses the idiomatic approach of returning errors as multiple return values, making it easy to handle errors explicitly.

Make sure to explore these advanced concepts as you progress in your Go programming journey.

### Download and install the Go programming language on your computer

- `curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz`
- `sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz`
- `sudo nano ~/.profile`
- `export PATH=$PATH:/usr/local/go/bin` add in ./profile
- `source ~/.profile`
- `go version`

#### Source

- <https://www.techtarget.com/searchitoperations/definition/Go-programming-language>
- <https://go.dev/doc/tutorial/getting-started>
- <https://www.codecademy.com/resources/blog/what-is-go/>
- <https://go.dev/doc/install>
- <https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04>
