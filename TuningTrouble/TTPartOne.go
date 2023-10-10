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

	fmt.Println(*q)
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
		fmt.Printf("%c\n", char)

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
		}

		fmt.Println(markerMap)

		if len(markerMap) == 4 {
			break
		}
	}

	for key := range markerMap {
		marker += string(key)
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

		fmt.Printf("Before the first start-of-packer is detacted, %c characters need to be processed", marker)
	}
}
