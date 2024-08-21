package puzzle

import "fmt"

type Cell struct {
	initialValue    int
	currentValue    int
	trueValue       int
	potentialValues []int
	Position        *Position
	Puzzle          *Puzzle
}

func (cell *Cell) GetCurrentValue() int {
	return cell.currentValue
}

func (cell *Cell) SetCurrentValue(newValue int) {
	cell.currentValue = newValue
}

func (cell *Cell) Valid() bool {
	return (cell.currentValue != cell.initialValue &&
		cell.validMacro() &&
		cell.validColumn() &&
		cell.validRow())
}

func (cell *Cell) validMacro() bool {
	return validGroup(cell, cell.Puzzle.GetMacro(*cell.Position))
}

func (cell *Cell) validColumn() bool {
	return validGroup(cell, cell.Puzzle.GetColumn(*cell.Position))
}

func (cell *Cell) validRow() bool {
	return validGroup(cell, cell.Puzzle.GetRow(*cell.Position))
}

func validGroup(cell *Cell, group []*Cell) bool {
	var conflictingCells []*Cell
	for _, c := range group {
		if cell.Conflict(c) {
			conflictingCells = append(conflictingCells, c)
		}
	}
	return len(conflictingCells) == 0
}

func (c Cell) Conflict(cell *Cell) bool {
	return !(c.currentValue == 0 ||
		cell.currentValue == 0 ||
		(c.Position.Equal(cell.Position) &&
			c.currentValue == cell.currentValue) ||
		(!c.Position.Equal(cell.Position) &&
			c.currentValue != cell.currentValue))
}

func (c *Cell) String() string {
	return fmt.Sprintf(
		"{initialValue: %d, currentValue: %d, trueValue: %d, potentialValue: %v, Position: %s}",
		c.initialValue,
		c.currentValue,
		c.trueValue,
		c.potentialValues,
		c.Position.String(),
	)
}
