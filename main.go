package main

import (
	"fmt"
	"os"

	day118 "github.com/mtslzr/advent-of-code-go/2018/day1"
	day119 "github.com/mtslzr/advent-of-code-go/2019/day1"
	day219 "github.com/mtslzr/advent-of-code-go/2019/day2"
	"github.com/mtslzr/advent-of-code-go/2021/day1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <year> <day>")
	} else {
		switch os.Args[1] {
		case "2021":
			switch os.Args[2] {
			case "day1":
				day1.Part1()
				day1.Part2()
			}
		case "2019":
			switch os.Args[2] {
			case "day1":
				day119.Part1()
				day119.Part2()
			case "day2":
				day219.Part1()
				day219.Part2()
			}
		case "2018":
			switch os.Args[2] {
			case "day1":
				day118.Part1()
				day118.Part2()
			}
		}
	}
}
