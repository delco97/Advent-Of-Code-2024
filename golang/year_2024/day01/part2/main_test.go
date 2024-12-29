package main

import (
	"os"
	"testing"
	"strings"
	"github.com/delco97/advent-of-code/utilities"
)

func TestExample(t *testing.T) {
	content := strings.TrimSpace(`
	3   4
	4   3
	2   5
	1   3
	3   9
	3   3
	`)
	tmpFile, err := utilities.CreateTempFile(content)
	if err != nil {
		t.Fatalf("An error occurred while creating the temp file %v", err)
	}
	defer os.Remove(tmpFile.Name())
	result := solve(tmpFile.Name())
	if result != 31 {
		t.Fatalf("Expected 31, got %d", result)
	}
}
