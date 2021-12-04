package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 1561344
	actual := SolvePart1(files.ReadFile(2))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 1848454425
	actual := SolvePart2(files.ReadFile(2))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
