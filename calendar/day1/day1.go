package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
)

func main() {
	input := files.ReadFile(1)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	for i, s := range input {
		if i > 0 && conv.StringToNumber(s) > conv.StringToNumber(input[i-1]) {
			result++
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	for i := 3; i < len(input); i++ {
		sum2 := conv.StringToNumber(input[i]) + conv.StringToNumber(input[i-1]) + conv.StringToNumber(input[i-2])
		sum1 := conv.StringToNumber(input[i-1]) + conv.StringToNumber(input[i-2]) + conv.StringToNumber(input[i-3])
		if sum2 > sum1 {
			result++
		}
	}
	return result
}
