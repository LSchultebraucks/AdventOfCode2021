package main

import (
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(8)
	println(SolvePart1(input))
}

func parseInput(input []string) [][]string {
	result := make([][]string, len(input))
	for i, line := range input {
		splitDelimiter := strings.Split(line, " | ")
		signalPattern := strings.Split(splitDelimiter[0], " ")
		outputValues := strings.Split(splitDelimiter[1], " ")
		result[i] = append(signalPattern, outputValues...)
	}
	return result
}

func SolvePart1(input []string) int {
	result := 0
	parsedInput := parseInput(input)
	for _, line := range parsedInput {
		for i := 10; i <= 13; i++ {
			outputDigit := line[i]
			if len(outputDigit) == 2 || len(outputDigit) == 3 || len(outputDigit) == 4 || len(outputDigit) == 7 {
				result++
			}
		}
	}
	return result
}
