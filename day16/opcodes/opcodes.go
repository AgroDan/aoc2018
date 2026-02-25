package opcodes

import (
	"fmt"
	"strings"
)

// This might be a bit of a twister, but I'm first going to define
// the opcodes as if they are objects, allowing for me to work
// with the specific orders

// opcode is first number
// a is input 0
// b is input 1
// c is output
type instruction struct {
	opcode, a, b, c int
}

// now registers
type registers [4]int

type InstructionSet struct {
	before, after registers
	instr         instruction
}

func (is InstructionSet) String() string {
	return fmt.Sprintf("Before: %v\nInstruction: %v\nAfter: %v\n",
		is.before, is.instr, is.after)
}

// now I'd like to work on an instructionset to determine
// which opcodes would work for a given instruction set.
// probably easiest to just make an isEqual() function
// to determine if the registers are the same.

func (reg registers) isEqual(other registers) bool {
	for i := 0; i < len(reg); i++ {
		if reg[i] != other[i] {
			return false
		}
	}
	return true
}

func (is InstructionSet) OpcodeMatches() int {
	// this will return the number of opcodes that match
	// this instruction set.
	matches := 0

	// now run through all the opcodes and see if they match
	if is.after.isEqual(addr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(addi(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(mulr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(muli(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(banr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(bani(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(borr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(bori(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(setr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(seti(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(gtir(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(gtri(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(gtrr(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(eqir(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(eqri(is.before, is.instr)) {
		matches++
	}
	if is.after.isEqual(eqrr(is.before, is.instr)) {
		matches++
	}

	return matches
}

// Now I need to create hinter functions that will help
// me narrow down each opcode. I think the best way to do this is to
// first find the instruction set that only one opcode works for
// then i'll try and narrow down each opcode by trying this opcode
// against all others that match the same instruction set.
// If any are wrong, then it's not that opcode.

func (is InstructionSet) GetPotentialOperations() []string {
	// this will return the number of opcodes that match
	// this instruction set.
	var matches []string

	// now run through all the opcodes and see if they match
	if is.after.isEqual(addr(is.before, is.instr)) {
		matches = append(matches, "addr")
	}
	if is.after.isEqual(addi(is.before, is.instr)) {
		matches = append(matches, "addi")
	}
	if is.after.isEqual(mulr(is.before, is.instr)) {
		matches = append(matches, "mulr")
	}
	if is.after.isEqual(muli(is.before, is.instr)) {
		matches = append(matches, "muli")
	}
	if is.after.isEqual(banr(is.before, is.instr)) {
		matches = append(matches, "banr")
	}
	if is.after.isEqual(bani(is.before, is.instr)) {
		matches = append(matches, "bani")
	}
	if is.after.isEqual(borr(is.before, is.instr)) {
		matches = append(matches, "borr")
	}
	if is.after.isEqual(bori(is.before, is.instr)) {
		matches = append(matches, "bori")
	}
	if is.after.isEqual(setr(is.before, is.instr)) {
		matches = append(matches, "setr")
	}
	if is.after.isEqual(seti(is.before, is.instr)) {
		matches = append(matches, "seti")
	}
	if is.after.isEqual(gtir(is.before, is.instr)) {
		matches = append(matches, "gtir")
	}
	if is.after.isEqual(gtri(is.before, is.instr)) {
		matches = append(matches, "gtri")
	}
	if is.after.isEqual(gtrr(is.before, is.instr)) {
		matches = append(matches, "gtrr")
	}
	if is.after.isEqual(eqir(is.before, is.instr)) {
		matches = append(matches, "eqir")
	}
	if is.after.isEqual(eqri(is.before, is.instr)) {
		matches = append(matches, "eqri")
	}
	if is.after.isEqual(eqrr(is.before, is.instr)) {
		matches = append(matches, "eqrr")
	}

	return matches
}

type opCodeMap map[int]map[string]int

func (ocm opCodeMap) String() string {
	var sb strings.Builder
	for opcode, potentialOps := range ocm {
		fmt.Fprintf(&sb, "Opcode %d:\n", opcode)
		for potentialOp, count := range potentialOps {
			fmt.Fprintf(&sb, "\t%s: %d\n", potentialOp, count)
		}
	}
	return sb.String()
}

// type opCode [16]string
func BuildOpCodeMap(instructionSets []InstructionSet) opCodeMap {
	// This will be retval[opcode][potential operation] = number of times
	// this potential operation is the correct one for this opcode
	retval := make(map[int]map[string]int)

	for _, instrSet := range instructionSets {
		potentialOps := instrSet.GetPotentialOperations()
		opcode := instrSet.instr.opcode
		if _, ok := retval[opcode]; !ok {
			retval[opcode] = make(map[string]int)
		}
		for _, potentialOp := range potentialOps {
			retval[opcode][potentialOp]++
		}
	}

	return retval
}

// Based on the above I was able to find the following mapping:
// 0: borr
// 1: addr
// 2: eqrr
// 3: addi
// 4: eqri
// 5: eqir
// 6: gtri
// 7: mulr
// 8: setr
// 9: gtir
// 10: muli
// 11: banr
// 12: seti
// 13: gtrr
// 14: bani
// 15: bori

// With the above mapping that I'll explain how I was able to determine in
// the readme, I can now run through the second half of the input file and
// determine the value of register 0 after setting the registers to [0, 0, 0, 0]

func RunProgram(inst []instruction) registers {
	reg := registers{0, 0, 0, 0}

	for _, instr := range inst {
		switch instr.opcode {
		case 0:
			reg = borr(reg, instr)
		case 1:
			reg = addr(reg, instr)
		case 2:
			reg = eqrr(reg, instr)
		case 3:
			reg = addi(reg, instr)
		case 4:
			reg = eqri(reg, instr)
		case 5:
			reg = eqir(reg, instr)
		case 6:
			reg = gtri(reg, instr)
		case 7:
			reg = mulr(reg, instr)
		case 8:
			reg = setr(reg, instr)
		case 9:
			reg = gtir(reg, instr)
		case 10:
			reg = muli(reg, instr)
		case 11:
			reg = banr(reg, instr)
		case 12:
			reg = seti(reg, instr)
		case 13:
			reg = gtrr(reg, instr)
		case 14:
			reg = bani(reg, instr)
		case 15:
			reg = bori(reg, instr)
		}
	}

	return reg
}
