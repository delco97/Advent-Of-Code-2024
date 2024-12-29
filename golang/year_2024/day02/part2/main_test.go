package main

import (
	"github.com/delco97/advent-of-code/utilities"
	"os"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	content := strings.TrimSpace(`
	7 6 4 2 1
	1 2 7 8 9
	9 7 6 2 1
	1 3 2 4 5
	8 6 4 4 1
	1 3 6 7 9
	`)
	tmpFile, err := utilities.CreateTempFile(content)
	if err != nil {
		t.Fatalf("An error occurred while creating the temp file %v", err)
	}
	defer os.Remove(tmpFile.Name())
	result := solve(tmpFile.Name())
	if result != 2 {
		t.Fatalf("Expected 2, got %d", result)
	}
}

func TestExample2(t *testing.T) {

	content := strings.TrimSpace(`
	1 10 1 10 1
	`)
	tmpFile, err := utilities.CreateTempFile(content)
	if err != nil {
		t.Fatalf("An error occurred while creating the temp file %v", err)
	}
	defer os.Remove(tmpFile.Name())
	result := solve(tmpFile.Name())
	if result != 0 {
		t.Fatalf("Expected 0, got %d", result)
	}
}

func TestExample3(t *testing.T) {

	content := strings.TrimSpace(`
	1 0 3 6 7
	`)
	tmpFile, err := utilities.CreateTempFile(content)
	if err != nil {
		t.Fatalf("An error occurred while creating the temp file %v", err)
	}
	defer os.Remove(tmpFile.Name())
	result := solve(tmpFile.Name())
	if result != 1 {
		t.Fatalf("Expected 1, got %d", result)
	}
}
