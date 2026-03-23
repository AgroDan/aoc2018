package main

import (
	"day19/compute"
	"flag"
	"fmt"
	"time"
	"utils"

	"github.com/AgroDan/aocutils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here

	instructions, err := compute.ParseInstructions(lines)
	if err != nil {
		fmt.Println("Fatal:", err)
		return
	}
	part1Regs := compute.InitRegisters()
	regs := instructions.Compute(&part1Regs)
	fmt.Printf("Final registers: %s\n", regs.String())
	fmt.Printf("Result of computation for part 1: %d\n", regs[0])

	// part2Regs := compute.InitRegistersPartTwo()
	// newRegs := instructions.Compute(&part2Regs)
	// fmt.Printf("Final registers: %s\n", newRegs.String())
	// fmt.Printf("Result of computation for part 2: %d\n", newRegs[0])

	testIter := 10000
	part2Regs := compute.InitRegistersPartTwo()
	newRegs := instructions.ComputeIteration(&part2Regs, testIter)

	fmt.Printf("Final registers after %d iterations: %s\n", testIter, newRegs.String())
	fmt.Printf("The number we are interested in is: %d\n", newRegs[4])

	divs := compute.GetDivisors(newRegs[4])
	fmt.Printf("Divisors of %d: %v\n", newRegs[4], divs)

	// I'm running out of names for variables
	// eliminate the first two numbers of divisors since it'll always start with 1 and the number itself
	// segfault if this fails i guess i dunno whatever i'm so tired
	// deezRegs := compute.SpecRegisters([6]int{newRegs[0], divs[3], divs[2], newRegs[3], newRegs[4], 3})
	// fmt.Printf("Setting registers to skip ahead to new values: %s\n", deezRegs.String())
	// finalRegs := instructions.Compute(&deezRegs)
	// fmt.Printf("dafuq\n")
	// fmt.Printf("%s\n", finalRegs.String())
	// fmt.Printf("Result of computation for part 2: %d\n", finalRegs[0])

	// it looks like thanks to this line: "addr 2 0 0" that it stores the sum of the divisors in register 0
	// so I'll just return the sum of what was already computed.

	fmt.Printf("I'm so tired, here's part 2: %d\n", aocutils.SumIntSlice(divs))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
