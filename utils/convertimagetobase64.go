package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func ConvertImageToBase64(imagePath string) (string, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Read the image data into a byte slice
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read image data: %w", err)
	}

	// Encode the image data to a Base64 string
	base64String := base64.StdEncoding.EncodeToString(imageData)

	// Prepend the MIME type to the Base64 string
	// Update MIME type according to your image type (e.g., image/png, image/gif)
	return "data:image/jpeg;base64," + base64String, nil
}
