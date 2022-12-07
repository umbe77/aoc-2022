package day07_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day07"
)

var sampleInput = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestPart1(t *testing.T) {
	folders := day07.Parse(sampleInput)
	s := day07.Part1(folders)
	if s != 95437 {
		t.Errorf("Expected 95437, Got %d", s)
	}
}

func TestPart2(t *testing.T) {
	folders := day07.Parse(sampleInput)
	s := day07.Part2(folders)
	if s != 24933642 {
		t.Errorf("Expected 24933642, Got %d", s)
	}
}
