package elfwar

import (
	"fmt"
	"strings"

	"github.com/AgroDan/aocutils"
)

// Time for my runemap library

type Battlefield struct {
	aocutils.Runemap
	Elves   []*Soldier
	Goblins []*Soldier
}

func (b *Battlefield) DeepCopy() *Battlefield {
	nb := &Battlefield{}
	nb.Runemap = *b.Runemap.DeepCopy()
	nb.Elves = make([]*Soldier, len(b.Elves))
	nb.Goblins = make([]*Soldier, len(b.Goblins))

	for i, elf := range b.Elves {
		newElf := Soldier{
			Team:   elf.Team,
			HP:     elf.HP,
			Attack: elf.Attack,
			Coord:  elf.Coord,
		}
		nb.Elves[i] = &newElf
	}

	for i, goblin := range b.Goblins {
		newGoblin := Soldier{
			Team:   goblin.Team,
			HP:     goblin.HP,
			Attack: goblin.Attack,
			Coord:  goblin.Coord,
		}
		nb.Goblins[i] = &newGoblin
	}

	return nb
}

func (b *Battlefield) IsWarOver() bool {
	elfAlive := false
	goblinAlive := false

	for _, elf := range b.Elves {
		if elf.IsAlive() {
			elfAlive = true
			break
		}
	}

	for _, goblin := range b.Goblins {
		if goblin.IsAlive() {
			goblinAlive = true
			break
		}
	}

	return !(elfAlive && goblinAlive)
}

// just a helper function for the end of the battle
func (b *Battlefield) HPSum() int {
	sum := 0
	for _, elf := range b.Elves {
		if elf.IsAlive() {
			sum += elf.HP
		}
	}

	for _, goblin := range b.Goblins {
		if goblin.IsAlive() {
			sum += goblin.HP
		}
	}

	return sum
}

func (b *Battlefield) OrderOfUnits() []*Soldier {
	var units []*Soldier
	for _, elf := range b.Elves {
		if elf.IsAlive() {
			units = append(units, elf)
		}
	}

	for _, goblin := range b.Goblins {
		if goblin.IsAlive() {
			units = append(units, goblin)
		}
	}

	// Now sort them in reading order
	// this is a stupidly inefficient bubble sort
	// but it's a small set and my head is kinda
	// spinning right now
	count := len(units)
	for i := 0; i < count; i++ {
		for j := 0; j < count-i-1; j++ {
			if units[j].Y > units[j+1].Y || (units[j].Y == units[j+1].Y && units[j].X > units[j+1].X) {
				units[j], units[j+1] = units[j+1], units[j]
			}
		}
	}
	return units
}

func NewBattlefield(input []string) *Battlefield {
	b := &Battlefield{
		Runemap: aocutils.NewRunemap(input),
	}

	// Get the Elves first
	elfLocs := b.FindAll('E')
	for _, loc := range elfLocs {
		b.Elves = append(b.Elves, &Soldier{
			Team:   "Elf",
			HP:     200,
			Attack: 3,
			Coord:  loc,
		})

		// Otherwise set the map to being empty
		// as these soldiers will be tracked seperately
		b.Set(loc, '.')
	}

	// now goblins
	goblinLocs := b.FindAll('G')
	for _, loc := range goblinLocs {
		b.Goblins = append(b.Goblins, &Soldier{
			Team:   "Goblin",
			HP:     200,
			Attack: 3,
			Coord:  loc,
		})

		b.Set(loc, '.')
	}

	return b
}

func (b *Battlefield) IsUnitPresent(coord aocutils.Coord) bool {
	for _, elf := range b.Elves {
		if elf.Coord == coord {
			return true
		}
	}

	for _, goblin := range b.Goblins {
		if goblin.Coord == coord {
			return true
		}
	}

	return false
}

func (b *Battlefield) IsUnitPresentAndAlive(coord aocutils.Coord) bool {
	// I did this because I need this exactly functionality but its
	// usage is decided in different ways in different areas.
	for _, elf := range b.Elves {
		if elf.Coord == coord && elf.IsAlive() {
			return true
		}
	}

	for _, goblin := range b.Goblins {
		if goblin.Coord == coord && goblin.IsAlive() {
			return true
		}
	}

	return false
}

// This is the function that does it all. Returns true once it's done
func (b *Battlefield) CycleOnce() bool {
	units := b.OrderOfUnits()

	// First, has the war ended?

	for _, unit := range units {
		if b.IsWarOver() {
			return true
		}

		// remember this unit may have already been killed!
		if !unit.IsAlive() {
			continue
		}
		// First, unit should have already been
		// checked for life in the b.OrderOfUnits() function
		chosenEnemy := unit.FocusAttack(b)
		if chosenEnemy == nil {
			// Then we're moving
			newCoord := unit.ChooseMove(b)
			unit.UpdateCoord(newCoord)

			// Now if we moved into range of an enemy, we should attack
			chosenEnemy = unit.FocusAttack(b)
			if chosenEnemy != nil {
				unit.AttackEnemy(chosenEnemy)
			}
		} else {
			unit.AttackEnemy(chosenEnemy)
		}
	}

	return false
}

// func (b *Battlefield) Print() {
// 	printCopy := b.DeepCopy()
// 	for _, elf := range b.Elves {
// 		if elf.IsAlive() {
// 			printCopy.Set(elf.Coord, 'E')
// 		}
// 	}

// 	for _, goblin := range b.Goblins {
// 		if goblin.IsAlive() {
// 			printCopy.Set(goblin.Coord, 'G')
// 		}
// 	}
// 	printCopy.Print()
// }

func (b *Battlefield) Print() {
	for y := 0; y < b.Height(); y++ {
		var line string
		var unitsInLine []string
		for x := 0; x < b.Width(); x++ {
			coord := aocutils.Coord{X: x, Y: y}
			if b.IsUnitPresentAndAlive(coord) {
				for _, elf := range b.Elves {
					if elf.Coord == coord && elf.IsAlive() {
						line += "E"
						unitsInLine = append(unitsInLine, "E("+fmt.Sprint(elf.HP)+")")
						break
					}
				}

				for _, goblin := range b.Goblins {
					if goblin.Coord == coord && goblin.IsAlive() {
						line += "G"
						unitsInLine = append(unitsInLine, "G("+fmt.Sprint(goblin.HP)+")")
						break
					}
				}
			} else {
				r, _ := b.Get(coord)
				line += string(r)
			}
		}

		if len(unitsInLine) > 0 {
			line += "   " + strings.Join(unitsInLine, ",")
		}

		fmt.Println(line)
	}
}

// this next part is for part 2. This will cycle over and over
// again until the elves win, basically. Also I'll need to set
// the elves attack power to a given power.

func (b *Battlefield) SetElfAttackPower(power int) {
	for _, elf := range b.Elves {
		elf.Attack = power
	}
}

func (b *Battlefield) AnyElvesDead() bool {
	for _, elf := range b.Elves {
		if !elf.IsAlive() {
			return true
		}
	}

	return false
}

func (b *Battlefield) CycleOncePartTwo() (bool, bool) {
	// returns (isWarOver, anyElvesDead)
	units := b.OrderOfUnits()

	// First, has the war ended?

	for _, unit := range units {
		if b.IsWarOver() {
			return true, b.AnyElvesDead()
		}

		// remember this unit may have already been killed!
		if !unit.IsAlive() {
			continue
		}

		// then, check if any elves are dead
		if b.AnyElvesDead() {
			return true, true
		}

		// otherwise, execute the fight as normal
		chosenEnemy := unit.FocusAttack(b)
		if chosenEnemy == nil {
			// Then we're moving
			newCoord := unit.ChooseMove(b)
			unit.UpdateCoord(newCoord)

			// Now if we moved into range of an enemy, we should attack
			chosenEnemy = unit.FocusAttack(b)
			if chosenEnemy != nil {
				unit.AttackEnemy(chosenEnemy)
			}
		} else {
			unit.AttackEnemy(chosenEnemy)
		}
	}

	return false, false
}
