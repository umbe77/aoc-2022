package day02_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day02"
)

var input = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	rounds := day02.Parse(input)
	total := day02.TotalStrategyPoints(rounds)
	if total != 15 {
		t.Errorf("Expected 15, Got %d", total)
	}
}

func TestPart2(t *testing.T) {
	rounds := day02.ParsePart2(input)
	total := day02.TotalStrategyPointsPart2(rounds)
	if total != 12 {
		t.Errorf("Expected 12, Got %d", total)
	}
}
