package sleighsteps

import "fmt"

func ParseInstructions(lines []string) *Instructions {
	inst := Instructions{}
	inst.manual = make(map[rune]*step)
	for _, line := range lines {
		var req, st rune
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &req, &st)

		if _, exists := inst.manual[st]; !exists {
			inst.manual[st] = &step{name: st}
		}
		if _, exists := inst.manual[req]; !exists {
			inst.manual[req] = &step{name: req}
		}

		_ = inst.manual[st].addRequirement(req)
	}
	return &inst
}
