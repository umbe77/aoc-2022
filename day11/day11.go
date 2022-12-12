package day11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day11/input.txt", func(line string) {
		input = append(input, line)
	})
	monkeys := Parse(input)
	fmt.Printf("Part 1: %d\n", Part1(monkeys))
	monkeys2 := Parse(input)
	fmt.Printf("Part 2: %d\n", Part2(monkeys2))
}

type Monkey struct {
	Name    int
	Items   []int
	Op      []string
	Test    int
	Test_ok int
	Test_ko int
}

func Add(a, b int) int {
	return a + b
}
func Mul(a, b int) int {
	return a * b
}

func Parse(data []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	var currentMonkey *Monkey
	for _, line := range data {
		if strings.HasPrefix(line, "Monkey") {
			monckeName := utils.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
			monkeys = append(monkeys, &Monkey{Name: monckeName})
			currentMonkey = monkeys[len(monkeys)-1]
		}
		line = strings.Trim(line, " ")
		token := strings.Split(line, ":")
		if strings.HasPrefix(token[0], "Starting") {
			itemsStr := strings.Split(strings.Trim(token[1], " "), ", ")
			items := make([]int, len(itemsStr))
			for i, item := range itemsStr {
				items[i] = utils.Atoi(item)
			}
			currentMonkey.Items = items
		}
		if strings.HasPrefix(token[0], "Operation") {
			vals := strings.ReplaceAll(strings.Trim(token[1], " "), "new = ", "")
			currentMonkey.Op = strings.Split(vals, " ")
		}
		if strings.HasPrefix(token[0], "Test") {
			tkTest := strings.Split(strings.Trim(token[1], " "), " ")
			currentMonkey.Test = utils.Atoi(tkTest[len(tkTest)-1])
		}
		if strings.HasPrefix(token[0], "If true") {
			currentMonkey.Test_ok = utils.Atoi(string(token[1][len(token[1])-1]))
		}
		if strings.HasPrefix(token[0], "If false") {
			currentMonkey.Test_ko = utils.Atoi(string(token[1][len(token[1])-1]))
		}
	}
	return monkeys
}

func Part1(monkeys []*Monkey) int {
	inspectedItems := make([]int, len(monkeys))
	for round := 1; round <= 20; round++ {
		for _, monkey := range monkeys {
			inspectedItems[monkey.Name] += len(monkey.Items)
			// fmt.Printf("%+v\n", monkey.Items)
			for _, item := range monkey.Items {
				new := 0
				p1 := item
				p2 := item
				if monkey.Op[0] != "old" {
					p1 = utils.Atoi(monkey.Op[0])
				}
				if monkey.Op[2] != "old" {
					p2 = utils.Atoi(monkey.Op[2])
				}
				switch monkey.Op[1] {
				case "+":
					new = Add(p1, p2)
					break
				case "*":
					new = Mul(p1, p2)
					break
				}
				new = new / 3
				if new%monkey.Test == 0 {
					// fmt.Printf("new: %d -- OK to %d\n", new, monkey.Test_ok)
					monkeys[monkey.Test_ok].Items = append(monkeys[monkey.Test_ok].Items, new)
					// fmt.Printf("%+v\n", monkeys[monkey.Test_ok].Items)
				} else {
					// fmt.Printf("new: %d -- KO to %d\n", new, monkey.Test_ko)
					monkeys[monkey.Test_ko].Items = append(monkeys[monkey.Test_ko].Items, new)
					// fmt.Printf("%+v\n", monkeys[monkey.Test_ko].Items)
				}
				// fmt.Println("---------")
			}
			monkey.Items = make([]int, 0)
		}
	}
	// for _, v := range inspectedItems {
	// 	fmt.Println(v)
	// }
	sort.Ints(inspectedItems)
	l := len(inspectedItems)
	return inspectedItems[l-1] * inspectedItems[l-2]
}
func Part2(monkeys []*Monkey) int {
	inspectedItems := make([]int, len(monkeys))
	for round := 1; round <= 10000; round++ {
		for mi := 0; mi < len(monkeys); mi++ {
			monkey := monkeys[mi]
			inspectedItems[monkey.Name] += len(monkey.Items)
			for _, item := range monkey.Items {
				new := 0
				p1 := item
				p2 := item
				if monkey.Op[0] != "old" {
					p1 = utils.Atoi(monkey.Op[0])
				}
				if monkey.Op[2] != "old" {
					p2 = utils.Atoi(monkey.Op[2])
				}
				switch monkey.Op[1] {
				case "+":
					new = Add(p1, p2)
					break
				case "*":
					new = Mul(p1, p2)
					break
				}
				destMonkey := 0
				if new%monkey.Test == 0 {
					destMonkey = monkey.Test_ok
				} else {
					destMonkey = monkey.Test_ko
				}
				monkeys[destMonkey].Items = append(monkeys[destMonkey].Items, new)
			}
			monkey.Items = nil
		}
	}
	for _, v := range inspectedItems {
		fmt.Println(v)
	}
	sort.Ints(inspectedItems)
	l := len(inspectedItems)
	return inspectedItems[l-1] * inspectedItems[l-2]
}
