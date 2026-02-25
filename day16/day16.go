package main

import (
	"day16/opcodes"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	// using this because the input challenge is varied so
	// i need to split this up myself
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	instructionSets := opcodes.PartOneParseInput(challengeText)

	// just to make sure this parsed properly:
	// fmt.Printf("First instruction set:\n%v\n", instructionSets[0])
	partOneSamples := 0
	for _, instrSet := range instructionSets {
		if instrSet.OpcodeMatches() >= 3 {
			partOneSamples++
		}
	}
	fmt.Printf("Part One: %d samples match 3 or more opcodes\n", partOneSamples)

	// Now let's see if any instruction set only matches 1 opcode.
	// This is just to make sure that the opcode matching is working properly.
	// for _, instrSet := range instructionSets {
	// 	if instrSet.OpcodeMatches() == 1 {
	// 		fmt.Printf("Instruction set that matches only 1 opcode:\n%v\n", instrSet)
	// 	}
	// }

	opCodeMap := opcodes.BuildOpCodeMap(instructionSets)
	fmt.Printf("Opcode map:\n%v\n", opCodeMap)

	fmt.Printf("Now run the program for part two...\n")
	partTwoInstructions := opcodes.PartTwoParseInput(challengeText)

	finalRegisters := opcodes.RunProgram(partTwoInstructions)
	fmt.Printf("Final registers: %v\n", finalRegisters)

	fmt.Printf("Part Two: value in register 0 is %d\n", finalRegisters[0])

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
