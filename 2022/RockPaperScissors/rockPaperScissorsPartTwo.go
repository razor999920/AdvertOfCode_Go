package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getRoundScore(opp, mine string) int {
	const (
		Rock     = 1
		Paper    = 2
		Scissors = 3

		LOSS = 0
		DRAW = 3
		WIN  = 6
	)

	switch opp {
	case "A":
		switch mine {
		case "X":
			return Scissors + LOSS
		case "Y":
			return Rock + DRAW
		case "Z":
			return Paper + WIN
		}
	case "B":
		switch mine {
		case "X":
			return Rock + LOSS
		case "Y":
			return Paper + DRAW
		case "Z":
			return Scissors + WIN
		}
	case "C":
		switch mine {
		case "X":
			return Paper + LOSS
		case "Y":
			return Scissors + DRAW
		case "Z":
			return Rock + WIN
		}
	}

	return LOSS
}

func main() {
	file, err := os.Open("strategyGuide.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close() // return the  file when the function ends

	totalScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Fields(line)

		// Each line has 2 columns
		// col 1: oppnent's move
		// col 2: strategy
		totalScore += getRoundScore(moves[0], moves[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	fmt.Printf("The final score is: %d", totalScore)
}
