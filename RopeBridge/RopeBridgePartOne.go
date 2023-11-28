package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MotionGrid [][]string

func main() {
	file, err := os.Open("RopeBridgeDemoInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	/* Initilze the motion grid */
	motionGrid := MotionGrid{
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{"S", ".", ".", ".", ".", "."},
	}

	/* Head & Tail */
	headRow := 4
	headCol := 0
	// tailRow := 5
	// tailCol := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		instuctionList := strings.Split(instruction, " ")
		if len(instruction) < 2 {
			log.Printf("Invalid instructions %v", instruction)
			continue
		}

		command := instuctionList[0]
		position, err := strconv.Atoi(instuctionList[1])
		if err != nil {
			log.Printf("Unable to parse instruction to int %v", instruction)
			continue
		}

		// Move Head
		headRow, headCol = motionGrid.moveHead(command, position, headRow, headCol)

		motionGrid.printMotionGrid()
		fmt.Println()
	}
}

func (motionGrid *MotionGrid) moveHead(command string, position, headRow, headCol int) (row, col int) {
	/* Set new Row & Col */
	newHeadRow := headRow
	newHeadCol := headCol

	switch command {
	case "L":
		newHeadCol = newHeadCol - position
	case "R":
		newHeadCol = newHeadCol + position
	case "U":
		newHeadRow = newHeadRow - position
	case "D":
		newHeadRow = newHeadRow + position
	}

	// Check if we are in the initial position
	if (*motionGrid)[headRow][headCol] != "S" {
		(*motionGrid)[headRow][headCol] = "."
	}

	// Update H's position
	(*motionGrid)[newHeadRow][newHeadCol] = "H"

	return newHeadRow, newHeadCol
}

func (motionGrid MotionGrid) printMotionGrid() {
	for _, row := range motionGrid {
		fmt.Println(row)
	}
}
