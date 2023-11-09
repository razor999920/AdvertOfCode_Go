package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("NSLInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	directorySizes := make(map[string]int)
	var currentDirectory string
	directoryStack := []string{"/"} // Initialize with root directory

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue // Skip invalid lines
		}

		command := parts[0]
		operation := parts[1]

		// Ignore the dir command
		if command == "dir" {
			continue
		}

		switch operation {
		case "cd":
			directory := parts[2]
			if directory == "/" {
				// Move to root directory
				currentDirectory = "/"
				directoryStack = []string{"/"}
			} else if directory == ".." {
				// Move up one level
				if len(directoryStack) > 1 {
					directoryStack = directoryStack[:len(directoryStack)-1]
					currentDirectory = directoryStack[len(directoryStack)-1]
				}
			} else {
				// Move to a subdirectory
				if len(directoryStack) <= 1 {
					currentDirectory += directory
				} else {
					currentDirectory += "/" + directory
				}
				directoryStack = append(directoryStack, currentDirectory)
			}
		case "ls":
			// Skip "ls" commands
		default:
			// Parse file size and update directory size
			size, err := strconv.Atoi(command)
			if err != nil {
				log.Printf("Error parsing size: %v", err)
				continue
			}
			directorySizes[currentDirectory] += size

			// Add the size to the total size of the root diractory
			directorySizes["/"] += size
			// Also add the size to the parent diractories
			parentDiractories := strings.Split(currentDirectory, "/")
			if len(parentDiractories) > 2 {
				for _, dir := range parentDiractories {
					if dir == "" {
						continue
					}

					if dir != string(currentDirectory[len(currentDirectory)-1]) {
						dir = "/" + dir
						directorySizes[dir] += size
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the sum of total sizes of directories <= 100000
	largetTotalSize := 0
	for _, size := range directorySizes {
		if size <= 100000 && largetTotalSize < size {
			largetTotalSize = size
		}
	}

	fmt.Println("The sum of total sizes of directories is: ", largetTotalSize)
}
