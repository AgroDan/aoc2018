package elfwar

import (
	"fmt"

	"github.com/AgroDan/aocutils"
)

// This will include all of the soldier-related functions
// so I can separate them and split them off from the
// battlefield related functions.

type Soldier struct {
	Team   string
	HP     int
	Attack int
	aocutils.Coord
}

// this might come in handy
func (s *Soldier) IsAlive() bool {
	return s.HP > 0
}

func (s *Soldier) IsInRange(other *Soldier) bool {
	otherRange := other.TrueAllAvailable()
	for _, point := range otherRange {
		if s.Coord == point {
			return true
		}
	}

	return false
}

func (s *Soldier) OpenInRange(b *Battlefield) []aocutils.Coord {
	// This will just return _only_ the open spaces that are in
	// range of _this_ soldier, accounting for walls _and_ other
	// soldiers. This is important

	var retval []aocutils.Coord
	soldierRange := s.TrueAllAvailable()
outer:
	for _, point := range soldierRange {
		if b.IsInBounds(point) {
			// First, check if there's a wall
			checkRune, err := b.Get(point)
			if err != nil {
				panic("this should never happen")
			}

			if checkRune != '.' {
				continue
			}

			// loop through the soldiers now
			for _, otherSoldier := range b.Elves {
				if point == otherSoldier.Coord && otherSoldier.IsAlive() {
					continue outer
				}
			}

			for _, otherSoldier := range b.Goblins {
				if point == otherSoldier.Coord && otherSoldier.IsAlive() {
					continue outer
				}
			}

			// if we made it here, this is a valid point
			retval = append(retval, point)
		}
	}

	return retval
}

// I named this "AttackEnemy" because I already named
// the attack stat as "Attack" and the code would have
// failed otherwise
func (s *Soldier) AttackEnemy(enemy *Soldier) {
	enemy.HP -= s.Attack
}

func (s *Soldier) ChooseMove(b *Battlefield) aocutils.Coord {
	// This is a little different from the previous function.
	// Same idea, but it will return a coordinate on where
	// this soldier should move to, or the current coordinate if
	// it should stay put. This will be used with the "CycleOnce()"
	// function in the battlefield, which will call this for each
	// soldier on each cycle.

	// First, we need to get all of the open spaces that are
	// in range of the enemy
	var enemyTeam []*Soldier
	if s.Team == "Elf" {
		enemyTeam = b.Goblins
	} else {
		enemyTeam = b.Elves
	}

	var targets []aocutils.Coord
	for _, enemy := range enemyTeam {
		if enemy.IsAlive() {
			targets = append(targets, enemy.OpenInRange(b)...)
		}
	}

	// If no targets, war is over
	if len(targets) == 0 {
		// fmt.Printf("No targets for %s at (%d,%d)\n", s.Team, s.Coord.X, s.Coord.Y)
		return s.Coord
	}

	// get a list of the "range coords" and determine the least
	// number of steps to get there

	type potentialAttention struct {
		Coord     aocutils.Coord
		Steps     int
		FirstStep aocutils.Coord
	}

	var potMoves = []potentialAttention{}
	for _, target := range targets {
		newCoord, steps := whichStep(s.Coord, target, b)
		if steps != -1 {
			potMoves = append(potMoves, potentialAttention{
				Coord:     target,
				Steps:     steps,
				FirstStep: newCoord,
			})
		}
	}

	// now get the lowest number of steps.
	if len(potMoves) == 0 {
		// no capable moves
		return s.Coord
	}

	lowestSteps := potMoves[0].Steps
	for _, move := range potMoves {
		if move.Steps < lowestSteps {
			lowestSteps = move.Steps
		}
	}

	var nextSteps []aocutils.Coord
	for _, move := range potMoves {
		if move.Steps == lowestSteps {
			nextSteps = append(nextSteps, move.FirstStep)
		}
	}

	// if there's a tie, use reading order
	readingOrder(&nextSteps)
	return nextSteps[0]
}

func (s *Soldier) ChooseMoveOld(b *Battlefield) aocutils.Coord {
	// Here is where the A* algorithm will come in.
	// This will return the coordinate that this soldier
	// should move to, or the current coordinate if it should
	// stay put.

	// First, we need to get all of the open spaces that are
	// in range of the enemy
	var enemyTeam []*Soldier
	if s.Team == "Elf" {
		enemyTeam = b.Goblins
	} else {
		enemyTeam = b.Elves
	}

	var targets []aocutils.Coord
	for _, enemy := range enemyTeam {
		if enemy.IsAlive() {
			targets = append(targets, enemy.OpenInRange(b)...)
		}
	}

	// If no targets, war is over
	if len(targets) == 0 {
		fmt.Printf("No targets for %s at (%d,%d)\n", s.Team, s.Coord.X, s.Coord.Y)
		return s.Coord
	}

	// now find the shortest path(s).

	allPaths := make(map[aocutils.Coord][]aocutils.Coord)
	shortestPath := -1
	for _, target := range targets {
		path, _ := aStar(s.Coord, target, b)
		if len(path) > 0 {
			allPaths[target] = path
			if shortestPath == -1 || len(path) < shortestPath {
				shortestPath = len(path)
			}
		}
	}

	// If no paths, stay put
	if shortestPath == -1 {
		fmt.Printf("No paths for %s at (%d,%d)\n", s.Team, s.Coord.X, s.Coord.Y)
		b.Print()
		return s.Coord
	}

	// I'll find the targets that are the shortest distance away,
	// and if there's a tie I'll use my helper function to determine
	// which one is in reading order
	var chosenEnemies []aocutils.Coord

	for target, path := range allPaths {
		if len(path) == shortestPath {
			chosenEnemies = append(chosenEnemies, target)
		}
	}

	readingOrder(&chosenEnemies)

	return allPaths[chosenEnemies[0]][1]
}

func (s *Soldier) FocusAttack(b *Battlefield) *Soldier {
	// This will return the enemy that this soldier should attack,
	// or nil if there are no enemies in range. This is separate
	// from the "AttackEnemy" function because the logic for
	// determining which enemy to attack is a little complex, and
	// I want to keep it separate from the actual attacking.

	var enemyTeam []*Soldier
	if s.Team == "Elf" {
		enemyTeam = b.Goblins
	} else {
		enemyTeam = b.Elves
	}

	var targets []*Soldier
	for _, enemy := range enemyTeam {
		if enemy.IsAlive() && s.IsInRange(enemy) {
			targets = append(targets, enemy)
		}
	}

	if len(targets) == 0 {
		return nil
	}

	// If there's only one target, attack it
	if len(targets) == 1 {
		return targets[0]
	}

	// If there's a tie, attack the one with the lowest HP
	lowestHP := targets[0].HP
	for _, target := range targets {
		if target.HP < lowestHP {
			lowestHP = target.HP
		}
	}

	var lowestHPTargets []*Soldier
	for _, target := range targets {
		if target.HP == lowestHP {
			lowestHPTargets = append(lowestHPTargets, target)
		}
	}

	// If there's still a tie, attack the one in reading order
	if len(lowestHPTargets) == 1 {
		return lowestHPTargets[0]
	}

	var coords []aocutils.Coord
	for _, target := range lowestHPTargets {
		coords = append(coords, target.Coord)
	}

	readingOrder(&coords)

	for _, target := range lowestHPTargets {
		if target.Coord == coords[0] {
			return target
		}
	}

	// if we got here, then logic as we know it is flawed
	panic("this should never happen")
}

func (s *Soldier) UpdateCoord(newCoord aocutils.Coord) {
	s.Coord = newCoord
}
