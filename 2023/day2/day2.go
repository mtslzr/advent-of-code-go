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
	file, err := os.Open("2023/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validGames := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		gameID, _ := strconv.Atoi(parts[0][5:])
		pulls := strings.Split(parts[1], "; ")
		validGame := true

		for _, pull := range pulls {
			colors := strings.Split(pull, ", ")
			for _, color := range colors {
				data := strings.Split(color, " ")
				numBalls, _ := strconv.Atoi(data[0])
				var maxBalls int

				switch data[1] {
				case "blue":
					maxBalls = 14
				case "red":
					maxBalls = 12
				case "green":
					maxBalls = 13
				}
				if numBalls > maxBalls {
					validGame = false
					break
				}
			}
			if !validGame {
				break
			}
		}
		if validGame {
			validGames += gameID
		}
	}
	fmt.Printf("Total of valid games: %d\n", validGames)
}

func Part2() {
	file, err := os.Open("2023/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		pulls := strings.Split(parts[1], "; ")
		minBalls := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		for _, pull := range pulls {
			colors := strings.Split(pull, ", ")
			for _, color := range colors {
				data := strings.Split(color, " ")
				numBalls, _ := strconv.Atoi(data[0])
				if numBalls > minBalls[data[1]] {
					minBalls[data[1]] = numBalls
				}
			}
		}
		totalSum += minBalls["blue"] * minBalls["green"] * minBalls["red"]
	}
	fmt.Printf("Total of powers of games: %d\n", totalSum)
}
