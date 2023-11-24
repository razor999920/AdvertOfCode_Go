package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Print the grid
	motionGrid.printMotionGrid()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}

func (motionGrid MotionGrid) printMotionGrid() {
	for _, row := range motionGrid {
		fmt.Println(row)
	}
}
