package day13_test

import (
	"testing"

	"github.com/umbe77/aoc-2022/day13"
)

var sampleInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestPart1(t *testing.T) {
	pairs := day13.Parse(sampleInput)
	r := day13.Part1(pairs)
	if r != 13 {
		t.Errorf("Expected 13, Got %d", r)
	}
}
