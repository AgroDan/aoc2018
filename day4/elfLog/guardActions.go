package elfLog

import (
	"fmt"
	"strings"
)

// This file will section guards into their respective log entries.
// This will assume that the logs are sorted chronologically.

type GuardActions struct {
	GuardID            int
	Actions            []LogEntry
	totalMinutesAsleep int
}

// This will loop over all of the log entries and add each line to a map of GuardActions
func SectionGuardActions(logs []LogEntry) map[int]GuardActions {
	// Define the map of guards and their actions
	guardMap := make(map[int]GuardActions)
	var currentGuardID int

	for _, log := range logs {
		if strings.Contains(log.action, "begins shift") {
			// Extract the guard ID from the action string
			var guardID int
			fmt.Sscanf(log.action, "Guard #%d begins shift", &guardID)
			currentGuardID = guardID
			// Initialize the GuardActions if not already present
			if _, exists := guardMap[guardID]; !exists {
				guardMap[guardID] = GuardActions{
					GuardID: guardID,
					Actions: []LogEntry{},
				}
			}
		} else {
			// Add the log entry to the current guard's actions
			guardActions := guardMap[currentGuardID]
			guardActions.Actions = append(guardActions.Actions, log)
			guardMap[currentGuardID] = guardActions
		}
	}

	return guardMap
}

// Now I'll define all the minutes in which a guard is asleep. Since this always
// happens during the midnight hour for some reason, I can just have a map slice
// of integers and count up the amount of times a minute counts as asleep.
func GetGuardSleepMinutes(guardActions GuardActions) []int {
	sleepMinutes := make([]int, 60) // Minutes 0-59

	var sleepStart int
	for _, action := range guardActions.Actions {
		if strings.Contains(action.action, "falls asleep") {
			sleepStart = action.minute
		} else if strings.Contains(action.action, "wakes up") {
			sleepEnd := action.minute
			// Mark the minutes as asleep
			for m := sleepStart; m < sleepEnd; m++ {
				sleepMinutes[m]++
			}
		}
	}

	return sleepMinutes
}

// Total minutes asleep
func GetTotalMinutesAsleep(sleepMinutes []int) int {
	total := 0
	for _, count := range sleepMinutes {
		total += count
	}
	return total
}

// This will now process all the guards actions and return the guard ID of the
// guard with the most minutes asleep, along with their sleep minutes slice.
func FindSleepiestGuard(guardMap map[int]GuardActions) (int, []int) {
	var sleepiestGuardID int
	maxMinutesAsleep := 0
	var sleepiestGuardMinutes []int

	for guardID, actions := range guardMap {
		sleepMinutes := GetGuardSleepMinutes(actions)
		totalAsleep := GetTotalMinutesAsleep(sleepMinutes)
		if totalAsleep > maxMinutesAsleep {
			maxMinutesAsleep = totalAsleep
			sleepiestGuardID = guardID
			sleepiestGuardMinutes = sleepMinutes
		}
	}

	return sleepiestGuardID, sleepiestGuardMinutes
}

// This will find the minute that the guard is most frequently asleep on
func FindSleepiestMinute(sleepMinutes []int) (int, int) {
	sleepiestMinute := 0
	maxCount := 0

	for minute, count := range sleepMinutes {
		if count > maxCount {
			maxCount = count
			sleepiestMinute = minute
		}
	}

	return sleepiestMinute, maxCount
}

// and now find out which guard is most frequently asleep on the same minute.
func FindGuardMostFrequentlyAsleepOnSameMinute(guardMap map[int]GuardActions) (int, int, int) {
	var guardIDMostFrequentlyAsleep int
	var minuteMostFrequentlyAsleep int
	maxFrequency := 0

	for guardID, actions := range guardMap {
		sleepMinutes := GetGuardSleepMinutes(actions)
		for minute, frequency := range sleepMinutes {
			if frequency > maxFrequency {
				maxFrequency = frequency
				guardIDMostFrequentlyAsleep = guardID
				minuteMostFrequentlyAsleep = minute
			}
		}
	}

	return guardIDMostFrequentlyAsleep, minuteMostFrequentlyAsleep, maxFrequency
}
