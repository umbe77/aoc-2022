package day12

import (
	"fmt"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day12/input.txt", func(line string) {
		input = append(input, line)
	})
	fmt.Printf("Part 1: %d\n", Part1(Parse(input)))
	grid, width, height, _, end := Parse(input)
	fmt.Printf("Part 2: %d\n", Part2(grid, width, height, end))
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func Parse(input []string) (map[Point]int, int, int, Point, Point) {
	grid := make(map[Point]int)
	start := Point{}
	end := Point{}
	height := len(input)
	width := 0
	for y, line := range input {
		for x, c := range line {
			if width == 0 {
				width = len(line)
			}
			v := 0
			if c == 'S' {
				start.X = x
				start.Y = y
				v = int('a') - 97
			} else if c == 'E' {
				end.X = x
				end.Y = y
				v = int('z') - 97
			} else {
				v = int(c) - 97
			}
			grid[Point{X: x, Y: y}] = v
		}
	}
	return grid, width, height, start, end
}

func isInPath(path []Point, val Point) bool {
	for _, v := range path {
		if v == val {
			return true
		}
	}
	return false
}

var directions = []Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}

func get(dst map[Point]int, p Point) int {
	v, ok := dst[p]
	if !ok {
		return 10000
	}
	return v
}
func Part1(grid map[Point]int, width, height int, start, end Point) int {
	return FindPath(grid, width, height, start, end)
}

func Part2(grid map[Point]int, width, height int, end Point) int {
	startingPoints := make([]Point, 0)
	for k, p := range grid {
		if p == 0 {
			startingPoints = append(startingPoints, k)
		}
	}

	shortestPath := 10000
	for _, sp := range startingPoints {
		path := FindPath(grid, width, height, sp, end)
		if path != -1 {
			shortestPath = utils.Min(shortestPath, path)
		}
	}
	return shortestPath
}

func FindPath(grid map[Point]int, width, height int, start, end Point) int {

	queue := utils.New[Point]()

	dst := make(map[Point]int)
	dst[start] = 0

	queue.Push(start)
	for !queue.IsEmpty() {
		c := queue.Dequeue()
		if c == end {
			return get(dst, end)
		}
		for _, dir := range directions {
			n := c.Add(dir)
			if n.X >= 0 && n.X < width && n.Y >= 0 && n.Y < height {
				if grid[c] >= grid[n]-1 {
					ndst := get(dst, c) + 1
					if ndst < get(dst, n) {
						queue.Push(n)
						dst[n] = ndst
					}
				}
			}
		}

	}
	return -1
}
