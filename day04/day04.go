package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day04/input.txt", func(line string) {
		input = append(input, line)
	})

	pairs := Parse(input)
	fmt.Printf("Part 1: %d\n", Part(pairs, IsContained))
	fmt.Printf("Part 2: %d\n", Part(pairs, IsOverlapped))

}

type CleaningRange struct {
	Start int
	End   int
}

type ElfPair struct {
	Elf1 CleaningRange
	Elf2 CleaningRange
}

func Parse(input []string) []ElfPair {
	result := make([]ElfPair, 0)
	for _, v := range input {
		ranges := strings.Split(v, ",")
		p1 := strings.Split(ranges[0], "-")
		p1Start, _ := strconv.Atoi(p1[0])
		p1End, _ := strconv.Atoi(p1[1])
		p2 := strings.Split(ranges[1], "-")
		p2Start, _ := strconv.Atoi(p2[0])
		p2End, _ := strconv.Atoi(p2[1])
		result = append(result, ElfPair{
			Elf1: CleaningRange{
				Start: p1Start,
				End:   p1End,
			},
			Elf2: CleaningRange{
				Start: p2Start,
				End:   p2End,
			},
		})
	}
	return result
}

func IsContained(r1, r2 CleaningRange) bool {
	return r1.Start >= r2.Start && r1.End <= r2.End
}

func IsOverlapped(r1, r2 CleaningRange) bool {
	return r1.End >= r2.Start && r1.End <= r2.End || r1.Start >= r2.Start && r1.Start <= r2.End
}

func Part(input []ElfPair, check func(r1, r2 CleaningRange) bool) int {
	count := 0

	for _, v := range input {
		if check(v.Elf1, v.Elf2) || check(v.Elf2, v.Elf1) {
			count++
		}
	}

	return count
}
