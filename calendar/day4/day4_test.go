package main

import (
	"AdventOfCode2021/utils/files"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 38594
	actual := SolvePart1(files.ReadFile(4))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 21184
	actual := SolvePart2(files.ReadFile(4))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
