package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func moveTotal(c string) int {
	const (
		Rock     = 1
		Paper    = 2
		Scissors = 3
	)

	switch c {
	case "X", "A":
		return Rock
	case "Y", "B":
		return Paper
	case "Z", "C":
		return Scissors
	default:
		return 0
	}
}

func getRoundScore(opp, mine string) int {
	player1 := moveTotal(opp)
	player2 := moveTotal(mine)

	if player1 == player2 {
		return 3
	} else if player2 > player1 {
		return 6
	}

	return 0
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
		sum := 0

		// Each line has 2 columns
		// col 1: oppnent's move
		// col 2: my move
		sum = moveTotal(moves[1])
		// Get the score from the round
		sum += getRoundScore(moves[0], moves[1])

		// Sum the score
		totalScore += sum
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	fmt.Printf("The final score is: %d", totalScore)
}
