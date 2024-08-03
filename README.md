# llamacpp-go

`llama.cpp` is a application designed to interact with LLM (Large Language Model) servers. This repo provides an easy way to configure and run the server with arguments defined in config file. It uses configuration files to define server settings and executes commands to run llma.cpp servers.

## Basic Usage
```
    ./llama-server -m your_model.gguf --port 8080

    # Basic web UI can be accessed via browser: http://host:port
    # Chat completion endpoint: http://host:port/v1/chat/completions
```

## Features

- **Dynamic Configuration**: Uses YAML configuration files to define LLM server settings.
- **Strategy Pattern**: Implements a strategy pattern to execute commands based on the configuration.
- **Logging**: Utilizes `zerolog` for logging to provide detailed execution information.

## Prerequisites

- Go 1.20 or higher
- `llama-cli` command-line tool (ensure it's installed and available in your PATH)
- Follow instructions from [llama.cpp](https://github.com/ggerganov/llama.cpp)] git repo to install the binaries.

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/llama-go.git
   cd llama-go
   ```

2. **Install dependencies:**

   ```sh
   go mod download
   ```

## Configuration

Edit the `config/config.yaml` file to define your LLM server settings. Here's an example configuration:

```yaml
servers:
  - name: llm_server
    model: /path/to/your/model.gguf
    host: 0.0.0.0
    port: 50051
    threads: 40
    ctx-size: 2048
    batch-size: 512
    embedding: false
    api_key: your_api_key
  - name: embedding_server
    model: /path/to/your/model.gguf
    host: 0.0.0.0
    port: 50052
    threads: 40
    ctx-size: 2048
    batch-size: 512
    embedding: true
    api_key: your_api_key
```

## Running the Application

To run the application, execute the following command:

```sh
go run cmd/main.go
```

This will start the application and use the configuration to execute the defined LLM server strategies.

## Logging

The application uses `zerolog` for logging. Logs will be output to the console, providing detailed information about the execution process.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) to get started.

## Contact

For any questions or issues, please open an issue on the [GitHub repository](https://github.com/yourusername/llama-go/issues).

```

This README provides a comprehensive overview of the project, including installation instructions, configuration details, and how to run the application. It also includes information about logging, licensing, and how to contribute to the project.
