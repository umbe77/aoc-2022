package day10_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day10"
	"github.com/umbe77/aoc-2022/utils"
)

func getInput() []string {

	input := make([]string, 0)
	utils.ReadFile("sample_input.txt", func(line string) {
		input = append(input, line)
	})
	return input
}

func TestPart1(t *testing.T) {
	operations := day10.Parse(getInput())
	st := day10.Part1(operations)
	if st != 13140 {
		t.Errorf("Expected 13140, Got %d", st)
	}
}
