package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"math"
)

func main() {
	input := files.ReadFile(15)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

type Coord struct {
	Y, X int
}

type Cavern map[Coord]int

func SolvePart1(input []string) int {
	result := 0
	cavern := parseInput(input)
	width := len(input[0])
	height := width // input is square
	start := Coord{Y: 0, X: 0}
	end := Coord{Y: height - 1, X: width - 1}
	result = lowestRisk(cavern, &start, &end)
	return result
}

func SolvePart2(input []string) int {
	result := 0
	cavern := parseInput(input)
	width := len(input[0])
	height := width // input is square
	start := Coord{Y: 0, X: 0}
	end := Coord{Y: height*5 - 1, X: width*5 - 1}
	bigCavern := make(Cavern)
	for k, v := range cavern {
		for r := 0; r < 5; r++ {
			for c := 0; c < 5; c++ {
				increase := r + c
				value := 1 + (v+increase-1)%9
				bigCavern[Coord{k.Y + r*height, k.X + c*width}] = value
			}
		}
	}
	result = lowestRisk(bigCavern, &start, &end)
	return result
}

func lowestRisk(cavern Cavern, start *Coord, end *Coord) int {
	gScore := Cavern{
		*start: 0,
	}
	fScore := Cavern{
		*start: start.distance(end),
	}
	coordMap := map[Coord]bool{
		*start: true,
	}
	historyMap := make(map[Coord]Coord)

	for len(coordMap) != 0 {
		currentScore := math.MaxInt
		var current Coord
		for v := range coordMap {
			if score, ok := fScore[v]; ok {
				if score < currentScore {
					current = v
					currentScore = score
				}
			}
		}
		delete(coordMap, current)

		if current == *end {
			score := 0
			for current != *start {
				score += cavern[current]
				current = historyMap[current]
			}
			return score
		}
		for _, n := range cavern.Neighbors(current) {
			possibleScore := gScore[current] + cavern[n]
			if previousScore, ok := gScore[n]; !ok || possibleScore < previousScore {
				historyMap[n] = current
				gScore[n] = possibleScore
				fScore[n] = possibleScore + n.distance(end)
				if !coordMap[n] {
					coordMap[n] = true
				}
			}
		}
	}
	panic("No path found")
	return -1
}

func (c Coord) distance(dst *Coord) int {
	return (dst.Y - c.X) + (dst.Y - c.X)
}

func (r Cavern) Neighbors(pos Coord) []Coord {
	var coords []Coord
	up := Coord{pos.Y - 1, pos.X}
	down := Coord{pos.Y + 1, pos.X}
	left := Coord{pos.Y, pos.X - 1}
	right := Coord{pos.Y, pos.X + 1}
	for _, v := range []Coord{up, down, left, right} {
		if _, ok := r[v]; ok {
			coords = append(coords, v)
		}
	}
	return coords
}

func parseInput(input []string) Cavern {
	cavern := make(Cavern, 0)
	for y, line := range input {
		for x, c := range line {
			cavern[Coord{y, x}] = conv.StringToNumber(string(c))
		}
	}
	return cavern
}
