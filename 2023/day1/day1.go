package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Part1() {
	file, err := os.Open("2023/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totals []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		coord := []rune(scanner.Text())
		coordNums := []string{}

		for _, x := range coord {
			if unicode.IsDigit(x) {
				coordNums = append(coordNums, string(x))
			}
		}
		outNum, _ := strconv.Atoi(coordNums[0] + coordNums[len(coordNums)-1])
		totals = append(totals, outNum)
	}

	total := 0
	for _, x := range totals {
		total += x
	}

	fmt.Printf("Part 1 Total of Sums: %d\n\n", total)
}

func Part2() {
	file, err := os.Open("2023/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordNums := map[string]string{
		"one": "1",
		"two": "2",
		"thr": "3",
		"fou": "4",
		"fiv": "5",
		"six": "6",
		"sev": "7",
		"eig": "8",
		"nin": "9",
	}

	var totals []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		coord := []rune(scanner.Text())
		coordNums := []string{}
		fmt.Printf("\n%s\n", string(coord))

		for y, x := range coord {
			if unicode.IsDigit(x) {
				coordNums = append(coordNums, string(x))
			} else if (y + 3) <= len(string(coord)) {
				val, ok := wordNums[string(coord[y:y+3])]
				if ok {
					coordNums = append(coordNums, val)
				}
			}
		}
		fmt.Println(coordNums)

		outNum, _ := strconv.Atoi(coordNums[0] + coordNums[len(coordNums)-1])
		fmt.Println(outNum)
		totals = append(totals, outNum)
	}

	total := 0
	for _, x := range totals {
		total += x
	}

	fmt.Printf("Part 2 Total of Sums: %d", total)
}
