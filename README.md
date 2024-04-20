# NestFile

NestFile is a lightweight, easy-to-use file browser for managing files on your local server, inspired by [FileBrowser](https://github.com/filebrowser/filebrowser). It provides a simple and intuitive interface for navigating, uploading, downloading, and organizing files and directories.


## Features

- **File Management Interface:** NestFile offers an intuitive web interface for managing files within specified directories on your server.
- **File Operations:** Perform common file operations such as uploading, downloading, renaming, deleting, and previewing files directly within the interface.
- **User Management:** Create multiple users, each with their own directory for managing files independently.
- **Standalone Usage:** NestFile can be used as a standalone application, allowing easy integration into existing server environments.
- **Customizable Settings:** Configure NestFile to specify directories for users, set permissions, and customize appearance and functionality.

## Technologies Used

- **Backend:** NestFile is built using Go, a robust and efficient programming language known for its concurrency support and performance.
- **Frontend:** The frontend of NestFile is developed using React, a popular JavaScript library for building interactive user interfaces.
- **File System Operations:** NestFile utilizes Go's standard library for file system operations, ensuring reliability and compatibility across platforms.
- **User Authentication:** User authentication is implemented using JWT (JSON Web Tokens) for secure access control.

## Getting Started

To get started with NestFile, follow these steps:

1. **Clone the Repository:** Clone the NestFile repository to your local machine:

```
  git clone https://github.com/yourusername/nestfile.git
```

2. **Install Dependencies:** Navigate to the project directory and install dependencies for both the backend (Go) and frontend (React):

```
  cd nestfile
  npm install # Install frontend dependencies
  go mod download # Install backend dependencies
```
3. **Configure Settings:** Customize NestFile by configuring settings such as user directories, permissions, and authentication method.

4. **Build and Run:** Build the frontend assets and run the backend server:

```
  npm run build # Build frontend assets
  go run main.go # Run the backend server
```

5. **Access NestFile:** Once the server is running, access NestFile in your web browser at `http://localhost:8000` (or the specified port).

## Contributing

Contributions to NestFile are welcome! If you find any bugs or have suggestions for new features, please open an issue or submit a pull request. Before contributing, please review our [contribution guidelines](CONTRIBUTING.md).

## License

NestFile is open-source software licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code for your own purposes.


