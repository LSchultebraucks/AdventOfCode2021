package main

import (
	"AdventOfCode2021/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(1)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	for i, s := range input {
		if i > 0 && stringToNumber(s) > stringToNumber(input[i-1]) {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	for i := 3; i < len(input); i++ {
		sum2 := stringToNumber(input[i]) + stringToNumber(input[i-1]) + stringToNumber(input[i-2])
		sum1 := stringToNumber(input[i-1]) + stringToNumber(input[i-2]) + stringToNumber(input[i-3])
		if sum2 > sum1 {
			result++
		}
	}
	return result
}

func stringToNumber(value string) int {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intVar
}
