package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(5)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	result := 0
	coordinatePairs := parseInputToCoordinatePairs(input)
	area := createArea()
	for _, pair := range coordinatePairs {
		area = markArea(area, pair)
	}
	result = countOverlappingArea(area, 2)
	return result
}

func SolvePart2(input []string) int {
	result := 0
	coordinatePairs := parseInputToCoordinatePairs(input)
	area := createArea()
	for _, pair := range coordinatePairs {
		area = markAreaWithDiagonal(area, pair)
	}
	result = countOverlappingArea(area, 2)
	return result
}

func countOverlappingArea(area [][]int, overlappedByAtLeast int) int {
	cnt := 0
	for _, row := range area {
		for _, cell := range row {
			if cell >= overlappedByAtLeast {
				cnt++
			}
		}
	}
	return cnt
}

func createArea() [][]int {
	area := make([][]int, 1000)
	for i := range area {
		area[i] = make([]int, 1000)
	}
	return area
}

func markArea(area [][]int, pair coordinatePair) [][]int {
	if pair.start.x == pair.end.x {
		start := 0
		end := 0
		if pair.start.y < pair.end.y {
			start = pair.start.y
			end = pair.end.y
		} else {
			start = pair.end.y
			end = pair.start.y
		}
		for y := start; y <= end; y++ {
			area[y][pair.start.x]++
		}
	} else if pair.start.y == pair.end.y {
		start := 0
		end := 0
		if pair.start.x < pair.end.x {
			start = pair.start.x
			end = pair.end.x
		} else {
			start = pair.end.x
			end = pair.start.x
		}
		for x := start; x <= end; x++ {
			area[pair.start.y][x]++
		}
	}
	return area
}

func markAreaWithDiagonal(area [][]int, pair coordinatePair) [][]int {
	if pair.start.x == pair.end.x || pair.start.y == pair.end.y {
		return markArea(area, pair)
	} else {
		if pair.start.x < pair.end.x {
			yOffset := 1
			if pair.start.y > pair.end.y {
				yOffset = -1
			}
			y := pair.start.y
			for x := pair.start.x; x <= pair.end.x; x++ {
				area[y][x]++
				y += yOffset
			}
		} else {
			yOffset := 1
			if pair.start.y > pair.end.y {
				yOffset = -1
			}
			y := pair.start.y
			for x := pair.start.x; x >= pair.end.x; x-- {
				area[y][x]++
				y += yOffset
			}
		}
	}
	return area
}

func parseInputToCoordinatePairs(input []string) []coordinatePair {
	coordinatePairs := make([]coordinatePair, len(input))
	for i, line := range input {
		coordinatesStart := strings.Split(strings.Split(line, " -> ")[0], ",")
		coordinatesEnd := strings.Split(strings.Split(line, " -> ")[1], ",")
		coordinatePairs[i] = newCoordinatePair(
			newCoordinates(conv.StringToNumber(coordinatesStart[0]), conv.StringToNumber(coordinatesStart[1])),
			newCoordinates(conv.StringToNumber(coordinatesEnd[0]), conv.StringToNumber(coordinatesEnd[1])))
	}
	return coordinatePairs
}

type coordinates struct {
	x int
	y int
}

func newCoordinates(x int, y int) coordinates {
	return coordinates{x: x, y: y}
}

type coordinatePair struct {
	start coordinates
	end   coordinates
}

func newCoordinatePair(start coordinates, end coordinates) coordinatePair {
	return coordinatePair{start: start, end: end}
}
