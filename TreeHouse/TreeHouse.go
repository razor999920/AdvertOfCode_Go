package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func generateListFromInput(forestGrid [][]int, gridRow int, input string) {

	gridCounter := 0
	/*
		Iterate over the input
		convert each char to int and add it to the 2D array
	*/
	for _, char := range input {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// Convert the character to an integer
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error converting character to integer:", err)
				continue
			}
			// Append the integer to the 2D slice
			forestGrid[gridRow][gridCounter] = num

			gridCounter++
		}
	}
}

func getVisibleTree(forestGrid [][]int) int {
	visibleTree := 0

	return visibleTree
}

func main() {
	file, err := os.Open("TreeHouseDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close this file when system stops executing
	defer file.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(file)

	rows := 0

	for scanner.Scan() {
		rows += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Reset the file pointer to the start of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Now, scan the file again to create a 2D slice
	scanner = bufio.NewScanner(file)
	// Create the 2D array
	cols := rows
	forestGrid := make([][]int, rows)
	for i := range forestGrid {
		forestGrid[i] = make([]int, cols)
	}
	gridRowCounter := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Generate the 2D array from every line
		generateListFromInput(forestGrid, gridRowCounter, line)

		// Increment the row counter
		gridRowCounter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	/*
	 Now we have the populated 2D array
	 Figure out how many trees are visible
	*/
	visibleTree := getVisibleTree(forestGrid)

	fmt.Printf("There are %d trees visible from outside the grid", visibleTree)
}
