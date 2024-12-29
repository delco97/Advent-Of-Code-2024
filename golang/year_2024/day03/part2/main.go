package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"sort"
)

var multEnabled = false

type Position struct {
	start int
	end   int
}

type Token interface {
	position() Position
}

type DoToken struct {
	pos Position
}

func (o DoToken) position() Position {
	return o.pos
}

type DontToken struct {
	pos Position
}

func (o DontToken) position() Position {
	return o.pos
}

type MulToken struct {
	pos     Position
	factor1 int
	factor2 int
}

func (o MulToken) position() Position {
	return o.pos
}

var mul_operation_regex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var do_regex = regexp.MustCompile(`do\(\)`)
var dont_operation_regex = regexp.MustCompile(`don't\(\)`)

func extractMulTokens(text string) ([]MulToken, error) {
	matches := mul_operation_regex.FindAllStringSubmatchIndex(text, -1)
	result := make([]MulToken, len(matches))
	for i, match := range matches {
		match_start := match[0]
		match_end := match[1]
		// matched_text := text[match_start: match_end]
		// fmt.Println(matched_text)

		raw_factor1 := text[match[2]:match[3]]
		raw_factor2 := text[match[4]:match[5]]
		factor1, err := strconv.Atoi(raw_factor1)
		if err != nil {
			fmt.Printf("Errore durante la conversione da string a int del primo fattore: %v\n", err)
			return nil, err
		}
		factor2, err := strconv.Atoi(raw_factor2)
		if err != nil {
			fmt.Printf("Errore durante la conversione da string a int del secondo fattore: %v\n", err)
			return nil, err
		}
		result[i] = MulToken{
			pos: Position{
				start: match_start,
				end:   match_end,
			},
			factor1: factor1,
			factor2: factor2,
		}
	}
	return result, nil
}

func extractDoTokens(text string) ([]DoToken, error) {
	matches := do_regex.FindAllStringSubmatchIndex(text, -1)
	result := make([]DoToken, len(matches))
	for i, match := range matches {
		match_start := match[0]
		match_end := match[1]
		// matched_text := text[match_start: match_end]
		// fmt.Println(matched_text)

		result[i] = DoToken{
			pos: Position{
				start: match_start,
				end:   match_end,
			},
		}
	}
	return result, nil
}

func extractDontTokens(text string) ([]DontToken, error) {
	matches := dont_operation_regex.FindAllStringSubmatchIndex(text, -1)
	result := make([]DontToken, len(matches))
	for i, match := range matches {
		match_start := match[0]
		match_end := match[1]
		// matched_text := text[match_start: match_end]
		// fmt.Println(matched_text)

		result[i] = DontToken{
			pos: Position{
				start: match_start,
				end:   match_end,
			},
		}
	}
	return result, nil
}

func extractTokens(text string) ([]Token, error) {
	mulTokens, err := extractMulTokens(text)
	if err != nil {
		fmt.Printf("Errore durante l'estrazione delle operazioni mul: %v\n", err)
		return nil, err
	}
	doTokens, err := extractDoTokens(text)
	if err != nil {
		fmt.Printf("Errore durante l'estrazione delle operazioni do: %v\n", err)
		return nil, err
	}
	dontTokens, err := extractDontTokens(text)
	if err != nil {
		fmt.Printf("Errore durante l'estrazione delle operazioni dont: %v\n", err)
		return nil, err
	}
	var tokens []Token = make([]Token, len(mulTokens)+len(doTokens)+len(dontTokens))
	for i, mulToken := range mulTokens {
		tokens[i] = mulToken
	}
	for i, doToken := range doTokens {
		tokens[i+len(mulTokens)] = doToken
	}
	for i, dontToken := range dontTokens {
		tokens[i+len(mulTokens)+len(doTokens)] = dontToken
	}
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].position().start < tokens[j].position().start
	})
	return tokens, nil
}

func evalTokens(tokens []Token) (int, error) {
	sum := 0
	multEnabled = true
	for _, token := range tokens {
		switch token.(type) {
		case DoToken:
			multEnabled = true
		case DontToken:
			multEnabled = false
		case MulToken:
			mulToken, ok := token.(MulToken)
			if !ok {
				return -1, errors.New("errore durante il cast a MulToken")
			}
			if multEnabled {
				sum += mulToken.factor1 * mulToken.factor2
			}
		}
	}
	return sum, nil
}

func solve(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Errore durante l'apertura del file: %v\n", err)
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	file_content := ""
	for scanner.Scan() {
		line := scanner.Text()
		file_content += line
	}

	tokens, err := extractTokens(file_content)
	if err != nil {
		fmt.Printf("Errore durante l'estrazione dei tokens: %v\n", err)
		return -1, err
	}

	result, err := evalTokens(tokens)
	if err != nil {
		fmt.Printf("Errore durante l'evaluazione dei tokens: %v\n", err)
		return -1, err
	}
	return result, nil
}

func main() {
	content, _ := solve("../input.txt")
	fmt.Println(content)
}
