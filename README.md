# DockerizeIt

DockerizeIt is a CLI tool that generates Dockerfile and docker-compose.yml files for various tech stacks based on the project's folder structure.

## Usage

```
./dockerize-it -path /path/to/your/project
```

## Supported Stacks

- Node.js
- Go

More stacks will be added in future updates.

## Building the Tool

To build the DockerizeIt tool, run:

```
go build -o dockerize-it cmd/dockerize-it/main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
