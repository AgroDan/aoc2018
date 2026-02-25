package opcodes

import (
	"fmt"
	"strings"
)

// This will just parse the input into the instruction set

func PartOneParseInput(input string) []InstructionSet {
	// because I don't know what part 2 will entail
	// and I was explicitly told to ignore the second
	// half of the input file, I'm going to toss it in
	// the garbage while I work on the first half.

	challengeText := strings.Split(input, "\n\n\n\n")[0]

	// now split this into the individual instruction sets
	instructionSetStrs := strings.Split(challengeText, "\n\n")

	instructionSets := make([]InstructionSet, len(instructionSetStrs))

	// now work on each instruction set
	for i, instrSet := range instructionSetStrs {
		createdInstructionSet := InstructionSet{
			before: registers{},
			after:  registers{},
			instr:  instruction{},
		}

		lines := strings.Split(instrSet, "\n")
		// should be 3 lines now
		fmt.Sscanf(lines[0], "Before: [%d, %d, %d, %d]",
			&createdInstructionSet.before[0],
			&createdInstructionSet.before[1],
			&createdInstructionSet.before[2],
			&createdInstructionSet.before[3])
		fmt.Sscanf(lines[1], "%d %d %d %d",
			&createdInstructionSet.instr.opcode,
			&createdInstructionSet.instr.a,
			&createdInstructionSet.instr.b,
			&createdInstructionSet.instr.c)
		fmt.Sscanf(lines[2], "After:  [%d, %d, %d, %d]",
			&createdInstructionSet.after[0],
			&createdInstructionSet.after[1],
			&createdInstructionSet.after[2],
			&createdInstructionSet.after[3])

		instructionSets[i] = createdInstructionSet

		// WHAT A MESS
	}

	return instructionSets
}

func PartTwoParseInput(input string) []instruction {
	// this is just going to parse the second half of the input file
	// into a list of instructions that we can run through the opcode map
	challengeText := strings.Split(input, "\n\n\n\n")[1]

	lines := strings.Split(challengeText, "\n")

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		var instr instruction
		fmt.Sscanf(line, "%d %d %d %d",
			&instr.opcode,
			&instr.a,
			&instr.b,
			&instr.c)
		instructions[i] = instr
	}

	return instructions
}
