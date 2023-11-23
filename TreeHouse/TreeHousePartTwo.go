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

	var rowLength int
	for scanner.Scan() {
		rowLength = len(scanner.Text())
		break
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Reset to point of the file to the start
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// 2D GRID
	forestGrid := make([][]int, 0, rowLength)
	// Scenic Score GRID
	scenicGrid := make([]int, rowLength)

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Populate the grid from the scanner
		forestGridRow, err := generateGridRow(line)
		if err != nil {
			log.Fatal(err)
			continue
		}

		forestGrid = append(forestGrid, forestGridRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	/*
		Uinsg the forest grid, we need to calculate the scenic score at each index
	*/
	for row, rowVal := range forestGrid {
		for col := range rowVal {
			if row-1 < 0 || col-1 < 0 || row+1 >= rowLength || col+1 >= rowLength {
				continue
			}

			// Find the scenic score of the current tree
			scenicScore := 0

			left := getScenicScoreForTree(forestGrid, row, col, 0, -1)
			right := getScenicScoreForTree(forestGrid, row, col, 0, +1)
			up := getScenicScoreForTree(forestGrid, row, col, -1, 0)
			down := getScenicScoreForTree(forestGrid, row, col, +1, 0)

			scenicScore = up * down * left * right

			// Add Scenic score for the current tree
			scenicGrid = append(scenicGrid, scenicScore)
		}
	}

	fmt.Printf("The highest scenic score for a tree in the following grid is %v", getHighestScenicScore(scenicGrid))
}

func generateGridRow(line string) ([]int, error) {
	forestRow := make([]int, len(line))
	for i, char := range line {
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

func getScenicScoreForTree(forestGrid [][]int, row, col, rowDir, colDir int) int {
	currentTree := forestGrid[row][col]
	scenicScore := 0

	for r, c := row+rowDir, col+colDir; r >= 0 && c >= 0 && r < len(forestGrid) && c < len(forestGrid); r, c = r+rowDir, c+colDir {
		scenicScore++

		if currentTree <= forestGrid[r][c] {
			break
		}
	}

	if scenicScore == 0 {
		scenicScore = 1
	}

	return scenicScore
}

func getHighestScenicScore(scenicGrid []int) int {
	if len(scenicGrid) == 0 {
		return 0
	}

	scenicScore := 0

	for _, score := range scenicGrid {
		if score >= scenicScore {
			scenicScore = score
		}
	}

	return scenicScore
}
