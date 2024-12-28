package utils

import (
	"os"
	"path/filepath"
)

func ScanFolder(path string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// Read the directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// Recursively scan subdirectory
			subDirPath := filepath.Join(path, entry.Name())
			subDirContent, err := ScanFolder(subDirPath)
			if err != nil {
				return nil, err
			}
			result[entry.Name()] = subDirContent
		} else {
			// Add file to the result
			result[entry.Name()] = nil
		}
	}

	return result, nil
}
