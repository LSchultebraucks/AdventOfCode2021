package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile(13)
	println(SolvePart1(input))
	SolvePart2(input)
}

type Coord struct {
	X, Y int
}
type Grid map[Coord]bool

func SolvePart1(input []string) int {
	result := 0
	grid, folds := parseInput(input)
	foldedGrid := doFolds(grid, folds, 1)
	result = countDots(foldedGrid)
	return result
}

func SolvePart2(input []string) {
	grid, folds := parseInput(input)
	foldedGrid := doFolds(grid, folds, len(folds))
	printGrid(foldedGrid)
}

func printGrid(grid Grid) {
	maxX := 0
	maxY := 0
	for coords := range grid {
		if coords.X > maxX {
			maxX = coords.X
		}
		if coords.Y > maxY {
			maxY = coords.Y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[Coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Printf(".")
			}
		}
		println()
	}
}

func doFolds(grid Grid, folds []string, foldCnt int) Grid {
	foldedGrid := grid
	for i := 0; i < foldCnt; i++ {
		foldStatement := folds[i]
		foldedGrid = performSingleFold(foldedGrid, foldStatement)
	}
	return foldedGrid
}

func performSingleFold(grid Grid, foldStatement string) Grid {
	foldedGrid := make(Grid)
	splitStatement := strings.Split(foldStatement, "=")
	offset := conv.StringToNumber(splitStatement[1])
	for coord := range grid {
		difference := 0
		if strings.Contains(splitStatement[0], "x") {
			difference = coord.X - offset
			if difference > 0 {
				coord.X = offset - difference
				foldedGrid[coord] = true
			} else if difference < 0 {
				foldedGrid[coord] = true
			}
		} else if strings.Contains(splitStatement[0], "y") {
			difference = coord.Y - offset
			if difference > 0 {
				coord.Y = offset - difference
				foldedGrid[coord] = true
			} else if difference < 0 {
				foldedGrid[coord] = true
			}
		}
	}
	return foldedGrid
}

func countDots(grid Grid) int {
	result := 0
	for coord := range grid {
		if grid[coord] {
			result++
		}
	}
	return result
}

func parseInput(input []string) (Grid, []string) {
	grid := make(Grid)
	folds := make([]string, len(input))
	cIdx := 0
	for i, line := range input {
		if line == "" {
			cIdx = i
			break
		}
		lineSplit := strings.Split(line, ",")
		x := conv.StringToNumber(lineSplit[0])
		y := conv.StringToNumber(lineSplit[1])
		grid[Coord{x, y}] = true
	}
	for i := cIdx; i < len(input); i++ {
		if input[i] != "" {
			folds[i-cIdx-1] = input[i]
		}
	}
	folds = folds[:len(input)-cIdx-1]
	return grid, folds
}
