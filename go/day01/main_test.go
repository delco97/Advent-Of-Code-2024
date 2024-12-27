package main

import (
	"testing"
)

func TestSolveExample(t *testing.T) {
	col_a := []int{
		3,
		4,
		2,
		1,
		3,
		3,
	}
	col_b := []int{
		4,
		3,
		5,
		3,
		9,
		3,
	}
	result := solve(col_a, col_b)
	if result != 11 {
		t.Error("Got ", result)
	}
}

func TestSolve(t *testing.T) {
	col_a := []int{1, 2, 3}
	col_b := []int{1, 2, 3}
	result := solve(col_a, col_b)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
