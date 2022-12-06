package day06

import (
	"fmt"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day06/input.txt", func(line string) {
		input = append(input, line)
	})
	fmt.Printf("Part 1: %d\n", Part1(input[0]))
	fmt.Printf("Part 2: %d\n", Part2(input[0]))
}

func Part1(signal string) int {
	for i := 3; i < len(signal); i++ {
		c1 := signal[i-3]
		c2 := signal[i-2]
		c3 := signal[i-1]
		c4 := signal[i]

		if c1 != c2 && c1 != c3 && c1 != c4 &&
			c2 != c3 && c2 != c4 &&
			c3 != c4 {
			return i + 1
		}
	}
	return 0
}

func Part2(signal string) int {

	for i := 13; i < len(signal); i++ {
		distinctValues := make(map[string]struct{})
		for j := i - 13; j <= i; j++ {
			distinctValues[string(signal[j])] = struct{}{}
		}
		if len(distinctValues) == 14 {
			return i + 1
		}
	}
	return 0
}
