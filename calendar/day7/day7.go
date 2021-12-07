package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(7)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	positions := parseInput(input)
	return determineFuel(positions, dist)
}

func SolvePart2(input []string) int {
	positions := parseInput(input)
	return determineFuel(positions, dist2)
}

func determineFuel(positions []int, distFunc func(int, int) int) int {
	maxPositions := arrays.MaxValue(positions)
	fuelMap := make(map[int]int, maxPositions)
	for pos := 0; pos < maxPositions; pos++ {
		if fuelMap[pos] == 0 {
			for _, pos2 := range positions {
				fuelMap[pos] += distFunc(pos, pos2)
			}
		}
	}
	return getMinFuel(fuelMap)
}

func getMinFuel(fuel map[int]int) int {
	min := math.MaxInt32
	for _, v := range fuel {
		if v < min {
			min = v
		}
	}
	return min
}

func parseInput(input []string) []int {
	split := strings.Split(input[0], ",")
	result := make([]int, len(split))
	for i, line := range split {
		result[i] = conv.StringToNumber(line)
	}
	return result
}

func dist(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func dist2(a, b int) int {
	result := 0
	distance := dist(a, b)
	velocity := 1
	for i := 0; i < distance; i++ {
		result += velocity
		velocity++
	}
	return result
}
