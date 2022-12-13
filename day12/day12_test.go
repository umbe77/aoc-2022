package day12_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day12"
)

var sampleInput = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestPart1(t *testing.T) {
	grid, width, height, start, end := day12.Parse(sampleInput)
	p := day12.Part1(grid, width, height, start, end)
	if p != 31 {
		t.Errorf("Expected 31, Got %d", p)
	}
}

func TestPart2(t *testing.T) {
	grid, width, height, _, end := day12.Parse(sampleInput)
	p := day12.Part2(grid, width, height, end)
	if p != 29 {
		t.Errorf("Expected 29, Got %d", p)
	}
}
