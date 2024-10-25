package detector

import (
	"os"
	"path/filepath"
)

func DetectStack(projectPath string) string {
	if fileExists(filepath.Join(projectPath, "package.json")) {
		return "node"
	} else if fileExists(filepath.Join(projectPath, "go.mod")) {
		return "go"
	}
	// Add more stack detection logic here
	return "unknown"
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
