package main

import (
	"day13/minecarts"
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

	minecartTrack := minecarts.NewTrack(lines)
	for {
		collision := minecartTrack.Tick()
		if collision.X >= 0 && collision.Y >= 0 {
			fmt.Printf("Part 1, First Collision at: %d,%d\n", collision.X, collision.Y)
			// minecarts.PrintTrack(minecartTrack)
			break
		}
	}

	// for part two, start over
	minecartTrack2 := minecarts.NewTrack(lines)
	for {
		lastCartCoord := minecartTrack2.TickPartTwo()
		if lastCartCoord.X >= 0 && lastCartCoord.Y >= 0 {
			fmt.Printf("Part 2, Last cart at: %d,%d\n", lastCartCoord.X, lastCartCoord.Y)
			break
		}
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
