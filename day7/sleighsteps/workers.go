package sleighsteps

// Now we have to implement workers. There will be 5 workers total
// and each instruction will take 60 seconds + the letter's position
// in the alphabet (A=1, B=2, etc). So A takes 61 seconds, B takes 62 seconds, etc.

type worker struct {
	currentStep rune
	timeLeft    int
}

func (w *worker) assignStep(s rune) {
	w.currentStep = s
	// Check this out:
	w.timeLeft = 60 + int(s-'A') + 1
	// Did this for the example input, too lazy to add weird flags
	// w.timeLeft = int(s-'A') + 1

	// Clever way to assign a numerical value to a letter.
	// Since a rune is basically a number, we can just subtract
	// the value of the letter 'A' from the current step to get
	// its position in the alphabet. Then we add 1 because A should be 1, not 0.
}

// this will run once a "second"
func (w *worker) work() {
	if w.timeLeft > 0 {
		w.timeLeft--
	}
}

// easy way to tell if a worker isn't doing anything
func (w worker) isIdle() bool {
	return w.timeLeft == 0
}

// This will be for part two, because in this case we'll have to consider
// that there may be more than one potential instruction that can be worked
// on a time and could work on things in parallel.

func (inst *Instructions) GetOrderWithWorkers(workerCount int) (string, int) {
	// You wanna get nuts? Huh?
	// Let's get nuts!

	var retval string
	workerCollection := make([]worker, workerCount)

	tick := 0

	// for each tick, we need to check if any workers have completed their tasks.
	// If they have, we need to remove the completed step from the requirements of
	// any other steps. Then we need to check if any workers are idle and if there
	// are any available steps that can be worked on. If there are, we need to
	// assign them to the idle workers. Seems straightforward!
	for {
		// Check for completed tasks
		for i := range workerCollection {
			if workerCollection[i].isIdle() && workerCollection[i].currentStep != 0 {
				// This worker has completed their task, so we need to remove the
				// completed step from the requirements of any other steps.
				for _, s := range inst.manual {
					s.removeRequirement(workerCollection[i].currentStep)
				}
				// Mark the job to the return value
				retval += string(workerCollection[i].currentStep)

				// Then we need to set the worker's current step to 0 to indicate that they are idle.
				workerCollection[i].currentStep = 0
			}
		}

		// Now we'll check to see if there are any idle workers and
		// if there are any available steps that can be worked on.
		for i := range workerCollection {
			if workerCollection[i].isIdle() {
				availableSteps, err := inst.getNextSteps()
				if err != nil {
					// No available steps, so we can just continue to the next worker.
					continue
				}

				// If there are available steps, we need to assign them to the idle workers.
				// We should assign the first available step to the first idle worker, the second
				// available step to the second idle worker, etc. So we can just iterate through
				// the available steps and assign them to the idle workers until we run out of
				// either idle workers or available steps.
				nextStep := availableSteps[0]
				workerCollection[i].assignStep(nextStep)

				// Remove the assigned step from the manual so it doesn't get assigned again.
				delete(inst.manual, nextStep)
			}
		}

		// Now we need to have each worker do their work for this tick.
		for i := range workerCollection {
			workerCollection[i].work()
		}

		// Check if all work is done
		if len(inst.manual) == 0 {
			allIdle := true
			for i := range workerCollection {
				if !workerCollection[i].isIdle() {
					allIdle = false
					break
				}
			}
			if allIdle {
				// If all the work is done and we're all idle,
				// dump the last letter if one exists! THis is a
				// lot of processing for something that I'm only doing
				// because I'm a pedantic nerd
				for i := range workerCollection {
					if workerCollection[i].currentStep != 0 {
						retval += string(workerCollection[i].currentStep)
						workerCollection[i].currentStep = 0
					}
				}
				break
			}
		}

		tick++
	}

	return retval, tick + 1 // add one to account for the final tick
}
