package day02

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day02/input.txt", func(line string) {
		input = append(input, line)
	})

	rounds := Parse(input)
	fmt.Printf("Part 1: %d\n", TotalStrategyPoints(rounds))

	roundsPart2 := ParsePart2(input)
	fmt.Printf("Part 2: %d\n", TotalStrategyPointsPart2(roundsPart2))
}

type Round struct {
	You      Shape
	Opponent Shape
}

type Shape int

const (
	Rock    = 1
	Paper   = 2
	Scissor = 3
)

func Parse(input []string) []Round {
	rounds := make([]Round, len(input))
	for i, v := range input {
		shapes := strings.Split(v, " ")
		r := Round{}
		switch shapes[0] {
		case "A":
			r.Opponent = Rock
		case "B":
			r.Opponent = Paper
		case "C":
			r.Opponent = Scissor
		}
		switch shapes[1] {
		case "X":
			r.You = Rock
		case "Y":
			r.You = Paper
		case "Z":
			r.You = Scissor
		}
		rounds[i] = r
	}
	return rounds
}

func EvaluateRound(you, opponent Shape) int {
	if you == opponent {
		return int(you) + 3
	}
	result := 0
	switch you {
	case Rock:
		if opponent == Scissor {
			result = 6
		}
	case Paper:
		if opponent == Rock {
			result = 6
		}
	case Scissor:
		if opponent == Paper {
			result = 6
		}
	}
	return int(you) + result
}

func TotalStrategyPoints(rounds []Round) int {
	total := 0
	for _, r := range rounds {
		total += EvaluateRound(r.You, r.Opponent)
	}
	return total
}

type Action int

const (
	Lose Action = iota
	Draw
	Win
)

type RoundPart2 struct {
	Opponent Shape
	Act      Action
}

func ParsePart2(input []string) []RoundPart2 {
	rounds := make([]RoundPart2, len(input))
	for i, v := range input {
		shapes := strings.Split(v, " ")
		r := RoundPart2{}
		switch shapes[0] {
		case "A":
			r.Opponent = Rock
		case "B":
			r.Opponent = Paper
		case "C":
			r.Opponent = Scissor
		}
		switch shapes[1] {
		case "X":
			r.Act = Lose
		case "Y":
			r.Act = Draw
		case "Z":
			r.Act = Win
		}
		rounds[i] = r
	}
	return rounds
}

func EvaluateRoundPart2(round RoundPart2) int {
	opponent := round.Opponent
	if round.Act == Draw {
		return int(opponent) + 3
	}

	result := 0
	switch round.Opponent {
	case Rock:
		result = int(Scissor)
		if round.Act == Win {
			result = int(Paper) + 6
		}
	case Paper:
		result = int(Rock)
		if round.Act == Win {
			result = int(Scissor) + 6
		}
	case Scissor:
		result = int(Paper)
		if round.Act == Win {
			result = int(Rock) + 6
		}
	}
	return result
}

func TotalStrategyPointsPart2(rounds []RoundPart2) int {
	total := 0
	for _, r := range rounds {
		total += EvaluateRoundPart2(r)
	}
	return total
}
