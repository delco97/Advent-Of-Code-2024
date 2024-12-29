package utilities

import (
	"os"
	"sort"
	"strings"
	"golang.org/x/exp/constraints"
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

func BinarySearchInt[T constraints.Ordered](array []T, target T) int {
	i := sort.Search(len(array), func(i int) bool {
		return array[i] >= target
	}) 
	if i < len(array) && array[i] == target {
		return i
	} else {
		return -1
	}
}

func CountOccurrencesInSortedSliceInt[T constraints.Ordered](array []T, target T) int {
	count := 0
	i := BinarySearchInt(array, target)
	if i == -1 {
		return count
	}
	for ;i < len(array) && array[i] == target; i++ {
		count += 1
	}
	return count
}

func Remove[T any](slice []T, index int) []T {
    slice_copy := make([]T, len(slice))
    copy(slice_copy, slice)
	return append(slice_copy[:index], slice_copy[index+1:]...)
}