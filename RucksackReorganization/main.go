package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// PriorityItems counter
	priorityItemsTotal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		priorityItemsTotal += getInvalidItem(compartment1, compartment2)
	}

	fmt.Printf("The sum of the priorities of the listed items types is: %d", priorityItemsTotal)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getInvalidItem(compartment1, compartment2 string) int {
	compartmentSet := make(map[rune]struct{})

	for _, key := range compartment1 {
		compartmentSet[key] = struct{}{}
	}

	for _, key := range compartment2 {
		_, exists := compartmentSet[key]
		if exists {
			fmt.Printf("%c", key)
			acsiPosition := int(key)

			if acsiPosition < 97 {
				return acsiPosition - int('A') + 27
			}

			return acsiPosition - int('a') + 1
		}

	}

	return 0
}
