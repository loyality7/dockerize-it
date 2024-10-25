package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/loyality7/dockerize-it/internal/detector"
	"github.com/loyality7/dockerize-it/internal/generator"
)

func main() {
	projectPath := flag.String("path", ".", "Path to the project directory")
	flag.Parse()

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	fmt.Printf("Analyzing project at: %s\n", absPath)

	// Detect the stack
	stack := detector.DetectStack(absPath)
	fmt.Printf("Detected stack: %s\n", stack)

	// Analyze the project structure
	structure := detector.AnalyzeStructure(absPath)

	// Generate Dockerfile
	dockerfileContent := generator.GenerateDockerfile(stack, structure)
	dockerfilePath := filepath.Join(absPath, "Dockerfile")
	err = os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		log.Fatalf("Error writing Dockerfile: %v", err)
	}
	fmt.Println("Dockerfile generated successfully.")

	// Generate docker-compose.yml
	dockerComposeContent := generator.GenerateDockerCompose(stack, structure)
	dockerComposePath := filepath.Join(absPath, "docker-compose.yml")
	err = os.WriteFile(dockerComposePath, []byte(dockerComposeContent), 0644)
	if err != nil {
		log.Fatalf("Error writing docker-compose.yml: %v", err)
	}
	fmt.Println("docker-compose.yml generated successfully.")
}
