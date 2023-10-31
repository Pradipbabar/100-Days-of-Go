# Day 16

## Advanced File Operations

Advanced file handling operations in Go involve working with directories, managing file permissions, and implementing more complex file manipulation techniques. Here's an overview of these concepts and related strategies for error handling:

1. **Working with Directories**: Go provides various functions for working with directories, such as creating directories, reading directory contents, and removing directories. Use the `os` package for directory operations.

2. **File Permissions**: You can set and manage file permissions using the `os` package. Use the `Chmod` function to change the permissions of a file or directory, and the `Stat` function to retrieve file information, including permissions.

3. **Complex File Manipulation**: Advanced file manipulation involves tasks such as copying files, moving files, and renaming files. You can use the `os` package and functions like `Copy`, `Rename`, and `Remove` for these operations.

4. **Error Handling Strategies**: When working with file operations, it's crucial to implement robust error handling strategies. Use conditional statements, error checks, and the `os.IsNotExist` function to handle errors effectively. Additionally, consider using `defer` statements to ensure that resources are properly released.

5. **File Metadata and Attributes**: Go provides methods to access and manipulate file metadata and attributes. Use the `FileInfo` type from the `os` package to retrieve information about a file, such as its size, modification time, and mode.


### Sources

- <https://cihanozhan.medium.com/file-operations-in-golang-292825c9fb3d>
- <https://github.com/cihanozhan/golang-file-samples/blob/master/01-CreateEmptyFile.go>

#### Test Code By Uncomment one by one function
