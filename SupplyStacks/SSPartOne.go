package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCargeFromLine(crateMap map[int][]rune, inputRow string) {
	for index, ascii := range inputRow {
		if ascii >= 65 && ascii <= 90 {
			crateMap[index/2] = append(crateMap[index/2], ascii)
		}
	}
}

func main() {
	file, err := os.Open("SSDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Map for the crates
	cargoMap := make(map[int][]rune)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		fmt.Println(line)

		if strings.Contains(line, "move") {
			// Move the cargo around based on the intructions

			continue
		}

		if !strings.Contains(line, "1") {
			// Get carge from each line
			getCargeFromLine(cargoMap, line)
		}
	}

	fmt.Println(cargoMap)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
