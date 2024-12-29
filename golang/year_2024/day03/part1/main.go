package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MulOperation struct {
	factor1 int
	factor2 int
}

var mul_operation_regex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func extractMulOperations(text string) ([]MulOperation, error) {
	mul_operations := mul_operation_regex.FindAllStringSubmatch(text, -1)
	var result []MulOperation = make([]MulOperation, len(mul_operations))
	for _, mul_operation := range mul_operations {
		raw_num1 := mul_operation[1]
		raw_num2 := mul_operation[2]
		num1, err := strconv.Atoi(raw_num1)
		if err != nil {
			fmt.Printf("Errore durante la conversione da string a int del primo fattore: %v\n", err)
			return nil, err
		}
		num2, err := strconv.Atoi(raw_num2)
		if err != nil {
			fmt.Printf("Errore durante la conversione da string a int del secondo fattore: %v\n", err)
			return nil, err
		}		
		result = append(result, MulOperation{factor1: num1, factor2: num2})
	}
	return result, nil
}

func solve(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Errore durante l'apertura del file: %v\n", err)
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		mul_operations, err := extractMulOperations(line)
		if err != nil {
			fmt.Printf("Errore durante l'estrazione delle operazioni di moltiplicazione: %v\n", err)	
			return -1, err
		}
		for _, mul_operation := range mul_operations {
			sum += mul_operation.factor1 * mul_operation.factor2
		}
	}
	return sum, nil
}

func main() {
	content, _ := solve("../input.txt")
	fmt.Println(content)
}
