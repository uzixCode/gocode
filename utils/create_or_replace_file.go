package utils

import "os"

func CreateOrReplaceFile(filePath string, content string) error {
	// Open the file with read-write permissions.
	// If the file does not exist, it will be created with 0644 permissions.
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file.
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
