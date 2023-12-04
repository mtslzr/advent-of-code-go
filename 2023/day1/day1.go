package day1

import (
	"bufio"
	"bytes"
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
		newNum := bytes.Buffer{}
		coord := []rune(scanner.Text())
		var revCoord string
		for _, x := range coord {
			if unicode.IsDigit(x) && len(newNum.String()) < 1 {
				newNum.WriteString(string(x))
			}
			revCoord = string(x) + revCoord
		}
		for _, x := range revCoord {
			if unicode.IsDigit(x) && len(newNum.String()) < 2 {
				newNum.WriteString(string(x))
				break
			}
		}
		outNum, _ := strconv.Atoi(newNum.String())
		totals = append(totals, outNum)
	}

	total := 0
	for _, x := range totals {
		total += x
	}

	fmt.Printf("Total of sums: %d", total)
}
