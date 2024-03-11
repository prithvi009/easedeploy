package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GenerateUniqueFileName(fileName string) string {
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)
	uniqueName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), extension)
	return uniqueName
}
