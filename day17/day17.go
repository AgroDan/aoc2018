package main

import (
	"day17/sand"
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

	// challenge dictates that the source starts at X:500, Y:0
	waterSource := aocutils.Coord{X: 500, Y: 0}

	claymap := sand.GenerateMap(lines)
	// Set the initial water source to '|'
	claymap.Set(waterSource, '|')
	// claymap.PrintOffset()
	seeds := []aocutils.Coord{waterSource}
	// count := 0
	for {

		// if count >= 1000 {
		// 	break
		// }
		seeds = claymap.Tick(seeds)
		if len(seeds) == 0 {
			break
		}
		// count++
	}
	// fmt.Printf("Finished in %d ticks.\n", count)
	// fmt.Printf("Map maxes: X: %d, Y: %d\n", claymap.MaxX, claymap.MaxY)
	// fmt.Println()
	// claymap.PrintOffset()

	fmt.Printf("Total water for part 1: %d\n", claymap.CountWater())
	fmt.Printf("Total settled water for part 2: %d\n", claymap.CountSettledWater())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
