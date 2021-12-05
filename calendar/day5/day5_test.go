package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 5
	actual := SolvePart1(files.ReadTestInputFile(5))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 7318
	actual := SolvePart1(files.ReadFile(5))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 12
	actual := SolvePart2(files.ReadTestInputFile(5))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 19939
	actual := SolvePart2(files.ReadFile(5))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
