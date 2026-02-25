package opcodes

// this is just all of the functions performed on
// the registers, taking an instruction as input
// and returning the new registers. This does not
// assume the opcode is correct, just runs this
// particular function given the instruction provided

// addition
func addr(reg registers, instr instruction) registers {
	// add register
	// stores into register C the result of adding regA and regB
	reg[instr.c] = reg[instr.a] + reg[instr.b]
	return reg
}

func addi(reg registers, instr instruction) registers {
	// add immediate
	// stores into register C the result of adding regA and value B
	reg[instr.c] = reg[instr.a] + instr.b
	return reg
}

// multiplication
func mulr(reg registers, instr instruction) registers {
	// multiply register
	// stores into register C the result of multipying regA and regB
	reg[instr.c] = reg[instr.a] * reg[instr.b]
	return reg
}

func muli(reg registers, instr instruction) registers {
	// multiply immediate
	// stores into register C the result of multipying regA and value B
	reg[instr.c] = reg[instr.a] * instr.b
	return reg
}

// bitwise AND
func banr(reg registers, instr instruction) registers {
	// bitwise AND register
	// stores into register C the result of the bitwise AND of regA and regB
	reg[instr.c] = reg[instr.a] & reg[instr.b]
	return reg
}

func bani(reg registers, instr instruction) registers {
	// bitwise AND immediate
	// stores into register C the result of the bitwise AND of regA and value B
	reg[instr.c] = reg[instr.a] & instr.b
	return reg
}

// bitwise OR
func borr(reg registers, instr instruction) registers {
	// bitwise OR register
	// stores into register C the result of the bitwise OR of regA and regB
	reg[instr.c] = reg[instr.a] | reg[instr.b]
	return reg
}

func bori(reg registers, instr instruction) registers {
	// bitwise OR immediate
	// stores into register C the result of the bitwise OR of regA and value B
	reg[instr.c] = reg[instr.a] | instr.b
	return reg
}

// assignment
func setr(reg registers, instr instruction) registers {
	// set register
	// copies the contents of regA into register C. (Input B is ignored.)
	reg[instr.c] = reg[instr.a]
	return reg
}

func seti(reg registers, instr instruction) registers {
	// set immediate
	// stores value A into register C. (Input B is ignored.)
	reg[instr.c] = instr.a
	return reg
}

// greater-than testing
func gtir(reg registers, instr instruction) registers {
	// greater-than immediate/register
	// sets register C to 1 if value A is greater than regB. Otherwise, register C is set to 0.
	if instr.a > reg[instr.b] {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}

func gtri(reg registers, instr instruction) registers {
	// greater-than register/immediate
	// sets register C to 1 if regA is greater than value B. Otherwise, register C is set to 0.
	if reg[instr.a] > instr.b {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}

func gtrr(reg registers, instr instruction) registers {
	// greater-than register/register
	// sets register C to 1 if regA is greater than regB. Otherwise, register C is set to 0.
	if reg[instr.a] > reg[instr.b] {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}

// equality testing
func eqir(reg registers, instr instruction) registers {
	// equal immediate/register
	// sets register C to 1 if value A is equal to regB. Otherwise, register C is set to 0.
	if instr.a == reg[instr.b] {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}

func eqri(reg registers, instr instruction) registers {
	// equal register/immediate
	// sets register C to 1 if regA is equal to value B. Otherwise, register C is set to 0.
	if reg[instr.a] == instr.b {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}

func eqrr(reg registers, instr instruction) registers {
	// equal register/register
	// sets register C to 1 if regA is equal to regB. Otherwise, register C is set to 0.
	if reg[instr.a] == reg[instr.b] {
		reg[instr.c] = 1
	} else {
		reg[instr.c] = 0
	}
	return reg
}
