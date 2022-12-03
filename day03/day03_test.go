package day03_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day03"
)

var input = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPart1(t *testing.T) {
	r := day03.Parse(input)
	sum := day03.Part1(r)
	if sum != 157 {
		t.Errorf("Expected 157, Got %d", sum)
	}
}

func TestPart2(t *testing.T) {
	r := day03.Parse(input)
	sum := day03.Part2(r)
	if sum != 70 {
		t.Errorf("Expected 70, Got %d", sum)
	}
}
