package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2022/day01"
	"github.com/umbe77/aoc-2022/day02"
	"github.com/umbe77/aoc-2022/day03"
	"github.com/umbe77/aoc-2022/day04"
	"github.com/umbe77/aoc-2022/day05"
	"github.com/umbe77/aoc-2022/day06"
	"github.com/umbe77/aoc-2022/day07"
	"github.com/umbe77/aoc-2022/day08"
	"github.com/umbe77/aoc-2022/day09"
	"github.com/umbe77/aoc-2022/day10"
	"github.com/umbe77/aoc-2022/day11"
	"github.com/umbe77/aoc-2022/day12"
	"github.com/umbe77/aoc-2022/day13"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "01", "day in format dd")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	switch day {
	case "01":
		day01.Execute()
	case "02":
		day02.Execute()
	case "03":
		day03.Execute()
	case "04":
		day04.Execute()
	case "05":
		day05.Execute()
	case "06":
		day06.Execute()
	case "07":
		day07.Execute()
	case "08":
		day08.Execute()
	case "09":
		day09.Execute()
	case "10":
		day10.Execute()
	case "11":
		day11.Execute()
	case "12":
		day12.Execute()
	case "13":
		day13.Execute()
	}
}
