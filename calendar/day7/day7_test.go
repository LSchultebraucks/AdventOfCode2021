package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 37
	actual := SolvePart1(files.ReadTestInputFile(7))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 355764
	actual := SolvePart1(files.ReadFile(7))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 168
	actual := SolvePart2(files.ReadTestInputFile(7))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 99634572
	actual := SolvePart2(files.ReadFile(7))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
