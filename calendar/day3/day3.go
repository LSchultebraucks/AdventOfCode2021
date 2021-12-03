package main

import (
	"AdventOfCode2021/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(3)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	commonBinaryOne := make([]int, len(input[0]))
	gamma := ""
	epsilon := ""
	for _, s := range input {
		for ydx := range s {
			if s[ydx] == '1' {
				commonBinaryOne[ydx]++
			}
		}
	}
	halfSize := len(input) / 2
	for _, s := range commonBinaryOne {
		if s >= halfSize {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaNum, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilon, 2, 64)
	result = int(gammaNum * epsilonNum)
	return result
}

func solvePart2(input []string) int {
	result := 0
	result = determineFilter(input, oxygenFilter) * determineFilter(input, co2Filter)
	return result
}
func determineFilter(input []string, filter func(int, int) int) int {
	newInput := make([]string, len(input))
	copy(newInput, input)

	offset := 0

	for len(newInput) > 1 {
		left := make([]string, 0, len(newInput))
		target := filter(mostLeast(newInput, offset))

		for _, line := range newInput {
			if int(line[offset]) == target {
				left = append(left, line)
			}
		}

		newInput = left
		offset++
	}

	filterNum, _ := strconv.ParseInt(newInput[0], 2, 32)
	return int(filterNum)
}

func oxygenFilter(most, _ int) int {
	return most
}

func co2Filter(_, least int) int {
	return least
}

func mostLeast(input []string, idx int) (int, int) {
	zeroes := 0
	ones := 0

	for _, line := range input {
		if line[idx] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes > ones {
		return '0', '1'
	}

	return '1', '0'
}
