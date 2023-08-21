package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("partTwoDemo.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// PriorityItems counter
	groupBadgePriority := 0

	scanner := bufio.NewScanner(file)
	lines := make([]string, 3)
	for i := 0; scanner.Scan(); i = (i + 1) % 3 {
		lines[i] = scanner.Text()

		if i == 2 {
			groupBadgePriority += getGroupBadgePriority(lines[0], lines[1], lines[2])
		}
	}

	fmt.Printf("The sum of the priorities of the listed items types is: %d", groupBadgePriority)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getGroupBadgePriority(elfOneItems, elfTwoItems, elfThreeItems string) int {

	return 0
}
