package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 1656
	actual := SolvePart1(files.ReadTestInputFile(11))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 1667
	actual := SolvePart1(files.ReadFile(11))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 195
	actual := SolvePart2(files.ReadTestInputFile(11))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 488
	actual := SolvePart2(files.ReadFile(11))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
