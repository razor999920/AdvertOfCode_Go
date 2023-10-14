package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
