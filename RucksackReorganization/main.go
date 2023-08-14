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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		compartment1 := line[:len(line)/2]
		comompartment2 := line[len(line)/2:]
		println(compartment1, comompartment2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
