package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 3882564
	actual := SolvePart1(files.ReadFile(3))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 3385170
	actual := SolvePart2(files.ReadFile(3))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
