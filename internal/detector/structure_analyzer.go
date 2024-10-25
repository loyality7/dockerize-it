package detector

import (
	"os"
)

type ProjectStructure struct {
	RootDir     string
	Files       []string
	Directories []string
}

func AnalyzeStructure(projectPath string) ProjectStructure {
	structure := ProjectStructure{
		RootDir: projectPath,
	}

	entries, err := os.ReadDir(projectPath)
	if err != nil {
		return structure
	}

	for _, entry := range entries {
		if entry.IsDir() {
			structure.Directories = append(structure.Directories, entry.Name())
		} else {
			structure.Files = append(structure.Files, entry.Name())
		}
	}

	return structure
}
