package main

import (
	"AdventOfCode2021/utils/conv"
	"AdventOfCode2021/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(2)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	horizontal := 0
	depth := 0
	for _, s := range input {
		command := strings.Split(s, " ")
		direction := command[0]
		amount := conv.StringToNumber(command[1])
		if direction == "forward" {
			horizontal += amount
		} else if direction == "down" {
			depth += amount
		} else if direction == "up" {
			depth -= amount
		}
	}
	result = horizontal * depth
	return result
}

func solvePart2(input []string) int {
	result := 0
	horizontal := 0
	depth := 0
	aim := 0
	for _, s := range input {
		command := strings.Split(s, " ")
		direction := command[0]
		amount := conv.StringToNumber(command[1])
		if direction == "forward" {
			horizontal += amount
			depth += aim * amount
		} else if direction == "down" {
			aim += amount
		} else if direction == "up" {
			aim -= amount
		}
	}
	result = horizontal * depth
	return result
}
