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

	for key, value := range diractoryMap {
		fmt.Println(key, value)
	}

	return largestSum
}

func main() {
	file, err := os.Open("NSLDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	diractoryMap := make(map[rune]int)
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
					directoryStack.pop()
				} else {
					fileName = []rune(file)[0]

					// Push the file name to the stack. For example, stack.Push('a')
					// Only push the parent file to the stack
					if fileName == '/' || !directoryStack.isEmpty() {
						continue
					}

					directoryStack.push(fileName)
					// Add the diractory to the map
					diractoryMap[fileName] = 0
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

			// fmt.Println(directoryStack)

			// Store value in the parent file (store in the stack)
			fileName = directoryStack.peak()

			_, ok := diractoryMap[fileName]
			if !ok {
				fileName = []rune(commands[1])[0]
			}

			diractoryMap[fileName] += value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print DS
	fmt.Println(diractoryMap)

	// Count total size of all the diractories
	getLargetValue(diractoryMap)

	fmt.Println("The sum of the total size of directories is:", 0)
}
