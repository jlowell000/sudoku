package puzzle

import (
	"errors"
	"math"
)

type Puzzle struct {
	Cells []*Cell
	Min   int
	Max   int
}

func New(min, max int) (*Puzzle, error) {
	macroSize := int(math.Sqrt(float64(max)))
	if macroSize*macroSize != max {
		return nil, errors.New("max must be a square of an integer")
	}

	puzzle := Puzzle{Min: min, Max: max}
	for i := puzzle.Min; i <= puzzle.Max; i++ {
		for j := puzzle.Min; j <= puzzle.Max; j++ {
			mx, my := determineMacro(i, macroSize), determineMacro(j, macroSize)
			puzzle.Cells = append(
				puzzle.Cells,
				&Cell{
					currentValue: 0,
					initialValue: 0,
					trueValue:    0,
					Position:     &Position{i, j, mx, my},
					Puzzle:       &puzzle,
				},
			)
		}
	}
	return &puzzle, nil
}

func determineMacro(value, macroSize int) int {
	macro := value / macroSize
	if value%macroSize != 0 {
		macro++
	}
	return macro
}

func (puzzle *Puzzle) GetCell(position Position) *Cell {
	for _, c := range puzzle.Cells {
		if c.Position.X == position.X && c.Position.Y == position.Y {
			return c
		}
	}
	return nil
}

func (puzzle *Puzzle) GetMacro(position Position) []*Cell {
	var cells []*Cell
	for _, c := range puzzle.Cells {
		if position.SameMacro(c.Position) {
			cells = append(cells, c)
		}
	}

	return cells
}

func (puzzle *Puzzle) GetColumn(position Position) []*Cell {
	return puzzle.GetLine(position, false)
}

func (puzzle *Puzzle) GetRow(position Position) []*Cell {
	return puzzle.GetLine(position, true)
}

func (puzzle *Puzzle) GetLine(position Position, row bool) []*Cell {
	var cells []*Cell
	for i := puzzle.Min; i <= puzzle.Max; i++ {
		var p Position
		if row {
			p = Position{X: position.X, Y: i}
		} else {
			p = Position{X: i, Y: position.Y}
		}

		cells = append(cells, puzzle.GetCell(p))
	}
	return cells
}

func (puzzle *Puzzle) setPotentialValues(position *Position) {
	cell := puzzle.GetCell(*position)
	cell.potentialValues = []int{}
	if cell.initialValue == cell.trueValue {
		cell.potentialValues = []int{cell.trueValue}
	} else {
		cell.potentialValues = []int{}
		for i := cell.Puzzle.Min; i <= cell.Puzzle.Max; i++ {
			cell.SetCurrentValue(i)
			if cell.Valid() {
				cell.potentialValues = append(cell.potentialValues, i)
			}
		}
		cell.currentValue = cell.initialValue
	}
}

func (puzzle *Puzzle) Valid() (bool, map[Position]bool) {
	valid := true
	validMap := make(map[Position]bool)
	for _, c := range puzzle.Cells {
		var v = c.Valid()
		validMap[*c.Position] = v
		if !v {
			valid = false
		}
	}
	return valid, validMap
}
