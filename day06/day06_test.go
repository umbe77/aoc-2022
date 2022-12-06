package day06_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day06"
)

func TestPart1(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	r := day06.Part1(input)
	if r != 7 {
		t.Errorf("Exptectd 7, Got %d", r)
	}
	input = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	r = day06.Part1(input)
	if r != 5 {
		t.Errorf("Exptectd 5, Got %d", r)
	}
	input = "nppdvjthqldpwncqszvftbrmjlhg"
	r = day06.Part1(input)
	if r != 6 {
		t.Errorf("Exptectd 6, Got %d", r)
	}
	input = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	r = day06.Part1(input)
	if r != 10 {
		t.Errorf("Exptectd 10, Got %d", r)
	}
	input = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	r = day06.Part1(input)
	if r != 11 {
		t.Errorf("Exptectd 11, Got %d", r)
	}
}

func TestPart2(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	r := day06.Part2(input)
	if r != 19 {
		t.Errorf("Exptectd 19, Got %d", r)
	}
	input = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	r = day06.Part2(input)
	if r != 23 {
		t.Errorf("Exptectd 23, Got %d", r)
	}
	input = "nppdvjthqldpwncqszvftbrmjlhg"
	r = day06.Part2(input)
	if r != 23 {
		t.Errorf("Exptectd 23, Got %d", r)
	}
	input = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	r = day06.Part2(input)
	if r != 29 {
		t.Errorf("Exptectd 29, Got %d", r)
	}
	input = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	r = day06.Part2(input)
	if r != 26 {
		t.Errorf("Exptectd 26, Got %d", r)
	}
}
