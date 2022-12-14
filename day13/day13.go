package day13

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := utils.ReadAllFile("day13/input.txt")
	pairs := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(pairs))
}

type Pair struct {
	Left  []interface{}
	Right []interface{}
}

func Parse(input string) []Pair {
	pairs := make([]Pair, 0)

	for _, p := range strings.Split(input, "\n\n") {
		var l []interface{}
		var r []interface{}
		packets := strings.Split(p, "\n")

		json.Unmarshal([]byte(packets[0]), &l)
		json.Unmarshal([]byte(packets[1]), &r)

		pairs = append(pairs, Pair{
			Left:  l,
			Right: r,
		})
	}

	return pairs
}

func isFloat64(a interface{}) bool {
	if _, ok := a.(float64); ok {
		return true
	}
	return false
}

func isArray(a interface{}) bool {
	if _, ok := a.([]interface{}); ok {
		return true
	}
	return false
}

func isValid(left, right interface{}) (bool, int) {

	if isFloat64(left) && isFloat64(right) {
		if left.(float64) < right.(float64) {
			return true, 1
		}
		if left.(float64) > right.(float64) {
			return false, 1
		}
	} else if isArray(left) && isArray(right) {
		left_len := len(left.([]interface{}))
		right_len := len(right.([]interface{}))
		bound := utils.Min(left_len, right_len)
		for i := 0; i < bound; i++ {
			valid, ok := isValid(left.([]interface{})[i], right.([]interface{})[i])
			if ok != 0 {
				return valid, 1
			}
		}
		if left_len < right_len {
			return true, 1
		}
		if left_len > right_len {
			return false, 1
		}
	} else if isFloat64(left) && isArray(right) {
		l := []interface{}{left.(float64)}
		return isValid(l, right)
	} else {
		r := []interface{}{right.(float64)}
		return isValid(left, r)
	}
	return false, 0
}

func Part1(pairs []Pair) int {
	sumIndices := 0

	// if valid, ok := isValid(pairs[2].Left, pairs[2].Right); valid && ok == 1 {
	// 	sumIndices += 1
	// }
	for i, pair := range pairs {
		if valid, ok := isValid(pair.Left, pair.Right); valid && ok == 1 {
			sumIndices += i + 1
		}
	}

	return sumIndices
}
