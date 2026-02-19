package main

import (
	"day15/elfwar"
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
	bf := elfwar.NewBattlefield(lines)
	// Let's see what this looks like after X iterations

	// var iterations int = 26
	fmt.Printf("Initial state:\n")
	bf.Print()
	// for i := 0; i < iterations; i++ {
	// 	fmt.Printf("Iteration %d:\n", i+1)
	// 	bf.CycleOnce()
	// 	bf.Print()
	// }
	// oou := bf.OrderOfUnits()
	// fmt.Printf("Order of units: ")
	// for _, unit := range oou {
	// 	fmt.Printf("%s(%d,%d) ", unit.Team, unit.Coord.X, unit.Coord.Y)
	// }
	// fmt.Println()
	cycles := 0
	for !bf.CycleOnce() {
		// bf.Print()
		// fmt.Println()
		cycles++
	}
	bf.Print()
	fmt.Printf("Battle ended after %d cycles\n", cycles)
	fmt.Printf("HP sum: %d\n", bf.HPSum())
	fmt.Printf("Outcome: %d\n", cycles*bf.HPSum())

	// ingest the map again
	bfpart2 := elfwar.NewBattlefield(lines)
	// Now we need to find the minimum attack power for the Elves such that they win without any casualties
	attackPower := 4
finish:
	for {
		bfCopy := bfpart2.DeepCopy()
		bfCopy.SetElfAttackPower(attackPower)
		cycles := 0
		for {
			warOver, anyElvesDead := bfCopy.CycleOncePartTwo()
			if warOver {
				if !anyElvesDead {
					fmt.Printf("Elves win with attack power %d after %d cycles\n", attackPower, cycles)
					totalHP := bfCopy.HPSum()
					fmt.Printf("Outcome => total HP: %d, rounds: %d, score: %d\n", totalHP, cycles, cycles*totalHP)
					bfCopy.Print()
					break finish
				} else {
					fmt.Printf("Elves lose with attack power %d after %d cycles\n", attackPower, cycles)
				}
				break
			}
			cycles++
		}
		fmt.Printf("Trying attack power %d...\n", attackPower)
		attackPower++
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
