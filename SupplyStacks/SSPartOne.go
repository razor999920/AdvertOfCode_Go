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
		var key int

		if ascii >= 65 && ascii <= 90 {
			if index == 1 {
				key = index
			} else {
				key = ((index - 1) / 4) + 1
			}

			crateMap[key] = append([]rune{ascii}, crateMap[key]...)
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

		if strings.Contains(line, "move") {
			// Move the cargo around based on the intructions
			instructions := strings.Fields(line)

			var num int
			var from int
			var to int
			for index, instruction := range instructions {
				var value int

				if _, err := fmt.Sscanf(instruction, "%d", &value); err != nil {
					continue
				}

				if index == 1 {
					num = value

					continue
				}
				if index == 3 {
					from = value

					continue
				}

				if index == 5 {
					to = value

					i := 0
					for num > i {
						top := cargoMap[from][len(cargoMap[from])-1]
						cargoMap[to] = append(cargoMap[to], top)
						// Readjuct the from list
						cargoMap[from] = cargoMap[from][:len(cargoMap[from])-1]

						i++
					}
				}
			}

			continue
		}

		if !strings.Contains(line, "1") {
			// Get carge from each line
			getCargeFromLine(cargoMap, line)
		}
	}

	// Get head
	var crateStack string
	for index := 1; index <= len(cargoMap); index++ {
		top := cargoMap[index][len(cargoMap[index])-1]
		// Add the cargo to the string
		crateStack += string(top)
	}

	// Result
	fmt.Println("The crates that will end up on top are: ", crateStack)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
