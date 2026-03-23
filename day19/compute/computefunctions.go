package compute

import "fmt"

// First, I'll define the registers and create the subfunctions
// for the instructionset. Note that this is basically a copy
// of the functions written from day16, with some additional
// adjustments for this particular challenge

type State struct {
	Registers     // pointing to the registers themselves
	IP        int // the intended instruction pointer register
}

type Registers [6]int

func InitRegisters() Registers {
	return Registers{0, 0, 0, 0, 0, 0}
}

func InitRegistersPartTwo() Registers {
	return Registers{1, 0, 0, 0, 0, 0}
}

func (r *Registers) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d, %d, %d]", r[0], r[1], r[2], r[3], r[4], r[5])
}

// add register
func (r *Registers) addr(a, b, c int) {
	r[c] = r[a] + r[b]
}

// add immediate
func (r *Registers) addi(a, b, c int) {
	r[c] = r[a] + b
}

// multiply register
func (r *Registers) mulr(a, b, c int) {
	r[c] = r[a] * r[b]
}

// multiply immediate
func (r *Registers) muli(a, b, c int) {
	r[c] = r[a] * b
}

// bitwise AND
func (r *Registers) banr(a, b, c int) {
	r[c] = r[a] & r[b]
}

// bitwise AND immediate
func (r *Registers) bani(a, b, c int) {
	r[c] = r[a] & b
}

// bitwise OR
func (r *Registers) borr(a, b, c int) {
	r[c] = r[a] | r[b]
}

// bitwise OR immediate
func (r *Registers) bori(a, b, c int) {
	r[c] = r[a] | b
}

// assignment, set register
func (r *Registers) setr(a, _, c int) {
	r[c] = r[a]
}

// assignment, set immediate
func (r *Registers) seti(a, _, c int) {
	r[c] = a
}

// greater than testing immediate/register
// sets register C to 1 if VALUE A > VALUE B, else 0
func (r *Registers) gtir(a, b, c int) {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

// greater than testing register/immediate
// sets register C to 1 if REGISTER A > VALUE B, else 0
func (r *Registers) gtri(a, b, c int) {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

// greater than testing register/register
// sets register C to 1 if REGISTER A > REGISTER B, else 0
func (r *Registers) gtrr(a, b, c int) {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

// equality testing immediate/register
// sets register C to 1 if VALUE A == VALUE B, else 0
func (r *Registers) eqir(a, b, c int) {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

// equality testing register/immediate
// sets register C to 1 if REGISTER A == VALUE B, else 0
func (r *Registers) eqri(a, b, c int) {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

// equality testing register/register
// sets register C to 1 if REGISTER A == REGISTER B, else 0
func (r *Registers) eqrr(a, b, c int) {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}
