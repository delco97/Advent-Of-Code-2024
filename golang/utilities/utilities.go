package utilities

import (
	"os"
	"sort"
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

func BinarySearchInt(array []int, target int) int {
	i := sort.Search(len(array), func(i int) bool {
		return array[i] >= target
	}) 
	if i < len(array) && array[i] == target {
		return i
	} else {
		return -1
	}
}

func CountOccurrencesInSortedSliceInt(array []int, target int) int {
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