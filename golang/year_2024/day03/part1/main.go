package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Errore durante l'apertura del file: %v\n", err)
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	mul_operation_regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		mul_operations := mul_operation_regex.FindAllStringSubmatch(line, -1)
		for _, mul_operation := range mul_operations {
			raw_num1 := mul_operation[1]
			raw_num2 := mul_operation[2]
			num1, err := strconv.Atoi(raw_num1)
			if err != nil {
				fmt.Printf("Errore durante la conversione in intero: %v\n", err)
				return -1, err
			}
			num2, err := strconv.Atoi(raw_num2)
			if err != nil {
				fmt.Printf("Errore durante la conversione in intero: %v\n", err)
				return -1, err
			}
			sum += num1 * num2
		}
	}
	return sum, nil
}

func main() {
	content, _ := solve("../input.txt")
	fmt.Println(content)
}
