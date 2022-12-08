package day08_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day08"
)

var sampleInput = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestPart1(t *testing.T) {
	treeMap := day08.Parse(sampleInput)
	result := day08.Part1(treeMap)
	if result != 21 {
		t.Errorf("Expected 21, Got %d", result)
	}
}
func TestPart2(t *testing.T) {
	treeMap := day08.Parse(sampleInput)
	result := day08.Part2(treeMap)
	if result != 8 {
		t.Errorf("Expected 8, Got %d", result)
	}
}
