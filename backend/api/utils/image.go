package utils

import (
	"fmt"
)

var AllowedImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
}

func GenerateRandomImage(name string, width, height uint) string {
	return fmt.Sprintf(
		"https://picsum.photos/%d/%d?random=%s",
		width,
		height,
		name,
	)
}

func IsImage(contentType string) bool {
	return AllowedImageTypes[contentType]
}
