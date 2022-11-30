package main

import (
	"flag"
	"fmt"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "01", "day in format dd")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	switch day {
	case "01":
		fmt.Println("Hello World!")
	}
}
