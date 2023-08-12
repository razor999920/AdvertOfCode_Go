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

	switch mine {
	case "X":
		switch opp {
		case "A":
			return Rock + DRAW
		case "B":
			return Rock + LOSS
		case "C":
			return Rock + WIN
		}
	case "Y":
		switch opp {
		case "A":
			return Paper + WIN
		case "B":
			return Paper + DRAW
		case "C":
			return Paper + LOSS
		}
	case "Z":
		switch opp {
		case "A":
			return Scissors + LOSS
		case "B":
			return Scissors + WIN
		case "C":
			return Scissors + DRAW
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
		// col 2: my move
		totalScore += getRoundScore(moves[0], moves[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
	fmt.Printf("The final score is: %d", totalScore)
}
