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
	packets := ParsePacket(input)
	fmt.Printf("Part 2: %d\n", Part2(packets))

}

type Pair struct {
	Left  []interface{}
	Right []interface{}
}

type Packet struct {
	P   []interface{}
	Str string
}

func ParsePacket(input string) []Packet {
	packets := make([]Packet, 0)

	for _, packet := range strings.Split(input, "\n") {
		if packet == "" {
			continue
		}
		var p []interface{}
		json.Unmarshal([]byte(packet), &p)
		packets = append(packets, Packet{
			P:   p,
			Str: packet,
		})
	}

	var pDivider2 []interface{}
	json.Unmarshal([]byte("[[2]]"), &pDivider2)
	packets = append(packets, Packet{P: pDivider2, Str: "[[2]]"})
	var pDivider6 []interface{}
	json.Unmarshal([]byte("[[6]]"), &pDivider6)
	packets = append(packets, Packet{P: pDivider6, Str: "[[6]]"})

	return packets
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

func isInPath(path []Packet, val Packet) bool {
	for _, v := range path {
		if v.Str == val.Str {
			return true
		}
	}
	return false
}

func findPath(pairValids map[string]bool, packets []Packet, start Packet, path []Packet) []Packet {
	if len(path) == len(packets) {
		return path
	}

	path = append(path, start)

	for _, right := range packets {
		if start.Str == right.Str {
			continue
		}
		if valid, ok := pairValids[fmt.Sprintf("{%s,%s}", start.Str, right.Str)]; valid && ok && !isInPath(path, right) {
			newPath := findPath(pairValids, packets, right, path)
			if len(newPath) == len(packets) {
				return newPath
			}
		}
	}
	return path
}

func Part2(packets []Packet) int {
	pairCache := make(map[string]bool)
	for _, left := range packets {
		for _, right := range packets {
			if left.Str == right.Str {
				continue
			}
			valid, ok := isValid(left.P, right.P)
			pairCache[fmt.Sprintf("{%s,%s}", left.Str, right.Str)] = valid && ok == 1
		}
	}

	decoderKey := 1
	for _, start := range packets {
		path := findPath(pairCache, packets, start, nil)
		if len(path) == len(packets) {
			for i, p := range path {
				if p.Str == "[[2]]" || p.Str == "[[6]]" {
					decoderKey *= (i + 1)
				}
			}
		}
	}

	return decoderKey
}
