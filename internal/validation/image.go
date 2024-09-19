package validation

import (
	"strings"
)

func IsValidImageExtension(imagePath string, allowedExtensions []string) bool {
	for _, allowedExtension := range allowedExtensions {
		if strings.HasSuffix(imagePath, allowedExtension) {
			return true
		}
	}
	return false
}
