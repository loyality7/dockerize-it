package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/loyality7/dockerize-it/internal/detector"
	"github.com/loyality7/dockerize-it/internal/generator"
)

func main() {
	// Use current directory as default
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	projectPath := flag.String("path", currentDir, "Path to the project directory")
	flag.Parse()

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	fmt.Printf("Analyzing project at: %s\n", absPath)

	// Analyze the project structure
	structure := detector.AnalyzeStructure(absPath)

	// Print file structure
	fmt.Println("Project structure:")
	printStructure(structure, 0)

	// Detect the stack
	stacks := detector.DetectStack(structure)
	fmt.Printf("Detected stacks: %v\n", getStackNames(stacks))

	// Generate Dockerfile
	dockerfileContent := generator.GenerateDockerfile(stacks, structure)
	dockerfilePath := filepath.Join(absPath, "Dockerfile")
	err = os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		log.Fatalf("Error writing Dockerfile: %v", err)
	}
	fmt.Println("Dockerfile generated successfully in the root directory.")

	// Generate docker-compose.yml
	dockerComposeContent := generator.GenerateDockerCompose(stacks, structure)
	dockerComposePath := filepath.Join(absPath, "docker-compose.yml")
	err = os.WriteFile(dockerComposePath, []byte(dockerComposeContent), 0644)
	if err != nil {
		log.Fatalf("Error writing docker-compose.yml: %v", err)
	}
	fmt.Println("docker-compose.yml generated successfully in the root directory.")
}

func printStructure(structure detector.ProjectStructure, level int) {
	indent := strings.Repeat("  ", level)

	fmt.Printf("%s%s\n", indent, filepath.Base(structure.RootDir))

	for _, dir := range structure.Directories {
		fmt.Printf("%s  %s/\n", indent, dir)
	}

	for _, file := range structure.Files {
		fmt.Printf("%s  %s\n", indent, file)
	}
}

func getStackNames(stacks []detector.Stack) []string {
	names := make([]string, len(stacks))
	for i, stack := range stacks {
		names[i] = stack.Name
	}
	return names
}
