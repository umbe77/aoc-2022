package day08

import (
	"fmt"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day08/input.txt", func(line string) {
		input = append(input, line)
	})
	treeMap := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1_Mine(treeMap))
	fmt.Printf("Part 1: %d\n", Part1_Refactored(treeMap))
	fmt.Printf("Part 2: %d\n", Part2(treeMap))
}

func Parse(input []string) [][]int {
	treeMap := make([][]int, len(input))
	for y, v := range input {
		treeMap[y] = make([]int, len(v))
		for x, h := range v {
			treeMap[y][x] = utils.Atoi(string(h))
		}
	}
	return treeMap
}

func coordinates(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

func isVisible(treeMap [][]int, x, y int) bool {
	height := len(treeMap)
	width := len(treeMap[0])
	isVisibleRight := true
	for i := x + 1; i < width; i++ {
		if treeMap[y][x] <= treeMap[y][i] {
			isVisibleRight = false
			break
		}
	}
	isVisibleLeft := true
	for i := x - 1; i >= 0; i-- {
		if treeMap[y][x] <= treeMap[y][i] {
			isVisibleLeft = false
			break
		}
	}
	isVisibleDown := true
	for i := y + 1; i < height; i++ {
		if treeMap[y][x] <= treeMap[i][x] {
			isVisibleDown = false
			break
		}
	}
	isVisibleUp := true
	for i := y - 1; i >= 0; i-- {
		if treeMap[y][x] <= treeMap[i][x] {
			isVisibleUp = false
			break
		}
	}
	return isVisibleRight || isVisibleLeft || isVisibleDown || isVisibleUp
}

func Part1_Refactored(treeMap [][]int) int {
	visibleTrees := len(treeMap)*2 - 2 + len(treeMap[0])*2 - 2
	height := len(treeMap) - 1
	width := len(treeMap[0]) - 1
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if isVisible(treeMap, x, y) {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func Part1_Mine(treeMap [][]int) int {
	visibleTrees := make(map[string]struct{})
	width := len(treeMap[0]) - 1
	height := len(treeMap) - 1
	for col := 0; col <= width; col++ {
		visibleTrees[coordinates(col, 0)] = struct{}{}
		visibleTrees[coordinates(col, height)] = struct{}{}
	}
	for row := 1; row < height; row++ {
		visibleTrees[coordinates(0, row)] = struct{}{}
		visibleTrees[coordinates(width, row)] = struct{}{}
	}
	for col := 1; col < width; col++ {
		currentMaxHeight := 0
		for row := 1; row < height; row++ {
			currentMaxHeight = utils.Max(currentMaxHeight, treeMap[row-1][col])
			if currentMaxHeight < treeMap[row][col] {
				visibleTrees[coordinates(col, row)] = struct{}{}
			}
		}
	}
	for col := 1; col < width; col++ {
		currentMaxHeight := 0
		for row := height - 1; row >= 0; row-- {
			currentMaxHeight = utils.Max(currentMaxHeight, treeMap[row+1][col])
			if currentMaxHeight < treeMap[row][col] {
				visibleTrees[coordinates(col, row)] = struct{}{}
			}
		}
	}
	for row := 1; row < height; row++ {
		currentMaxWidth := 0
		for col := 1; col < width; col++ {
			currentMaxWidth = utils.Max(currentMaxWidth, treeMap[row][col-1])
			if currentMaxWidth < treeMap[row][col] {
				visibleTrees[coordinates(col, row)] = struct{}{}
			}
		}
	}
	for row := 1; row < height; row++ {
		currentMaxWidth := 0
		for col := width - 1; col >= 0; col-- {
			currentMaxWidth = utils.Max(currentMaxWidth, treeMap[row][col+1])
			if currentMaxWidth < treeMap[row][col] {
				visibleTrees[coordinates(col, row)] = struct{}{}
			}
		}
	}

	return len(visibleTrees)
}

type Dir int

const (
	Left Dir = iota
	Down
	Right
	Up
)

func Part2(treeMap [][]int) int {
	maxScenicScore := 0
	width := len(treeMap[0]) - 1
	height := len(treeMap) - 1
	for y, row := range treeMap {
		for x := range row {
			scenicScoreRight := 0
			for i := x + 1; i <= width; i++ {
				scenicScoreRight++
				if treeMap[y][x] <= treeMap[y][i] {
					break
				}
			}
			scenicScoreleft := 0
			for i := x - 1; i >= 0; i-- {
				scenicScoreleft++
				if treeMap[y][x] <= treeMap[y][i] {
					break
				}
			}
			scenicScoreDown := 0
			for i := y + 1; i <= height; i++ {
				scenicScoreDown++
				if treeMap[y][x] <= treeMap[i][x] {
					break
				}
			}
			scenicScoreUp := 0
			for i := y - 1; i >= 0; i-- {
				scenicScoreUp++
				if treeMap[y][x] <= treeMap[i][x] {
					break
				}
			}
			maxScenicScore = utils.Max(maxScenicScore, scenicScoreRight*scenicScoreleft*scenicScoreDown*scenicScoreUp)
		}
	}
	return maxScenicScore
}
