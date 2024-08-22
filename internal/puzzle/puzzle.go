package puzzle

import (
	"errors"
	"fmt"
	"math"
)

type Puzzle struct {
	CurrentValues   map[Position]int
	InitialValues   map[Position]int
	TrueValues      map[Position]int
	PotentialValues map[Position][]int
	Positions       []Position
	Max             int
	macroSize       int
}

func New(max int) (Puzzle, error) {
	ms := int(math.Sqrt(float64(max)))
	if ms*ms != max {
		return Puzzle{}, errors.New("max must be a square of an integer")
	}

	puzzle := Puzzle{
		Max: max, macroSize: ms,
		CurrentValues:   map[Position]int{},
		InitialValues:   map[Position]int{},
		TrueValues:      map[Position]int{},
		PotentialValues: map[Position][]int{},
	}
	for j := 1; j <= puzzle.Max; j++ {
		for i := 1; i <= puzzle.Max; i++ {
			p := Position{i, j}
			puzzle.Positions = append(puzzle.Positions, p)
			puzzle.CurrentValues[p] = 0
			puzzle.InitialValues[p] = 0
			puzzle.TrueValues[p] = 0
			puzzle.PotentialValues[p] = []int{0}
		}
	}

	return puzzle, nil
}

func (puzzle *Puzzle) GetMacro(position Position) []Position {
	var macroGroup []Position
	for _, p := range puzzle.Positions {
		if position.SameMacro(puzzle.macroSize, p) {
			macroGroup = append(macroGroup, p)
		}
	}
	return macroGroup
}

func (puzzle *Puzzle) GetColumn(position Position) []Position {
	return puzzle.GetLine(position, false)
}

func (puzzle *Puzzle) GetRow(position Position) []Position {
	return puzzle.GetLine(position, true)
}

func (puzzle *Puzzle) GetLine(position Position, row bool) []Position {
	var cells []Position
	for i := 1; i <= puzzle.Max; i++ {
		var p Position
		if row {
			p = Position{X: position.X, Y: i}
		} else {
			p = Position{X: i, Y: position.Y}
		}

		cells = append(cells, p)
	}
	return cells
}

func (puzzle *Puzzle) SetPotentialValues() {
	for _, p := range puzzle.Positions {
		puzzle.setPotentialValues(p)
	}
}

func (puzzle *Puzzle) setPotentialValues(position Position) {
	value := puzzle.CurrentValues[position]
	puzzle.PotentialValues[position] = []int{}
	if puzzle.InitialValues[position] == puzzle.TrueValues[position] {
		puzzle.PotentialValues[position] = []int{puzzle.TrueValues[position]}
	} else {
		puzzle.PotentialValues[position] = []int{}
		for i := 1; i <= puzzle.Max; i++ {
			puzzle.CurrentValues[position] = i
			if puzzle.valid(position) {
				puzzle.PotentialValues[position] = append(puzzle.PotentialValues[position], i)
			}
		}
		puzzle.CurrentValues[position] = value
	}
}

func (puzzle *Puzzle) Valid() bool {
	valid := true
	for _, p := range puzzle.Positions {
		if !puzzle.valid(p) {
			valid = false
		}
	}
	return valid
}

func (puzzle *Puzzle) valid(position Position) bool {
	return puzzle.validGroup(position, puzzle.GetMacro(position)) &&
		puzzle.validGroup(position, puzzle.GetColumn(position)) &&
		puzzle.validGroup(position, puzzle.GetRow(position))

}

func (puzzle *Puzzle) validGroup(position Position, group []Position) bool {
	var conflicting []Position
	for _, g := range group {
		if puzzle.Conflict(position, g) {
			conflicting = append(conflicting, g)
		}
	}
	return len(conflicting) == 0
}

func (puzzle *Puzzle) Conflict(a, b Position) bool {
	return !(puzzle.CurrentValues[a] == 0 ||
		puzzle.CurrentValues[b] == 0 ||
		(a.Equal(b) && puzzle.CurrentValues[a] == puzzle.CurrentValues[b]) ||
		(!a.Equal(b) && puzzle.CurrentValues[a] != puzzle.CurrentValues[b]))
}

func (puzzle *Puzzle) CurrentValueString() string {
	str := ""
	for i := 1; i <= puzzle.Max; i++ {
		for j := 1; j <= puzzle.Max; j++ {
			str = fmt.Sprintf("%s%d", str, puzzle.CurrentValues[Position{i, j}])
		}
	}
	return str
}

func (puzzle *Puzzle) ReadableCurrentValueString() string {
	str := "\n"
	for i := 1; i <= puzzle.Max; i++ {
		str = fmt.Sprintf("%s[", str)
		for j := 1; j <= puzzle.Max; j++ {
			str = fmt.Sprintf("%s%d", str, puzzle.CurrentValues[Position{i, j}])
			if j != puzzle.Max {
				str = fmt.Sprintf("%s,", str)
			} else {
				str = fmt.Sprintf("%s]", str)
			}
		}
		if i != puzzle.Max {
			str = fmt.Sprintf("%s\n", str)
		}
	}
	return str
}
