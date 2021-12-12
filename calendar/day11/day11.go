package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
)

func main() {
	input := files.ReadFile(11)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart2(input []string) int {
	result := 0
	octopuses := parseInput(input)
	for i := 0; true; i++ {
		if simulateFlashStep(octopuses) == 100 {
			result = i + 1
			break
		}

	}
	return result
}

func SolvePart1(input []string) int {
	result := 0
	octopuses := parseInput(input)
	for i := 0; i < 100; i++ {
		result += simulateFlashStep(octopuses)
	}
	return result
}

func simulateFlashStep(octopuses [][]int) int {
	result := 0
	increaseEnergyLevelByOne(octopuses)
	result = octopusesFlashes(octopuses)
	return result
}

func octopusesFlashes(octopuses [][]int) int {
	result := 0
	flashed := generateHasFlashedFalse(octopuses)
	for {
		done := true
		for r, row := range octopuses {
			for c, v := range row {
				if !flashed[r][c] && v > 9 {
					done = false
					octopuses[r][c] = 0
					flashed[r][c] = true
					result++
					checkNeighbours(octopuses, r, c, flashed)
				}
			}
		}
		if done {
			break
		}
	}
	return result
}

func checkNeighbours(octopuses [][]int, r int, c int, flashed [][]bool) {
	for neighborR := r - 1; neighborR <= r+1; neighborR++ {
		for neighborC := c - 1; neighborC <= c+1; neighborC++ {
			if neighborR == r && neighborC == c {
				continue
			}
			incrementNeighbour(octopuses, neighborR, neighborC, flashed)
		}
	}
}

func incrementNeighbour(octopuses [][]int, r int, c int, flashed [][]bool) {
	if r < 0 || c < 0 || r >= 10 || c >= 10 || flashed[r][c] {
		return
	}
	octopuses[r][c]++
}

func generateHasFlashedFalse(octupuses [][]int) [][]bool {
	hasFlashed := make([][]bool, len(octupuses))
	for i := 0; i < len(octupuses); i++ {
		hasFlashed[i] = make([]bool, len(octupuses[i]))
		for j := 0; j < len(octupuses[i]); j++ {
			hasFlashed[i][j] = false
		}
	}
	return hasFlashed
}

func increaseEnergyLevelByOne(octopuses [][]int) {
	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[i]); j++ {
			octopuses[i][j]++
		}
	}
}

func parseInput(input []string) [][]int {
	var result = make([][]int, len(input))
	for i, line := range input {
		result[i] = make([]int, len(line))
		for j, char := range line {
			result[i][j] = conv.StringToNumber(string(char))
		}
	}
	return result
}
