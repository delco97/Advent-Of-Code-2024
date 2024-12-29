package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFileContent(fileName string) (reports [][]int, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Errore durante l'apertura del file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	matrix := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		parts_int := make([]int, len(parts))
		for i, part := range parts {
			parts_int[i], err = strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Errore durante la conversione in intero: %v\n", err)
				return nil, err
			}
		}
		matrix = append(matrix, parts_int)
	}
	return matrix, nil
}

func safeReports(reports [][]int) int {
	num_reports := len(reports)
	safe_reports := num_reports
	for i := 0; i < num_reports; i++ {
		report := reports[i]
		num_levels := len(report)
		is_increasing := report[1] >= report[0]
		for j := 1; j < num_levels; j++ {
			adjacent_difference := int(math.Abs(float64(report[j] - report[j-1])))
			if adjacent_difference == 0 || adjacent_difference > 3 {
				safe_reports -= 1
				break
			}
			if is_increasing != (report[j] >= report[j-1]) {
				safe_reports -= 1
				break
			}
		}
	}
	return safe_reports
}

func solve(filePath string) int {
	reports, err := readFileContent(filePath)
	if err != nil {
		fmt.Println("Errore durante la lettura del file")
		return -1
	}
	res := safeReports(reports)
	return res
}

func main() {
	fmt.Println(solve("../input.txt"))
}
