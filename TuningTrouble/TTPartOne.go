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

func getFirstMarker(signal string) string {
	var marker string

	var markerQueue Queue
	markerMap := make(map[rune]int)

	for _, char := range signal {
		_, ok := markerMap[char]
		if !ok {
			// Add the single signal to the maps
			markerQueue.Enqueue(char)
			markerMap[char] = 1
		} else {
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

			// Now add it to the queue
			markerQueue.Enqueue(char)
			markerMap[char] = 1
		}

		if len(markerMap) == 4 {
			break
		}

		for _, char := range markerQueue {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}

	for !markerQueue.IsEmpty() {
		signal := markerQueue.Dequeue()
		// Remove the value from the map as well
		if s, ok := signal.(rune); ok {
			delete(markerMap, s)
		}

		marker += string(signal.(rune))
	}

	return marker
}

func main() {
	file, err := os.Open("TTDemo.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		marker := getFirstMarker(scanner.Text())

		fmt.Printf("Before the first start-of-packer is detacted, %s characters need to be processed", marker)
	}
}
