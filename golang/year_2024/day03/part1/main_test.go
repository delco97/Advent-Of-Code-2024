package main

import (
	"os"
	"testing"
	"strings"
	"github.com/delco97/advent-of-code/utilities"
)

func TestExample(t *testing.T) {
	content := strings.TrimSpace(`
	xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
	`)
	tmpFile, err := utilities.CreateTempFile(content)
	if err != nil {
		t.Fatalf("An error occurred while creating the temp file %v", err)
	}
	defer os.Remove(tmpFile.Name())
	result, _ := solve(tmpFile.Name())
	if result != 161 {
		t.Fatalf("Expected 161, got %d", result)
	}
}
