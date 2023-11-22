package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("TreeHouseInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Assuming all rows have the same length, determine cols from the first line
	var cols int
	if scanner.Scan() {
		cols = len(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Reset the file pointer to the start of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Re-initialize the scanner and create the 2D array
	scanner = bufio.NewScanner(file)
	forestGrid := make([][]int, 0, cols) // Set capacity to cols, but length to 0
	for scanner.Scan() {
		line := scanner.Text()
		forestRow, err := generateListFromInput(line)
		if err != nil {
			log.Fatal(err)
		}
		forestGrid = append(forestGrid, forestRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	visibleTree := 0

	// Check visibility for each tree
	for rowIndex := range forestGrid {
		for colIndex := range forestGrid[rowIndex] {
			// Edge trees are always visible
			if rowIndex == 0 || rowIndex == len(forestGrid)-1 || colIndex == 0 || colIndex == len(forestGrid[rowIndex])-1 {
				visibleTree++
				continue
			}

			if isTreeVisible(forestGrid, rowIndex, colIndex, -1, 0) || isTreeVisible(forestGrid, rowIndex, colIndex, 1, 0) || isTreeVisible(forestGrid, rowIndex, colIndex, 0, -1) || isTreeVisible(forestGrid, rowIndex, colIndex, 0, 1) {
				visibleTree++
			}
		}
	}

	fmt.Printf("There are %d trees visible from outside the grid", visibleTree)
}

// Updated generateListFromInput function to return a slice and an error
func generateListFromInput(input string) ([]int, error) {
	forestRow := make([]int, len(input))
	for i, char := range input {
		if !unicode.IsDigit(char) {
			return nil, fmt.Errorf("invalid character in input: %v", string(char))
		}
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		forestRow[i] = num
	}
	return forestRow, nil
}

func isTreeVisible(forestGrid [][]int, row int, col int, rowDir int, colDir int) bool {
	currentTree := forestGrid[row][col]

	for r, c := row+rowDir, col+colDir; r >= 0 && r < len(forestGrid) && c >= 0 && c < len(forestGrid[row]); r, c = r+rowDir, c+colDir {
		if forestGrid[r][c] >= currentTree {
			return false
		}
	}

	return true
}

// func getVisibleTree(forestGrid [][]int) int {
// 	visibleTree := 0

// 	for rowIndex, row := range forestGrid {
// 		for colIndex := range row {
// 			// Edge trees are always visible
// 			if rowIndex == 0 || rowIndex == len(forestGrid)-1 || colIndex == 0 || colIndex == len(row)-1 {
// 				visibleTree++
// 				continue
// 			}

// 			currentTree := forestGrid[rowIndex][colIndex]
// 			isVisible := false

// 			// Check in all directions, breaking early if visibility is confirmed
// 			// LEFT
// 			for c := colIndex - 1; c >= 0; c-- {
// 				if forestGrid[rowIndex][c] < currentTree {
// 					// fmt.Println("LEFT", rowIndex, colIndex, forestGrid[rowIndex][colIndex])
// 					isVisible = true
// 				} else {
// 					isVisible = false
// 					break
// 				}
// 			}

// 			// RIGHT
// 			if !isVisible {
// 				for c := colIndex + 1; c < len(row); c++ {
// 					if forestGrid[rowIndex][c] < currentTree {
// 						// fmt.Println("RIGHT", rowIndex, colIndex, forestGrid[rowIndex][colIndex])
// 						isVisible = true
// 					} else {
// 						isVisible = false
// 						break
// 					}
// 				}
// 			}

// 			// TOP
// 			if !isVisible {
// 				for r := rowIndex - 1; r >= 0; r-- {
// 					if forestGrid[r][colIndex] < currentTree {
// 						// fmt.Println("TOP", rowIndex, colIndex, forestGrid[rowIndex][colIndex])
// 						isVisible = true
// 					} else {
// 						isVisible = false
// 						break
// 					}
// 				}
// 			}

// 			// BOTTOM
// 			if !isVisible {
// 				for r := rowIndex + 1; r < len(forestGrid); r++ {
// 					if forestGrid[r][colIndex] < currentTree {
// 						// fmt.Println("BOTTOM", rowIndex, colIndex, forestGrid[rowIndex][colIndex])
// 						isVisible = true
// 					} else {
// 						isVisible = false
// 						break
// 					}
// 				}
// 			}

// 			if isVisible {
// 				visibleTree++
// 			}
// 		}
// 	}

// 	return visibleTree
// }
