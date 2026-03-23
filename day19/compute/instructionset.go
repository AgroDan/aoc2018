package compute

import "fmt"

// This will be the object created for the challenge itself, indexing
// each instruction so that the program flow can be determined.

type Instruction struct {
	Opcode  string
	A, B, C int
}

type InstructionSet struct {
	IP           int // the intended instruction pointer register
	Instructions []Instruction
}

func ParseInstructions(lines []string) (InstructionSet, error) {
	var iset InstructionSet
	for i, line := range lines {
		if i == 0 {
			// first line is the instruction pointer register
			_, err := fmt.Sscanf(line, "#ip %d", &iset.IP)
			if err != nil {
				return iset, fmt.Errorf("failed to parse instruction pointer from line: %s", line)
			}
		} else {
			var instr Instruction
			_, err := fmt.Sscanf(line, "%s %d %d %d", &instr.Opcode, &instr.A, &instr.B, &instr.C)
			if err != nil {
				return iset, fmt.Errorf("failed to parse instruction from line: %s", line)
			}
			iset.Instructions = append(iset.Instructions, instr)
		}
	}
	return iset, nil
}

func (iset *InstructionSet) Compute(regs *Registers) *Registers {
	// used to init the registers here, but I'm leaving that
	// responsibility to the caller so I can do part 2 "easily"

	// execute the instructions until we go out of bounds
	for {
		ip := regs[iset.IP]
		if ip < 0 || ip >= len(iset.Instructions) {
			break
		}
		instr := iset.Instructions[ip]

		// fmt.Printf("ip=%d %s %d %d %d -> ", ip, instr.Opcode, instr.A, instr.B, instr.C)

		// execute the instruction based on the opcode
		switch instr.Opcode {
		case "addr":
			regs.addr(instr.A, instr.B, instr.C)
		case "addi":
			regs.addi(instr.A, instr.B, instr.C)
		case "mulr":
			regs.mulr(instr.A, instr.B, instr.C)
		case "muli":
			regs.muli(instr.A, instr.B, instr.C)
		case "banr":
			regs.banr(instr.A, instr.B, instr.C)
		case "bani":
			regs.bani(instr.A, instr.B, instr.C)
		case "borr":
			regs.borr(instr.A, instr.B, instr.C)
		case "bori":
			regs.bori(instr.A, instr.B, instr.C)
		case "setr":
			regs.setr(instr.A, 0, instr.C) // B is ignored for setr
		case "seti":
			regs.seti(instr.A, 0, instr.C) // B is ignored for seti
		case "gtir":
			regs.gtir(instr.A, instr.B, instr.C)
		case "gtri":
			regs.gtri(instr.A, instr.B, instr.C)
		case "gtrr":
			regs.gtrr(instr.A, instr.B, instr.C)
		case "eqir":
			regs.eqir(instr.A, instr.B, instr.C)
		case "eqri":
			regs.eqri(instr.A, instr.B, instr.C)
		case "eqrr":
			regs.eqrr(instr.A, instr.B, instr.C)
		default:
			panic(fmt.Sprintf("unknown opcode: %s", instr.Opcode))
		}

		// increment the instruction pointer
		regs[iset.IP]++

		// fmt.Printf("%s\n", regs.String())
	}

	return regs
}

func GetDivisors(n int) []int {
	divisors := []int{}

	// this is funky but the i*i basically means we only
	// loop until the square root of n, then we add both i
	// and n/i as divisors when we find a match
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	return divisors
}

func (iset *InstructionSet) ComputeIteration(regs *Registers, iter int) *Registers {
	// used to init the registers here, but I'm leaving that
	// responsibility to the caller so I can do part 2 "easily"

	// execute the instructions until we go out of bounds
	for range iter {
		ip := regs[iset.IP]
		if ip < 0 || ip >= len(iset.Instructions) {
			break
		}
		instr := iset.Instructions[ip]

		// fmt.Printf("ip=%d %s %d %d %d -> ", ip, instr.Opcode, instr.A, instr.B, instr.C)

		// execute the instruction based on the opcode
		switch instr.Opcode {
		case "addr":
			regs.addr(instr.A, instr.B, instr.C)
		case "addi":
			regs.addi(instr.A, instr.B, instr.C)
		case "mulr":
			regs.mulr(instr.A, instr.B, instr.C)
		case "muli":
			regs.muli(instr.A, instr.B, instr.C)
		case "banr":
			regs.banr(instr.A, instr.B, instr.C)
		case "bani":
			regs.bani(instr.A, instr.B, instr.C)
		case "borr":
			regs.borr(instr.A, instr.B, instr.C)
		case "bori":
			regs.bori(instr.A, instr.B, instr.C)
		case "setr":
			regs.setr(instr.A, 0, instr.C) // B is ignored for setr
		case "seti":
			regs.seti(instr.A, 0, instr.C) // B is ignored for seti
		case "gtir":
			regs.gtir(instr.A, instr.B, instr.C)
		case "gtri":
			regs.gtri(instr.A, instr.B, instr.C)
		case "gtrr":
			regs.gtrr(instr.A, instr.B, instr.C)
		case "eqir":
			regs.eqir(instr.A, instr.B, instr.C)
		case "eqri":
			regs.eqri(instr.A, instr.B, instr.C)
		case "eqrr":
			regs.eqrr(instr.A, instr.B, instr.C)
		default:
			panic(fmt.Sprintf("unknown opcode: %s", instr.Opcode))
		}

		// increment the instruction pointer
		regs[iset.IP]++

		// fmt.Printf("%s\n", regs.String())
	}

	return regs
}

func SpecRegisters(regValues [6]int) Registers {
	return Registers{regValues[0], regValues[1], regValues[2], regValues[3], regValues[4], regValues[5]}
}
