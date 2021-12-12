package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 226
	actual := SolvePart1(files.ReadTestInputFile(12))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 3298
	actual := SolvePart1(files.ReadFile(12))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2WithExample(t *testing.T) {
	expected := 3509
	actual := SolvePart2(files.ReadTestInputFile(12))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 93572
	actual := SolvePart2(files.ReadFile(12))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
