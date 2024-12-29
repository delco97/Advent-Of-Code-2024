package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/delco97/advent-of-code/utilities"
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

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}
	num_levels := len(report)
	is_increasing := report[1] >= report[0]
	for i := 1; i < num_levels; i++ {
		adjacent_difference := int(math.Abs(float64(report[i] - report[i-1])))
		if (adjacent_difference == 0 || adjacent_difference > 3) || (is_increasing != (report[i] >= report[i-1])) {
			return false
		}
	}
	return true
}

func safeReports(reports [][]int) int {
	safe_reports := 0
	for i := 0; i < len(reports); i++ {
		report := reports[i]
		if isReportSafe(report) {
			safe_reports++
		} else {
			for j := 0; j < len(report); j++ {
				fixed_report := utilities.Remove(report, j)
				if isReportSafe(fixed_report) {
					safe_reports++
					break
				}
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
