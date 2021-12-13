package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1WithExample(t *testing.T) {
	expected := 17
	actual := SolvePart1(files.ReadTestInputFile(13))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 647
	actual := SolvePart1(files.ReadFile(13))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
