package main

import (
	"AdventOfCode2021/utils/files"
	"AdventOfCode2021/utils/maps"
	"strings"
	"unicode"
)

func main() {
	input := files.ReadFile(12)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	result := 0
	nodes := parseInput(input)
	start := "start"
	visits := make(map[string]int)
	result = calculatePathsVisitedOnce(nodes, start, visits)
	return result
}

func SolvePart2(input []string) int {
	result := 0
	nodes := parseInput(input)
	start := "start"
	visits := make(map[string]int)
	result = calculatePathVisitTwice(nodes, start, visits, false)
	return result
}

func calculatePathsVisitedOnce(nodes map[string][]string, currentNode string, visits map[string]int) int {
	end := "end"
	visits[currentNode]++
	result := 0
	for _, c := range nodes[currentNode] {
		if c == end {
			result++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || visits[c] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = maps.CopyMap(visits)
			}
			result += calculatePathsVisitedOnce(nodes, c, nextVisits)
		}
	}
	return result
}

func calculatePathVisitTwice(nodes map[string][]string, currentNode string, visits map[string]int, hasVisitedTwice bool) int {
	end := "end"
	result := 0
	if visits[currentNode] > 0 && unicode.IsLower(rune(currentNode[0])) {
		hasVisitedTwice = true
	}
	visits[currentNode]++
	for _, c := range nodes[currentNode] {
		if c == end {
			result++
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = maps.CopyMap(visits)
			}
			result += calculatePathVisitTwice(nodes, c, nextVisits, hasVisitedTwice)
		}
	}
	return result
}

func parseInput(input []string) map[string][]string {
	nodes := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, "-")
		// undirected graph
		nodes[parts[0]] = append(nodes[parts[0]], parts[1])
		nodes[parts[1]] = append(nodes[parts[1]], parts[0])
	}
	return nodes
}
