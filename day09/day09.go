package day09

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

var sampleInput_multi = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day09/input.txt", func(line string) {
		input = append(input, line)
	})
	moves := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(moves))
	fmt.Printf("Part 2: %d\n", Part2(moves))
}

type Move struct {
	Dir  string
	Step int
}

func Parse(data []string) []Move {
	moves := make([]Move, len(data))
	for i, line := range data {
		tokens := strings.Split(line, " ")
		moves[i] = Move{
			Dir:  tokens[0],
			Step: utils.Atoi(tokens[1]),
		}
	}
	return moves
}

func coordinates(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

func moveHead(head map[string]int, m Move) map[string]int {
	next := map[string]int{"x": head["x"], "y": head["y"]}
	switch m.Dir {
	case "L":
		next["x"] = next["x"] - 1
		break
	case "R":
		next["x"] = next["x"] + 1
		break
	case "U":
		next["y"] = next["y"] + 1
		break
	case "D":
		next["y"] = next["y"] - 1
		break
	}
	return next
}

func isAdjacent(head, tail map[string]int) bool {
	isAdjacent := false
	for x := -1; x <= 1 && !isAdjacent; x++ {
		for y := -1; y <= 1 && !isAdjacent; y++ {
			isAdjacent = tail["x"]+x == head["x"] && tail["y"]+y == head["y"]
		}
	}
	return isAdjacent
}

func moveTail(head, tail map[string]int) map[string]int {
	next := map[string]int{"x": tail["x"], "y": tail["y"]}
	if next["y"] == head["y"] {
		if next["x"] < head["x"] {
			next["x"] = next["x"] + 1
		} else {
			next["x"] = next["x"] - 1
		}
	} else if next["x"] == head["x"] {
		if next["y"] < head["y"] {
			next["y"] = next["y"] + 1
		} else {
			next["y"] = next["y"] - 1
		}
	} else {
		dx := head["x"] - next["x"]
		dy := head["y"] - next["y"]
		if head["x"] > next["x"] && head["y"] > next["y"] {
			if utils.Abs(dx) > utils.Abs(dy) {
				next["x"] = head["x"] - 1
				next["y"] = head["y"]
			} else {
				next["x"] = head["x"]
				next["y"] = head["y"] - 1
			}
		}
		if head["x"] > next["x"] && head["y"] < next["y"] {
			if utils.Abs(dx) > utils.Abs(dy) {
				next["x"] = head["x"] - 1
				next["y"] = head["y"]
			} else {
				next["x"] = head["x"]
				next["y"] = head["y"] + 1
			}
		}
		if head["x"] < next["x"] && head["y"] < next["y"] {
			if utils.Abs(dx) > utils.Abs(dy) {
				next["x"] = head["x"] + 1
				next["y"] = head["y"]
			} else {
				next["x"] = head["x"]
				next["y"] = head["y"] + 1
			}
		}
		if head["x"] < next["x"] && head["y"] > next["y"] {
			if utils.Abs(dx) > utils.Abs(dy) {
				next["x"] = head["x"] + 1
				next["y"] = head["y"]
			} else {
				next["x"] = head["x"]
				next["y"] = head["y"] - 1
			}
		}
	}

	return next
}

func printLocations(locations map[string]struct{}, upperX, upperY, lowerX, lowerY int) {
	for y := upperY; y >= lowerY; y-- {
		for x := lowerX; x <= upperX; x++ {
			char := "."
			if _, ok := locations[coordinates(x, y)]; ok {
				char = "#"
			}
			if x == 0 && y == 0 {
				char = "s"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func Part1(moves []Move) int {
	maxY := 0
	minY := 0
	maxX := 0
	minX := 0
	positions := make(map[string]struct{})
	currentHead := map[string]int{"x": 0, "y": 0}
	currentTail := map[string]int{"x": 0, "y": 0}
	positions[coordinates(currentTail["x"], currentTail["y"])] = struct{}{}
	// fmt.Printf("HEAD %+v\n", currentHead)
	for _, m := range moves {
		for i := 0; i < m.Step; i++ {
			currentHead = moveHead(currentHead, m)
			// fmt.Printf("(%+v)HEAD %+v\n", m, currentHead)
			if !isAdjacent(currentHead, currentTail) {
				currentTail = moveTail(currentHead, currentTail)
				// fmt.Printf("(%+v)TAIL %+v\n", m, currentTail)
				maxY = utils.Max(maxY, currentTail["y"])
				maxX = utils.Max(maxX, currentTail["x"])
				minY = utils.Min(minY, currentTail["y"])
				minX = utils.Min(minX, currentTail["x"])
				positions[coordinates(currentTail["x"], currentTail["y"])] = struct{}{}
			}
		}
		// fmt.Println("------------")
	}
	// printLocations(positions, maxX, maxY, minX, minY)
	return len(positions)
}

func Part2(moves []Move) int {
	maxY := 0
	minY := 0
	maxX := 0
	minX := 0
	positions := make(map[string]struct{})
	currentHead := map[string]int{"x": 0, "y": 0}
	currentTails := []map[string]int{
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
		{"x": 0, "y": 0},
	}
	positions[coordinates(0, 0)] = struct{}{}
	// fmt.Printf("HEAD %+v\n", currentHead)
	for _, m := range moves {
		for i := 0; i < m.Step; i++ {
			currentHead = moveHead(currentHead, m)
			// fmt.Printf("(%+v)HEAD %+v\n", m, currentHead)
			for tailIndex := range currentTails {
				head := currentHead
				if tailIndex > 0 {
					head = currentTails[tailIndex-1]
				}
				// fmt.Printf("Current %+v\n", currentTail)
				if !isAdjacent(head, currentTails[tailIndex]) {
					currentTails[tailIndex] = moveTail(head, currentTails[tailIndex])
					// fmt.Printf("(%+v)TAIL[%d] %+v\n", m, tailIndex, currentTail)
					if tailIndex == 8 {
						maxY = utils.Max(maxY, currentTails[tailIndex]["y"])
						maxX = utils.Max(maxX, currentTails[tailIndex]["x"])
						minY = utils.Min(minY, currentTails[tailIndex]["y"])
						minX = utils.Min(minX, currentTails[tailIndex]["x"])
						positions[coordinates(currentTails[tailIndex]["x"], currentTails[tailIndex]["y"])] = struct{}{}
					}
				}
			}
		}
		// fmt.Println("------------")
	}
	// printLocations(positions, maxX, maxY, minX, minY)
	return len(positions)
}
