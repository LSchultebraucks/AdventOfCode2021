package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(6)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	return simulatePopulation(parseInput(input), 80)
}

func SolvePart2(input []string) int {
	return simulatePopulation(parseInput(input), 256)
}

func simulatePopulation(fish map[int]int, days int) int {
	for i := 0; i < days; i++ {
		fish = simulateDay(fish)
	}
	return mapSize(fish)
}

func simulateDay(fish map[int]int) map[int]int {
	oldFish := fish
	fish = make(map[int]int)
	for k, v := range oldFish {
		if k == 0 {
			fish[6] += v
			fish[8] += v
			continue
		}
		fish[k-1] += v
	}
	return fish
}

func mapSize(input map[int]int) int {
	cnt := 0
	for _, v := range input {
		cnt += v
	}
	return cnt
}

func parseInput(input []string) map[int]int {
	fish := make(map[int]int)
	splitted := strings.Split(input[0], ",")
	for _, s := range splitted {
		n := conv.StringToNumber(s)
		fish[n]++
	}
	return fish
}
