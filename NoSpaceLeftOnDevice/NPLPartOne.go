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

func main() {
	file, err := os.Open("NSLDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	commandMap := make(map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")

		if len(commands) == 3 && commands[1] == "cd" {
			dir := rune(commands[2][0])
			_, ok := commandMap[dir]
			if !ok {
				commandMap[dir] = 0
			}
		}

		fmt.Println(commandMap)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result string

	fmt.Println("The sum of the total size of directories is:", result)
}
