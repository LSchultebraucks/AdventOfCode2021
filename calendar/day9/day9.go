package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(9)
	println(SolvePart1(input))
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input))
	for i, line := range input {
		split := strings.Split(line, "")
		result[i] = make([]int, len(split))
		result[i] = arrays.StringArrayToIntArray(split)
	}
	return result
}

func SolvePart1(input []string) int {
	result := 0
	parsedInput := parseInput(input)
	risks := riskLevels(parsedInput)
	result = arrays.Sum(risks) + len(risks)
	return result
}

func SolvePart2(input []string) int {
	result := 0
	parsedInput := parseInput(input)
	println(parsedInput)
	return result
}

func riskLevels(input [][]int) []int {
	result := make([]int, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if i == 0 {
				if j == 0 {
					if input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] {
						result = append(result, input[i][j])
					}
				} else if j == len(input[i])-1 {
					if input[i][j] < input[i+1][j] && input[i][j] < input[i][j-1] {
						result = append(result, input[i][j])
					}
				} else {
					if input[i][j] < input[i][j-1] && input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] {
						result = append(result, input[i][j])
					}
				}
			} else {
				if j == 0 {
					if i == len(input)-1 {
						if input[i][j] < input[i][j+1] && input[i][j] < input[i-1][j] {
							result = append(result, input[i][j])
						}
					} else {
						if input[i][j] < input[i-1][j] && input[i][j] < input[i][j+1] && input[i][j] < input[i+1][j] {
							result = append(result, input[i][j])
						}
					}
				} else {
					if i == len(input)-1 {
						if j == len(input[i])-1 {
							if input[i][j] < input[i][j-1] && input[i][j] < input[i-1][j] {
								result = append(result, input[i][j])
							}
						} else {
							if input[i][j] < input[i][j-1] && input[i][j] < input[i-1][j] && input[i][j] < input[i][j+1] {
								result = append(result, input[i][j])
							}
						}
					} else {
						if j == len(input[i])-1 {
							if input[i][j] < input[i][j-1] && input[i][j] < input[i-1][j] && input[i][j] < input[i+1][j] {
								result = append(result, input[i][j])
							}
						} else {
							if input[i][j] < input[i-1][j] && input[i][j] < input[i][j-1] && input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] {
								result = append(result, input[i][j])
							}
						}
					}
				}
			}
		}
	}
	return result
}
