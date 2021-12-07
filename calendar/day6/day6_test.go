package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 5934
	actual := SolvePart1(files.ReadTestInputFile(6))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 365862
	actual := SolvePart1(files.ReadFile(6))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 26984457539
	actual := SolvePart2(files.ReadTestInputFile(6))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 1653250886439
	actual := SolvePart2(files.ReadFile(6))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
