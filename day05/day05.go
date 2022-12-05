package day05

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day05/input.txt", func(line string) {
		input = append(input, line)
	})
	crates1, moves1 := Parse(input)
	fmt.Printf("Part 1: %s\n", Part1(crates1, moves1))
	crates2, moves2 := Parse(input)
	fmt.Printf("Part 2: %s\n", Part2(crates2, moves2))
}

type Move struct {
	Count int
	From  int
	To    int
}

func Part1(crates map[int]utils.Stack[string], moves []Move) string {
	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			cFrom := crates[m.From]
			cTo := crates[m.To]
			cTo.Push(cFrom.Pop())
			crates[m.From] = cFrom
			crates[m.To] = cTo
		}
	}
	result := ""
	for i := 1; i <= len(crates); i++ {
		c := crates[i]
		result = fmt.Sprintf("%s%s", result, c.Peek())
	}

	return result
}

func Part2(crates map[int]utils.Stack[string], moves []Move) string {
	for _, m := range moves {
		cFrom := crates[m.From]
		cTo := crates[m.To]
		cTo.PushN(cFrom.PopN(m.Count))
		crates[m.From] = cFrom
		crates[m.To] = cTo
	}
	result := ""
	for i := 1; i <= len(crates); i++ {
		c := crates[i]
		result = fmt.Sprintf("%s%s", result, c.Peek())
	}

	return result
}

func Parse(data []string) (map[int]utils.Stack[string], []Move) {
	crates := make(map[int]utils.Stack[string])
	moves := make([]Move, 0)

	startMovesIndex := 0
	for i, line := range data {
		if line == "" {
			startMovesIndex = i
			break
		}
	}
	//parse crates
	numberOfCrates := 0
	for i := startMovesIndex - 1; i >= 0; i-- {
		if i == startMovesIndex-1 {
			for _, v := range strings.Split(data[i], "  ") {
				crates[utils.Atoi(strings.Trim(v, " "))] = utils.New[string]()
			}
			numberOfCrates = len(crates)
			continue
		}
		line := data[i]
		currentLineIndex := 0
		for j := 1; j <= numberOfCrates; j++ {
			if j == 1 {
				currentLineIndex = 1
			}
			crateValue := string(line[currentLineIndex])
			if crateValue != " " {
				c := crates[j]
				c.Push(crateValue)
				crates[j] = c
			}
			currentLineIndex += 4
		}
	}
	//parse moves
	for i := startMovesIndex + 1; i < len(data); i++ {
		mLine := strings.Split(data[i], " ")
		move := Move{}
		for j := 0; j < len(mLine); j += 2 {
			switch mLine[j] {
			case "move":
				move.Count = utils.Atoi(mLine[j+1])
				break
			case "from":
				move.From = utils.Atoi(mLine[j+1])
			case "to":
				move.To = utils.Atoi(mLine[j+1])
			}
		}
		moves = append(moves, move)
	}

	return crates, moves

}
