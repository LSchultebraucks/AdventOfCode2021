package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 15
	actual := SolvePart1(files.ReadTestInputFile(9))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 452
	actual := SolvePart1(files.ReadFile(9))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

/*
func TestSolvePart2WithExample(t *testing.T) {
	expected := 1134
	actual := SolvePart2(files.ReadTestInputFile(9))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 0
	actual := SolvePart2(files.ReadFile(9))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
*/
