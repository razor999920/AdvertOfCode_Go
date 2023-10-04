package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func addToMap(crateMap map[int]string, inputRow string) {

}

func main() {
	file, err := os.Open("SSDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Map for the crates
	createMap = make(map[int]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Println("Moves")
		}

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
