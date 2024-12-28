package utils

import (
	"os"
	"strings"
)

// Item represents a file or folder with its type
type Item struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ScanCurrentFolder scans only the current folder with filtering options
func ScanCurrentFolder(path string, filterType string, extension string, search string) ([]Item, error) {
	var result []Item

	// Read the directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		itemType := "file"
		if entry.IsDir() {
			itemType = "folder"
		}

		// Apply filters
		if filterType == "file" && itemType != "file" {
			continue
		}
		if filterType == "folder" && itemType != "folder" {
			continue
		}
		if extension != "" && itemType == "file" && !strings.HasSuffix(entry.Name(), extension) {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(entry.Name()), strings.ToLower(search)) {
			continue
		}

		// Add the filtered item to the result
		result = append(result, Item{
			Name: entry.Name(),
			Type: itemType,
		})
	}

	return result, nil
}
