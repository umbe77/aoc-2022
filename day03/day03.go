package day03

import (
	"fmt"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day03/input.txt", func(line string) {
		input = append(input, line)
	})

	rucksacks := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(rucksacks))
	fmt.Printf("Part 2: %d\n", Part2(rucksacks))
}

type Rucksack struct {
	Comp1 string
	Comp2 string
	All   string
}

func Parse(input []string) []Rucksack {
	result := make([]Rucksack, len(input))
	for i, v := range input {
		result[i] = Rucksack{
			Comp1: v[:(len(v) / 2)],
			Comp2: v[(len(v) / 2):],
			All:   v,
		}
	}
	return result
}

func prioritySum(items []rune) int {
	sum := 0
	for _, e := range items {
		a := int(e)
		if 97 <= a && a <= 97+26 {
			sum += a - 96
		} else {
			sum += a - 38
		}
	}
	return sum
}

func Part1(rucksacks []Rucksack) int {
	errors := make([]rune, len(rucksacks))
	for i, r := range rucksacks {
		for _, comp1 := range r.Comp1 {
			var foundError bool
			for _, comp2 := range r.Comp2 {
				if comp1 == comp2 {
					errors[i] = comp1
					foundError = true
					break
				}
			}
			if foundError {
				break
			}
		}
	}
	return prioritySum(errors)
}

func FindBadge(r1, r2, r3 Rucksack) rune {
	for _, badge1 := range r1.All {
		for _, badge2 := range r2.All {
			for _, badge3 := range r3.All {
				if badge1 == badge2 && badge2 == badge3 {
					return badge1
				}
			}
		}
	}
	return '0'
}

func Part2(rucksacks []Rucksack) int {
	badges := make([]rune, 0)
	for i := 0; i < len(rucksacks); {
		badges = append(badges, rune(FindBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])))
		i += 3
	}
	return prioritySum(badges)
}
