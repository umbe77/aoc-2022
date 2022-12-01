package day01

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day01/input.txt", func(line string) {
		input = append(input, line)
	})

	elfs := ParseInput(input)
	fmt.Printf("Part 1: %d\n", Part1(elfs))
	fmt.Printf("Part 2: %d\n", Part2(elfs))

}

func ParseInput(input []string) [][]int {
	elfs := make([][]int, 0)
	currentIndex := -1
	for _, line := range input {
		if line == "" || currentIndex == -1 {
			currentIndex++
			elfs = append(elfs, make([]int, 0))
		}
		calories, _ := strconv.Atoi(line)
		elfs[currentIndex] = append(elfs[currentIndex], calories)
	}
	return elfs
}

func getCalories(elfs [][]int) []int {
	totalCalories := make([]int, len(elfs))
	for i, calories := range elfs {
		for _, cal := range calories {
			totalCalories[i] += cal
		}
	}

	sort.Ints(totalCalories)
	return totalCalories

}

func Part1(elfs [][]int) int {
	totalCalories := getCalories(elfs)
	return totalCalories[len(totalCalories)-1]
}

func Part2(elfs [][]int) int {
	totalCalories := getCalories(elfs)
	return totalCalories[len(totalCalories)-1] +
		totalCalories[len(totalCalories)-2] +
		totalCalories[len(totalCalories)-3]
}
