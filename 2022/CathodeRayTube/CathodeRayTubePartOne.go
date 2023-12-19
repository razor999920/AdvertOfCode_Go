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
	file, err := os.Open("CRTDemoInputTwo.txt")
	if err != nil {
		log.Fatal(err)
	}

	/* Close the file */
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	/* Variables */
	cycleCount := 0
	cpuRegisterValue := 1
	cpuSignalStrength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions := scanner.Text()

		cycles, registerValue, signalStrengthValue, err := getRegisterValueFromInstruction(instructions, cycleCount, cpuRegisterValue)
		if err != nil {
			log.Fatal(err)
			return
		}

		cycleCount = cycles
		cpuRegisterValue += registerValue
		cpuSignalStrength += signalStrengthValue
	}

	fmt.Printf("The sum of all signal strenght is %v", cpuSignalStrength)
}

func getRegisterValueFromInstruction(instruction string, currentCycles, currentRegisterValue int) (int, int, int, error) {
	cycleCount := 0
	registerValue := 0
	signalStrengthValue := 0

	if strings.HasPrefix(instruction, INSTRUCTIONS_NOOP_PREFIX) {
		cycleCount = 1
	} else if strings.HasPrefix(instruction, INSTRUCTION_ADDX_PREFIX) {
		cycleCount = 2

		/* Get register value */
		instructionValue := 0
		instructionList := strings.Fields(instruction)
		if len(instructionList) <= 1 {
			return 0, 0, 0, fmt.Errorf("Invalid instructions provided %c", instructionList)
		}
		value, err := strconv.Atoi(instructionList[1])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("Invalid instruction provided %c", instructionList)
		}

		instructionValue = value

		// Now add the register value
		registerValue += instructionValue
	}

	// Add the cycle to the total count
	currentCycles += cycleCount

	/* Check if we are gathering signal strength this iteration */
	if cycleCount == 1 {
		if currentCycles == 20 || currentCycles == 60 || currentCycles == 100 || currentCycles == 140 || currentCycles == 180 || currentCycles == 220 {
			signalStrengthValue = currentRegisterValue * currentCycles
		}
	} else if cycleCount == 2 {
		if (currentCycles-1) == 20 || (currentCycles-1) == 60 || (currentCycles-1) == 100 || (currentCycles-1) == 140 || (currentCycles-1) == 180 || (currentCycles-1) == 220 {
			signalStrengthValue = currentRegisterValue * (currentCycles - 1)
		}

		if currentCycles == 20 || currentCycles == 60 || currentCycles == 100 || currentCycles == 140 || currentCycles == 180 || currentCycles == 220 {
			signalStrengthValue = currentRegisterValue * currentCycles
		}
	}

	return currentCycles, registerValue, signalStrengthValue, nil
}
