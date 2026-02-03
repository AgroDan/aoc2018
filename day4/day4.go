package main

import (
	"day4/elfLog"
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
	elfLogs := []elfLog.LogEntry{}
	for _, line := range lines {
		elfLogEntry := elfLog.ParseLogEntry(line)
		elfLogs = append(elfLogs, elfLogEntry)
	}

	elfLogs = elfLog.SortLogs(elfLogs)

	// Now that the logs are sorted, seciton them by guard
	guardActionsMap := elfLog.SectionGuardActions(elfLogs)

	// for each guard, get their sleep minutes
	for guardID, guardActions := range guardActionsMap {
		sleepMinutes := elfLog.GetGuardSleepMinutes(guardActions)
		fmt.Printf("Guard #%d sleep minutes: %v\n", guardID, sleepMinutes)
	}

	// Find the sleepiest guard now...
	sleepyGuardID, sleepyGuardMinutes := elfLog.FindSleepiestGuard(guardActionsMap)
	fmt.Printf("Sleepiest Guard is #%d with total sleep minutes: %d\n", sleepyGuardID, sleepyGuardMinutes)

	// Now find the minute they are most frequently asleep
	minute, frequency := elfLog.FindSleepiestMinute(sleepyGuardMinutes)
	fmt.Printf("Sleepiest minute for Guard #%d is minute %d with frequency %d\n", sleepyGuardID, minute, frequency)

	fmt.Printf("Part 1 Answer: %d\n", sleepyGuardID*minute)

	// Now to find the answer to part 2, find the guard who is most frequently asleep on
	// the same minute. I have a function for that!

	guardID2, minute2, frequency2 := elfLog.FindGuardMostFrequentlyAsleepOnSameMinute(guardActionsMap)
	fmt.Printf("Guard most frequently asleep on same minute is Guard #%d on minute %d with frequency %d\n", guardID2, minute2, frequency2)
	fmt.Printf("Part 2 Answer: %d\n", guardID2*minute2)

	// End of code

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
