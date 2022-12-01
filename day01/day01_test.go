package day01_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day01"
)

var input = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestPart1(t *testing.T) {

	elfs := day01.ParseInput(input)
	result := day01.Part1(elfs)

	if result != 24000 {
		t.Errorf("Exptected 24000, Got: %d", result)
	}

}

func TestPart2(t *testing.T) {

	elfs := day01.ParseInput(input)
	result := day01.Part2(elfs)

	if result != 45000 {
		t.Errorf("Exptected 45000, Got: %d", result)
	}

}
