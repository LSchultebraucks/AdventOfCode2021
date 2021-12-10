package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(10)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func parseInput(input []string) [][]string {
	result := make([][]string, len(input))
	for i, line := range input {
		split := strings.Split(line, "")
		result[i] = make([]string, len(split))
		result[i] = split
	}
	return result
}

func SolvePart1(input []string) int {
	result := 0
	parsedInput := parseInput(input)
	startingBrackets := "([{<"
	endingBrackets := ")]}>"
	bracketsValueMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	startingEndingBracketsMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	for i, line := range parsedInput {
		stack := make([]string, 0)
		for _, char := range line {
			if strings.Contains(startingBrackets, char) {
				stack = append(stack, startingEndingBracketsMap[char])
			} else if strings.Contains(endingBrackets, char) {
				if len(stack) == 0 {
					panic("Expected nothing, but found " + char)
				} else if char == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					println(i, "Expected "+stack[len(stack)-1]+", but found "+char)
					result += bracketsValueMap[char]
					break
				}
			}
		}
	}
	return result
}

func SolvePart2(input []string) int {
	result := 0
	parsedInput := parseInput(input)
	startingBrackets := "([{<"
	endingBrackets := ")]}>"
	bracketsValueMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	startingEndingBracketsMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	scores := make([]int, 0)
	for i, line := range parsedInput {
		stack := make([]string, 0)
		isLineCorrupted := false
		for _, char := range line {
			if strings.Contains(startingBrackets, char) {
				stack = append(stack, startingEndingBracketsMap[char])
			} else if strings.Contains(endingBrackets, char) {
				if len(stack) == 0 {
					panic("Expected nothing, but found " + char)
				} else if char == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					println(i, "Expected "+stack[len(stack)-1]+", but found "+char)
					isLineCorrupted = true
					break
				}
			}
		}
		if !isLineCorrupted {
			stack = arrays.ReverseStrings(stack)
			lineScore := 0
			for _, char := range stack {
				lineScore = lineScore * 5
				lineScore += bracketsValueMap[char]
			}
			scores = append(scores, lineScore)
		}
	}
	sort.Ints(scores)
	result = arrays.MiddleScore(scores)
	return result
}
