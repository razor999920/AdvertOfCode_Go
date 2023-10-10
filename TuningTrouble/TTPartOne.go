package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Peak() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return (*q)[0]
}

func (q *Queue) PeakLast() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return (*q)[len(*q)-1]
}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element) // add to the end of queue
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	element := (*q)[0]
	*q = (*q)[1:]

	return element
}

func getFirstMarker(signal string) int {
	var markerIndex int

	var markerQueue Queue
	markerMap := make(map[rune]int)

	for index, char := range signal {
		// Check if exists in the map
		_, ok := markerMap[char]

		if ok {
			for {
				signal := markerQueue.Dequeue()
				// Remove the value from the map as well
				if s, ok := signal.(rune); ok {
					delete(markerMap, s)
				}

				if signal == char {
					break
				}
			}
		}

		// Add the single signal to the maps
		markerQueue.Enqueue(char)
		markerMap[char] = index

		if len(markerMap) == 4 {
			break
		}
	}

	if !markerQueue.IsEmpty() {
		markerIndex = markerMap[markerQueue.PeakLast().(rune)] + 1
	}
	return markerIndex
}

func main() {
	file, err := os.Open("TTInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		markerIndex := getFirstMarker(scanner.Text())

		fmt.Printf("Before the first start-of-packer is detacted, %d characters need to be processed", markerIndex)
	}
}
