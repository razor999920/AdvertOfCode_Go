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
type Coordinates struct {
	x, y interface{}
}

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
	tailRow := 4
	tailCol := 0
	// Tail Position Counter
	tailCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		instructionList := strings.Split(instruction, " ")
		if len(instruction) < 2 {
			log.Printf("Invalid instructions %v", instruction)
			continue
		}

		command := instructionList[0]
		position, err := strconv.Atoi(instructionList[1])
		if err != nil {
			log.Printf("Unable to parse instruction to int %v", instruction)
			continue
		}

		// Move Head
		currentTailCounter := 0
		headRow, headCol, tailRow, tailCol, currentTailCounter = motionGrid.moveHead(command, position, headRow, headCol, tailRow, tailCol)
		// Add Tail's position
		tailCounter += currentTailCounter

		motionGrid.printMotionGrid()
		fmt.Println()
	}

	fmt.Printf("The Tail of the rope visits %d postions", tailCounter)
}

func (motionGrid *MotionGrid) moveHead(command string, position, headRow, headCol, tailRow, tailCol int) (hRow, hCol, tRow, tCol, tailCounter int) {
	/* Set new Row & Col */
	newHeadRow := headRow
	newHeadCol := headCol
	newTailRow := tailRow
	newTailCol := tailCol

	tailCounter = 0
	positionCounter := 1

	for positionCounter <= position {
		switch command {
		case "L":
			newHeadCol = newHeadCol - 1
		case "R":
			newHeadCol = newHeadCol + 1
		case "U":
			newHeadRow = newHeadRow - 1
		case "D":
			newHeadRow = newHeadRow + 1
		}

		// HEAD's new position on the grid
		// Check if we are in the initial position
		if (*motionGrid)[headRow][headCol] != "S" {
			(*motionGrid)[headRow][headCol] = "."
		}

		// Update H's position
		(*motionGrid)[newHeadRow][newHeadCol] = "H"
		headRow = newHeadRow
		headCol = newHeadCol

		positionCounter += 1

		// Figure out Tail's new position based on H's position
		// If Tail is touching Head to ignore
		if ((newHeadRow-1 == newTailRow || newHeadRow+1 == newTailRow) && (newHeadCol == newTailCol)) || ((newHeadCol-1 == newTailCol || newHeadCol+1 == newTailCol) && (newHeadRow == newTailRow)) || (newHeadRow == newTailRow && newHeadCol == newTailCol) {
			continue
		}

		if newHeadRow == newTailRow && newHeadCol != newTailCol {
			if newHeadCol > newTailCol {
				newTailCol += 1
			} else {
				newTailCol -= 1
			}
		} else if newHeadRow != newTailRow && newHeadCol == newTailCol {
			if newHeadRow > newTailRow {
				newTailRow += 1
			} else {
				newTailRow -= 1
			}
		} else {
			if newHeadRow == newTailRow+2 {
				newTailRow += 1
				if newHeadCol > newTailCol {
					newTailCol += 1
				} else if newHeadCol < newTailCol {
					newTailCol -= 1
				}
			} else if newHeadRow == newTailRow-2 {
				newTailRow -= 1

				if newHeadCol > newTailCol {
					newTailCol += 1
				} else if newHeadCol < newTailCol {
					newTailCol -= 1
				}
			} else if newHeadCol == newTailCol+2 {
				newTailCol += 1

				if newHeadRow > newTailRow {
					newTailRow += 1
				} else if newHeadRow < newTailRow {
					newTailRow -= 1
				}
			} else if newHeadCol == newTailCol-2 {
				newTailCol -= 1

				if newHeadRow > newTailRow {
					newTailRow += 1
				} else if newHeadRow < newTailRow {
					newTailRow -= 1
				}
			}
		}

		// Make sure the tail position changed
		if tailRow == newTailRow && tailCol == newTailCol {
			continue
		}

		// TAIL's new position on the grid
		// Remove tail's trail from the grid (Ignore for first iteration)
		if (*motionGrid)[tailRow][tailCol] != "S" {
			(*motionGrid)[tailRow][tailCol] = "."
		}
		// Update T's position
		(*motionGrid)[newTailRow][newTailCol] = "T"
		tailRow = newTailRow
		tailCol = newTailCol

		tailCounter += 1
	}

	return newHeadRow, newHeadCol, newTailRow, newTailCol, tailCounter
}

func coordinatesExists(visitedCoordinates []Coordinates, coordinates Coordinates) bool {
	for _, c := range visitedCoordinates {
		if c == coordinates {
			return true
		}
	}

	return false
}

func (motionGrid MotionGrid) printMotionGrid() {
	for _, row := range motionGrid {
		fmt.Println(row)
	}
}
