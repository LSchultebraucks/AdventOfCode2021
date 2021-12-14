package main

import (
	"AdventOfCode2021/utils/files"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(14)
	println(SolvePart1(input))
	println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	result := 0
	template, rules := parseInput(input)
	result = doSteps(template, rules, 10)
	return result
}

func SolvePart2(input []string) int {
	result := 0
	template, rules := parseInput(input)
	result = doSteps(template, rules, 40)
	return result
}

func doSteps(template string, rules map[string]string, stepCnt int) int {
	result := 0
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		couple := template[i : i+2]
		pairs[couple]++
	}
	for i := 0; i < stepCnt; i++ {
		newPairs := make(map[string]int)
		for couple, cnt := range pairs {
			middle := rules[couple]
			right := middle + string(couple[1])
			left := string(couple[0]) + middle
			if newPairs[left] == 0 {
				newPairs[left] = cnt
			} else {
				newPairs[left] += cnt
			}
			if newPairs[right] == 0 {
				newPairs[right] = cnt
			} else {
				newPairs[right] += cnt
			}
		}
		pairs = newPairs
	}
	cntMap := make(map[string]int)
	for couple, cnt := range pairs {
		cntMap[string(couple[0])] += cnt
		cntMap[string(couple[1])] += cnt
	}
	cntMap[string(template[0])]++
	cntMap[string(template[len(template)-1])]++
	mostCommon, leastCommon := 0, math.MaxInt64
	for _, cnt := range cntMap {
		if cnt > mostCommon {
			mostCommon = cnt
		}
		if cnt < leastCommon {
			leastCommon = cnt
		}
	}
	result = mostCommon/2 - leastCommon/2
	return result
}

func parseInput(input []string) (string, map[string]string) {
	template := input[0]
	rules := make(map[string]string)
	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], " -> ")
		rules[split[0]] = split[1]
	}
	return template, rules
}
