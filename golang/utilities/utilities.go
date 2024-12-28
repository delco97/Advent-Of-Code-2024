package utilities

import (
	"os"
	"strings"
)

func CreateTempFile(content string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "test-*.txt")
	if err != nil {
		return nil, err
	}
	trimmed_content := strings.TrimSpace(content)
	if _, err := tmpFile.Write([]byte(trimmed_content)); err != nil {
		return nil, err
	}
	return tmpFile, nil
}