package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func doesElfContainFullRange(firstElf, secondElf string) bool {
	firstElfSlice := strings.Split(firstElf, "-")
	secondElfSlice := strings.Split(secondElf, "-")
	// Get int values
	firstElfStart, err := strconv.Atoi(firstElfSlice[0])
	if err != nil {
		log.Fatal(err)
		return false
	}

	firstElfEnd, err := strconv.Atoi(firstElfSlice[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	secondElfStart, err := strconv.Atoi(secondElfSlice[0])
	if err != nil {
		log.Fatal(err)
		return false
	}

	secondElfEnd, err := strconv.Atoi(secondElfSlice[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	if firstElfStart < secondElfStart && firstElfEnd >= secondElfEnd {
		return true
	} else if firstElfStart > secondElfStart && firstElfEnd <= secondElfEnd {
		return true
	} else if firstElfStart == secondElfStart {
		if firstElfEnd >= secondElfEnd || firstElfEnd <= secondElfEnd {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("CCInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close the file after processing
	defer file.Close()

	fullRangeContainCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elvesSlice := strings.Split(line, ",")
		if doesElfContainFullRange(elvesSlice[0], elvesSlice[1]) {
			fullRangeContainCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d elves fully conatin their partner", fullRangeContainCounter)
}
