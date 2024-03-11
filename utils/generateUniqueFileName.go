package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GenerateUniqueFileName(fileName string) string {
	// Generate a unique file name by appending a timestamp
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)
	uniqueName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), extension)
	return uniqueName
}
