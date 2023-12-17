package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INSTRUCTION_ADDX_PREFIX = "addx"
const INSTRUCTIONS_NOOP_PREFIX = "noop"

func main() {
	file, err := os.Open("CRTDemoInputOne.txt")
	if err != nil {
		log.Fatal(err)
	}

	/* Close the file */
	defer file.Close()

	/* Variables */
	cpuRegister := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions := scanner.Text()

		fmt.Println(instructions)
		registerValue, err := getTotalCPUCycles(instructions)
		if err != nil {
			log.Fatal(err)
			return
		}

		cpuRegister += registerValue
	}

	fmt.Printf("The total CPU cycles are %v", cpuRegister)
}

func getTotalCPUCycles(instructions string) (int, error) {
	totalCycles := 0

	if strings.HasPrefix(instructions, INSTRUCTIONS_NOOP_PREFIX) {
		totalCycles += 1
	} else if strings.HasPrefix(instructions, INSTRUCTION_ADDX_PREFIX) {
		instructionList := strings.Fields(instructions)
		if len(instructionList) <= 1 {
			return 0, fmt.Errorf("Invalid instructions provided %c", instructionList)
		}
		registerValue, err := strconv.Atoi(instructionList[1])
		if err != nil {
			return 0, fmt.Errorf("Invalid instruction provided %c", instructionList)
		}

		for i := 0; i >= 2; i++ {

		}

		// Now add the register value
		totalCycles += registerValue
	}

	return totalCycles, nil
}
