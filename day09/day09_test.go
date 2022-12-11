package day09_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day09"
)

var sampleInput = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

func TestPart1(t *testing.T) {
	moves := day09.Parse(sampleInput)
	locations := day09.Part1(moves)
	if locations != 13 {
		t.Errorf("Expected 13, Got %d", locations)
	}
}

var sampleInput_multi = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestPart2(t *testing.T) {
	moves := day09.Parse(sampleInput_multi)
	locations := day09.Part2(moves)
	if locations != 36 {
		t.Errorf("Expected 36, Got %d", locations)
	}
}
