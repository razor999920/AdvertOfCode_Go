package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// CrateStack represents a stack of crates
type CrateStack map[int][]rune

// Push adds a new crate to the specified stack
func (cs CrateStack) Push(stackIndex int, crate rune) {
	cs[stackIndex] = append([]rune{crate}, cs[stackIndex]...)
}

// Pop removes and returns the top crate from the specified stack
func (cs CrateStack) Pop(stackIndex int) (rune, bool) {
	stack := cs[stackIndex]
	if len(stack) == 0 {
		return 0, false
	}
	crate := stack[len(stack)-1]
	cs[stackIndex] = stack[:len(stack)-1]
	return crate, true
}

// ParseAndExecute parses a command and performs the necessary actions
func (cs CrateStack) ParseAndExecute(command string) {
	if command == "" {
		return
	}
	if strings.Contains(command, "move") {
		var num, from, to int
		fmt.Sscanf(command, "move %d from %d to %d", &num, &from, &to)
		cs.Move(num, from, to)
	} else if !strings.Contains(command, "1") {
		cs.AddCrates(command)
	}
}

// Move moves a number of crates from one stack to another
func (cs CrateStack) Move(num, from, to int) {
	crateStack := cs[from][(len(cs[from]) - num):]
	cs[from] = cs[from][:len(cs[from])-num]
	cs[to] = append(cs[to], crateStack...)
}

// AddCrates adds crates to the stacks based on the input string
func (cs CrateStack) AddCrates(input string) {
	for index, ascii := range input {
		if ascii >= 'A' && ascii <= 'Z' {
			key := index/4 + 1
			cs.Push(key, ascii)
		}
	}
}

func main() {
	file, err := os.Open("SSInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	crateStacks := CrateStack{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		crateStacks.ParseAndExecute(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result string
	for i := 1; i <= len(crateStacks); i++ {
		if crate, ok := crateStacks.Pop(i); ok {
			result += string(crate)
		}
	}

	fmt.Println("The crates that will end up on top are:", result)
}
