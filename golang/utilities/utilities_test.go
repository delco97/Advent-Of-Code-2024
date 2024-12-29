package utilities

import "testing"

func TestExample(t *testing.T) {
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	new_numbers := Remove(numbers, 3)
	
	// Returned slice should not contain the element at index 3.
	expected := []int{10, 20, 30, 50, 60, 70}
	for i := range new_numbers {
		if new_numbers[i] != expected[i] {
			t.Fatalf("Expected %d, got %d", expected[i], new_numbers[i])
		}
	}
	
	// Original slice should not be modified
	expected = []int{10, 20, 30, 40, 50, 60, 70}
	for i := range numbers {
		if numbers[i] != expected[i] {
			t.Fatalf("Expected %d, got %d", expected[i], new_numbers[i])
		}
	}
}