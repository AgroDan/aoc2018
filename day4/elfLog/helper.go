package elfLog

import "slices"

// This will do other misc stuff of sorting the logs since we'll be getting them
// out of order.

func SortLogs(logs []LogEntry) []LogEntry {
	slices.SortFunc(logs, func(a, b LogEntry) int {
		// Need to sort by time, so year, month, day, hour, minute
		if a.year != b.year {
			return a.year - b.year
		}
		if a.month != b.month {
			return a.month - b.month
		}
		if a.day != b.day {
			return a.day - b.day
		}
		if a.hour != b.hour {
			return a.hour - b.hour
		}
		return a.minute - b.minute
	})
	return logs
}
