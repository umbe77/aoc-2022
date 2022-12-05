package day05_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day05"
)

var sample = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	crates, moves := day05.Parse(sample)
	result := day05.Part1(crates, moves)
	if result != "CMZ" {
		t.Errorf("Expected 'CMZ', Got %s", result)
	}
}

func TestPart2(t *testing.T) {
	crates, moves := day05.Parse(sample)
	result := day05.Part2(crates, moves)
	if result != "MCD" {
		t.Errorf("Expected 'MCD', Got %s", result)
	}
}
