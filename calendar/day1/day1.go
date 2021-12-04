package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
)

func main() {
	input := files.ReadFile(1)
	numbers := arrays.StringArrayToIntArray(input)
	println(SolvePart1(numbers))
	println(SolvePart2(numbers))
}

func SolvePart1(input []int) int {
	result := 0
	for i, s := range input {
		if i > 0 && s > input[i-1] {
			result++
		}
	}

	return result
}

func SolvePart2(input []int) int {
	result := 0
	for i := 3; i < len(input); i++ {
		sum2 := arrays.Sum(input[i-2 : i+1])
		sum1 := arrays.Sum(input[i-3 : i])
		if sum2 > sum1 {
			result++
		}
	}
	return result
}
