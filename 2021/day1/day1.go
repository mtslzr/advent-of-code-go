package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Part1() {
	depths := readInput("2021/day1/input.txt")
	increases := 0
	for x := range depths {
		if x < 1 {
			continue
		}
		if depths[x] > depths[x-1] {
			increases += 1
		}
	}
	fmt.Println(increases)
}

func Part2() {
	depths := readInput("2021/day1/input.txt")
	increases := 0
	for x := range depths {
		if x < 3 || len(depths) - x < 2 {
			continue
		}
		curr := depths[x-2] + depths[x-1] + depths[x]
		prev := depths[x-3] + depths[x-2] + depths[x-1]
		if curr > prev {
			increases += 1
		}
	}
	fmt.Println(increases)
}

func readInput(file string) []int {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(buf), "\n")
	out := make([]int, len(data))
	for i := range data {
		out[i], _ = strconv.Atoi(data[i])
	}

	return out
}