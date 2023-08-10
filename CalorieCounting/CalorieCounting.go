package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
  content, err := os.Open("ElvesItems.txt")
  if (err != nil) {
    log.Fatal(err)
  }
  // Close the file
  defer content.Close()

  // Array of calories for each elf 
  calories := []int {0}
  currentIndex := 0
  maxIndex := 0

  scanner := bufio.NewScanner(content)
  for scanner.Scan() {
    line := scanner.Text()

    if (line !=  "") {
      calorie, err := strconv.Atoi(line)
      if err != nil {
        log.Fatal(err)
      }

      totalElfCalorie := calories[len(calories) - 1] + calorie
      // Sum the total calories for the elf 
      calories[len(calories) - 1] = totalElfCalorie 

      if totalElfCalorie >= calories[maxIndex] {
        maxIndex = currentIndex
      } 
    } else {
      calories = append(calories, 0)
      currentIndex++;
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }


  fmt.Println("The Elf " + strconv.Itoa(maxIndex + 1) + " is carrying " + strconv.Itoa(calories[maxIndex]) + " calories")
}
