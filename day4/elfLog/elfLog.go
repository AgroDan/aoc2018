package elfLog

import "fmt"

/*
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
*/

type LogEntry struct {
	year, month, day, hour, minute int
	action                         string
}

func ParseLogEntry(line string) LogEntry {
	l := LogEntry{}
	fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &l.year, &l.month, &l.day, &l.hour, &l.minute)
	l.action = line[19:]
	return l
}

func (l LogEntry) String() string {
	return fmt.Sprintf("[%04d-%02d-%02d %02d:%02d] %s", l.year, l.month, l.day, l.hour, l.minute, l.action)
}
