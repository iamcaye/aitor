package project

import (
	"fmt"
	"os"
	"strings"

	"github.com/iamcaye/aitor-cli/utils"
)

type ProjectType string

const (
	NPM_PRJ     ProjectType = "npm"
	UNKNOWN_PRJ             = "unknown"
) // suported project types

var ProjectLockFile = map[ProjectType]string{
	UNKNOWN_PRJ: "",
	NPM_PRJ:     "package-lock.json",
}

func SearchProjectFolder(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if path == os.Getenv("HOME") {
		return "", nil
	}

	if !fileInfo.IsDir() {
		return "", nil
	}

	if isProjectSupported(path) {
		return path, nil
	}

	dirs := strings.Split(path, "/")
	if len(dirs) == 1 {
		return "", nil
	}
	parentPath := strings.Join(dirs[:len(dirs)-1], "/")
	return SearchProjectFolder(parentPath)
}

func isProjectSupported(path string) bool {
	// Check if the project is a Node.js project
	if _, err := os.Stat(path + "/package.json"); err == nil {
		return true
	}

	return false
}

func RunDetector() (string, ProjectType, error) {
	// Check if the current directory is a project folder
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return "", UNKNOWN_PRJ, err
	}

	projectPath, err := SearchProjectFolder(currentDir)
	if err != nil {
		fmt.Println("Error searching project folder:", err)
		return "", UNKNOWN_PRJ, err
	}

	if projectPath == "" {
		fmt.Println("No project folder found")
		return "", UNKNOWN_PRJ, err
	}

	projectType, err := getProjectType(projectPath)
	if err != nil {
		return "", UNKNOWN_PRJ, err
	}

	return projectPath, projectType, nil
}

func getProjectType(projectPath string) (ProjectType, error) {
	files, err := utils.ReadDir(projectPath)
	if err != nil {
		return UNKNOWN_PRJ, err
	}

	var foundType ProjectType = UNKNOWN_PRJ
	i := 0
	for foundType == UNKNOWN_PRJ && i < len(files) {
		file := files[i]
		if strings.HasPrefix(file.Name(), "package-lock") {
			foundType = NPM_PRJ
		}
		i++
	}

	return foundType, nil
}
