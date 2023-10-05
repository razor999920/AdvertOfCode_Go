package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCargeFromLine(crateMap map[int]rune, inputRow string) {
	fmt.Println(inputRow)

	for index, ascii := range inputRow {
		if ascii >= 65 && ascii <= 90 {
			char := rune(ascii)
			crateMap[index/2] = char
			fmt.Println(ascii)
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
	cargoMap := make(map[int]rune)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "1") {
			// Get carge from each line
			getCargeFromLine(cargoMap, line)
		}

		if line == "" {
			continue
		}
	}

	fmt.Println(cargoMap)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
