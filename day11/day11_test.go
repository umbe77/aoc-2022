package day11_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day11"
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
	monkeys := day11.Parse(getInput())
	r := day11.Part1(monkeys)
	if r != 10605 {
		t.Errorf("Expected 10605, Got %d", r)
	}
}

// func TestPart2(t *testing.T) {
// 	monkeys := day11.Parse(getInput())
// 	r := day11.Part2(monkeys)
// 	if r != 2713310158 {
// 		t.Errorf("Expected 2713310158, Got %d", r)
// 	}
// }
