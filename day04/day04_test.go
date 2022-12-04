package day04_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day04"
)

var input = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	i := day04.Parse(input)
	r := day04.Part(i, day04.IsContained)
	if r != 2 {
		t.Errorf("Expected 2, Got %d", r)
	}
}
func TestPart2(t *testing.T) {
	i := day04.Parse(input)
	r := day04.Part(i, day04.IsOverlapped)
	if r != 4 {
		t.Errorf("Expected 4, Got %d", r)
	}
}
