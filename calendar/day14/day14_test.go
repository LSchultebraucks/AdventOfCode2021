package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 1588
	actual := SolvePart1(files.ReadTestInputFile(14))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 2590
	actual := SolvePart1(files.ReadFile(14))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 2188189693529
	actual := SolvePart2(files.ReadTestInputFile(14))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 2875665202438
	actual := SolvePart2(files.ReadFile(14))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
