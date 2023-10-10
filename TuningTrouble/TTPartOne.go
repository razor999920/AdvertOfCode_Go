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

func main() {
	file, err := os.Open("TTDemo.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
