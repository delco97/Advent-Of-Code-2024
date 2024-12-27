package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"sort"
	"math"
	"regexp"
)


func readFileContent(fileName string) ([]int, []int, error) {
	var col1 []int
	var col2 []int
	
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Errore durante l'apertura del file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := regexp.MustCompile("\\s+").Split(line, -1)
		a, err := strconv.Atoi(parts[0])
		b, err := strconv.Atoi(parts[1])	
		if err != nil {
			fmt.Printf("Errore durante la conversione del numero: %v\n", err)
			return nil, nil, err	
		}			
		col1 = append(col1, a)
		col2 = append(col2, b)
	}
	return col1, col2, nil
}

func solve(col_a []int, col_b []int) int {
	sort.Ints(col_a)
	sort.Ints(col_b)

	if len(col_a) != len(col_b) {
		fmt.Println("Le colonne non hanno la stessa lunghezza")
		return 0
	}
	n := len(col_a)

	distances := make([]int, len(col_a))
	for i := 0; i < n; i++ {
		distances[i] = int(math.Abs(float64(col_a[i] - col_b[i])))
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += distances[i]
	}
	return sum
}

func main() {
	col_a, col_b, err := readFileContent("input.txt")
	if err != nil {
		return
	}
	res := solve(col_a, col_b)
	fmt.Println(res)
}