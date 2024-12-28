package utils

import "os"

func LoadFile(filePath string) (string, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
