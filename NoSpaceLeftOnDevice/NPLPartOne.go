package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []rune
}

type Operation string
type OperationValue int

func (stack *Stack) isEmpty() bool {
	return len(stack.items) <= 0
}

func (stack *Stack) size() int {
	if stack.isEmpty() {
		return 0
	}

	return len(stack.items)
}

func (stack *Stack) push(item rune) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() rune {
	if len(stack.items) <= 0 {
		return -1
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem
}

func (stack *Stack) peak() rune {
	if len(stack.items) <= 0 {
		return -1
	}

	return stack.items[len(stack.items)-1]
}

const (
	COMMAND         Operation      = "$"
	DIRECTORY       Operation      = "dir"
	COMMAND_VALUE   OperationValue = 1
	DIRECTORY_VALUE OperationValue = 2
)

const MAX_DIRACTORY_SIZE int = 100000

/*
Add the size of the last popped diractory to all the parent diractories
*/
func addSizeToTheParentDiractory(stack Stack, diractoryMap map[rune]int, size int) {
	if stack.isEmpty() {
		return
	}

	diractoryMap[stack.peak()] += size
}

/*
Determine what kind of operation does the line is requesting
*/
func getOperation(operation string) int {
	switch operation {
	case string(COMMAND):
		return 1
	case string(DIRECTORY):
		return 2
	default:
		return -1
	}
}

/* Get the largest value close to 100,000 */
func getLargetValue(diractoryMap map[rune]int) int {
	var largestSum int

	for _, value := range diractoryMap {
		if value <= MAX_DIRACTORY_SIZE && value > largestSum {
			largestSum = value
		}
	}

	return largestSum
}

func main() {
	file, err := os.Open("NSLInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	diractorySizeMap := make(map[rune]int)
	var directoryStack Stack

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")

		operation := -1
		if len(commands) <= 1 {
			continue
		}

		// Find out what kind of command the line is doing
		operation = getOperation(commands[0])

		var fileName rune
		if operation != -1 {
			// Based on the operation treat each line seperately
			switch operation {
			case int(COMMAND_VALUE):
				// Ignore 'ls' command
				if len(commands) == 2 {
					continue
				}

				file := commands[2]

				if file == ".." {
					lastDiractory := directoryStack.pop()
					// Add the size to the parent diractory
					addSizeToTheParentDiractory(directoryStack, diractorySizeMap, diractorySizeMap[lastDiractory])
				} else {
					fileName = []rune(file)[0]

					// Add diractory to the stack
					directoryStack.push(fileName)
					// Add the diractory to the size map
					diractorySizeMap[fileName] = 0
				}
			case int(DIRECTORY_VALUE):
				continue
			default:
				fmt.Println("Invalid operation found.")
			}
		} else {
			// Add the file size to the map
			value, err := strconv.Atoi(commands[0])
			if err != nil {
				fmt.Println("Invalid operation:", err)
			}

			// Store value in the parent file (store in the stack)
			fileName = directoryStack.peak()

			_, ok := diractorySizeMap[fileName]
			if !ok {
				fileName = []rune(commands[1])[0]
			}

			diractorySizeMap[fileName] += value
		}
	}

	// There should only by root diractory in the stack
	if directoryStack.size() >= 2 {
		for {
			if directoryStack.size() <= 1 {
				break
			}

			lastDiractory := directoryStack.pop()
			// Add the size to the parent diractory
			addSizeToTheParentDiractory(directoryStack, diractorySizeMap, diractorySizeMap[lastDiractory])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(diractorySizeMap)

	// Count total size of all the diractories
	fmt.Println("The sum of the total size of directories is:", getLargetValue(diractorySizeMap))
}
