package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2022/day01"
	"github.com/umbe77/aoc-2022/day02"
	"github.com/umbe77/aoc-2022/day03"
	"github.com/umbe77/aoc-2022/day04"
	"github.com/umbe77/aoc-2022/day05"
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
	}
}
