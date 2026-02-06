package sleighsteps

import (
	"fmt"
	"slices"
)

// F**K IT WE'LL DO IT LIVE

type step struct {
	name         rune
	requirements []rune
}

func (s *step) addRequirement(r rune) bool {
	// returns true if requirement was added
	// false if it already exists
	for _, v := range s.requirements {
		if v == r {
			return false
		}
	}

	// otherwise add the requirement
	s.requirements = append(s.requirements, r)
	return true
}

func (s *step) removeRequirement(r rune) bool {
	// returns true if requirement was removed
	delIndex := slices.Index(s.requirements, r)
	if delIndex < 0 {
		return false
	}

	s.requirements = slices.Delete(s.requirements, delIndex, delIndex+1)
	return true
}

func (s step) isAvailable() bool {
	// returns true if no requirements exist
	return len(s.requirements) == 0
}

type Instructions struct {
	manual map[rune]*step
}

func (i Instructions) getNextSteps() ([]rune, error) {
	// This function returns the runes that are considered next. This
	// supplies an ordered list of runes, so technically the next step
	// if there are more than one should be at index 0. Or rather, the
	// next step should always be at index 0.
	retval := make([]rune, 0)

	for r, s := range i.manual {
		if s.isAvailable() {
			retval = append(retval, r)
		}
	}

	if len(retval) == 0 {
		return nil, fmt.Errorf("no available steps")
	}
	slices.Sort(retval)
	return retval, nil
}

func (i *Instructions) GetFullSteps() (string, error) {
	// This function returns the full steps in order as a string. It
	// also modifies the instructions, so it should only be called once.
	retval := make([]rune, 0)

	for {
		nextSteps, err := i.getNextSteps()
		if err != nil {
			break
		}

		nextStep := nextSteps[0]
		retval = append(retval, nextStep)

		for _, s := range i.manual {
			s.removeRequirement(nextStep)
		}
		delete(i.manual, nextStep)
	}

	return string(retval), nil
}
