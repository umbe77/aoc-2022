package day10

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day10/input.txt", func(line string) {
		input = append(input, line)
	})
	operations := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(operations))
	fmt.Println("Part 2")
	Part2(operations)
	// inputSample := make([]string, 0)
	// utils.ReadFile("day10/sample_input.txt", func(line string) {
	// 	inputSample = append(inputSample, line)
	// })
	// operationsSample := Parse(inputSample)
	// fmt.Println("Part 2 Sample")
	// Part2(operationsSample)
}

type Operation struct {
	Op    string
	Param int
}

func Parse(input []string) []Operation {
	result := make([]Operation, len(input))
	for i, line := range input {
		tokens := strings.Split(line, " ")
		op := Operation{Op: tokens[0]}
		if tokens[0] == "addx" {
			op.Param = utils.Atoi(tokens[1])
		}
		result[i] = op
	}
	return result
}

func Part1(operations []Operation) int {
	X := 1
	signalStrength := 0
	cycle := 0
	for _, op := range operations {
		cycleCount := 1
		tmpX := 0
		switch op.Op {
		case "addx":
			cycleCount = 2
			tmpX = op.Param
			break
		}
		for i := 1; i <= cycleCount; i++ {
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength += (X * cycle)
			}
		}
		X = X + tmpX
	}
	return signalStrength
}

func isSpriteHere(i int, sprite []int) bool {
	for _, v := range sprite {
		if i == v {
			return true
		}
	}
	return false
}

func Part2(operations []Operation) {
	X := 1
	cycle := 0
	sprite := []int{0, 1, 2}
	screen := [][]string{make([]string, 40), make([]string, 40), make([]string, 40), make([]string, 40), make([]string, 40), make([]string, 40)}

	for _, op := range operations {
		cycleCount := 1
		tmpX := 0
		switch op.Op {
		case "addx":
			cycleCount = 2
			tmpX = op.Param
			break
		}
		for i := 1; i <= cycleCount; i++ {
			cycle++
			rowIndex := (cycle / 40)
			colNumber := cycle % 40
			if colNumber == 0 {
				colNumber = 40
				rowIndex--
			}
			colIndex := colNumber - 1

			if isSpriteHere(colIndex, sprite) {
				screen[rowIndex][colIndex] = "#"
			} else {
				screen[rowIndex][colIndex] = "."
			}
		}
		X = X + tmpX
		sprite[0] = X - 1
		sprite[1] = X
		sprite[2] = X + 1
	}

	for _, row := range screen {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
}
