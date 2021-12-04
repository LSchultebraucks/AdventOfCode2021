package main

import (
	"AdventOfCode2021/utils/arrays"
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 1688
	actual := SolvePart1(arrays.StringArrayToIntArray(files.ReadFile(1)))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 1728
	actual := SolvePart2(arrays.StringArrayToIntArray(files.ReadFile(1)))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
