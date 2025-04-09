package project

import (
	"os"
	"strings"
)

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
