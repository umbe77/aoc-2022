package day15

import (
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2022/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day15/input.txt", func(line string) {
		// utils.ReadFile("day15/sample_input.txt", func(line string) {
		input = append(input, line)
	})
	positions, lower, upper := Parse(input)
	// fmt.Printf("%+v\nL: %+v, U: %+v\n", positions, lower, upper)
	// fmt.Println(upper.x - lower.x)
	fmt.Printf("Part 1: %d\n", Part1(positions, 2000000, lower, upper))
	// fmt.Printf("Part 1: %d\n", Part1(positions, 10, lower, upper))

	fmt.Printf("Part 2: %d\n", Part2(positions, Point{x: 0, y: 0}, Point{x: 4000000, y: 4000000}))
	// fmt.Printf("Part 2: %d\n", Part2(positions, Point{x: 0, y: 0}, Point{x: 20, y: 20}))

}

type Point struct {
	x, y int
}

type Position struct {
	sensor, beacon Point
	distance       int
}

func getDistance(a, b Point) int {
	return utils.Abs(a.x-b.x) + utils.Abs(a.y-b.y)
}
func Parse(input []string) ([]Position, Point, Point) {
	positions := make([]Position, len(input))
	lowerX, lowerY := 100000, 100000
	upperX, upperY := 0, 0

	toPoint := func(p string) Point {
		coord := strings.Split(p, ", ")
		return Point{
			x: utils.Atoi(strings.Split(coord[0], "=")[1]),
			y: utils.Atoi(strings.Split(coord[1], "=")[1]),
		}
	}

	for i, line := range input {
		pair := strings.Split(line, ": ")
		sensor_point := strings.ReplaceAll(pair[0], "Sensor at ", "")
		beacon_point := strings.ReplaceAll(pair[1], "closest beacon is at ", "")

		sensor := toPoint(sensor_point)
		beacon := toPoint(beacon_point)

		lowerX = utils.Min(lowerX, utils.Min(sensor.x, beacon.x))
		lowerY = utils.Min(lowerY, utils.Min(sensor.y, beacon.y))
		upperX = utils.Max(upperX, utils.Max(sensor.x, beacon.x))
		upperY = utils.Max(upperY, utils.Max(sensor.y, beacon.y))

		distance := getDistance(sensor, beacon)

		positions[i] = Position{
			sensor:   sensor,
			beacon:   beacon,
			distance: distance,
		}

	}

	return positions, Point{x: lowerX, y: lowerY}, Point{x: upperX, y: upperY}
}

func isBeacon(positions []Position, p Point) bool {
	for _, position := range positions {
		if p == position.beacon {
			return true
		}
	}
	return false
}

func Part1(positions []Position, line int, lower, upper Point) int {

	count := 0

	for i := lower.x - 1000000; i <= upper.x+1000000; i++ {
		p := Point{x: i, y: line}
		no_beacon := false
		for _, position := range positions {
			d := getDistance(position.sensor, p)
			if d <= position.distance && !isBeacon(positions, p) {
				no_beacon = true
				break
			}
		}
		if no_beacon {
			count++
		}
	}

	return count
}

func Part2(positions []Position, lower, upper Point) int {
	for y := lower.y; y <= upper.y; y++ {
		for x := lower.x; x <= upper.x; x++ {
			p := Point{x: x, y: y}
			outOfRange := true
			for _, position := range positions {
				if isBeacon(positions, p) {
					outOfRange = false
					break
				}
				d := getDistance(position.sensor, p)
				if d <= position.distance {
					outOfRange = false
					break
				}
			}
			if outOfRange {
				return p.x*4000000 + p.y
			}
		}
	}
	return 0
}
