package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Stack struct {
	items []rune
}

type Operation string
type OperationValue int

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
	FILE            Operation      = "file"
	COMMAND_VALUE   OperationValue = 1
	DIRECTORY_VALUE OperationValue = 2
	FILE_VALUE      OperationValue = 3
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
	case string(FILE):
		return 3
	default:
		return -1
	}
}

func main() {
	file, err := os.Open("NSLDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	commandMap := make(map[rune]int)
	var directoryStack Stack

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")

		fmt.Println(commands)

		operation := -1
		if len(commands) >= 1 {
			operation = getOperation(commands[0])
		}

		if operation != -1 {
			// Based on the operation treat each line seperately
			switch operation {
			case int(COMMAND_VALUE):
				if len(commands) == 2 {
					continue
				}

				file := commands[2]

				if file == ".." {
					directoryStack.pop()
				} else {
					directoryStack.push([]rune(file)[0]) // Push the file name to the stack. For example, stac.Push('a')
				}
			case int(DIRECTORY_VALUE):
			case int(FILE_VALUE):
			default:
				fmt.Println("Invalid operation found.")
			}
		}

		if len(commands) == 3 && commands[1] == "cd" {
			dir := rune(commands[2][0])
			_, ok := commandMap[dir]
			if !ok {
				commandMap[dir] = 0
			}
		}

		// Print DS
		// fmt.Println(commandMap)
		// fmt.Println(directoryStack)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result string

	fmt.Println("The sum of the total size of directories is:", result)
}
