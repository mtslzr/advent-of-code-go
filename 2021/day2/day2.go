package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("2021/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	horiz := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		switch coords[0] {
		case "forward":
			horiz = horiz + num
		case "up":
			depth = depth - num
		case "down":
			depth = depth + num
		}
	}
	fmt.Println(horiz * depth)
}

func Part2() {
	file, err := os.Open("2021/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	aim := 0
	horiz := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}

		switch coords[0] {
		case "forward":
			horiz = horiz + num
			depth = depth + (aim * num)
		case "up":
			aim = aim - num
		case "down":
			aim = aim + num
		}
	}
	fmt.Println(horiz * depth)
}