package puzzle

import "fmt"

type Position struct {
	X      int
	Y      int
	MacroX int
	MacroY int
}

func (a *Position) Equal(position *Position) bool {
	return a.X == position.X && a.Y == position.Y &&
		a.MacroX == position.MacroX && a.MacroY == position.MacroY
}

func (a *Position) SameMacro(position *Position) bool {
	return a.MacroX == position.MacroX && a.MacroY == position.MacroY
}

func (p *Position) String() string {
	return fmt.Sprintf("{X: %d, Y: %d, MacroX: %d, MacroY: %d}", p.X, p.Y, p.MacroX, p.MacroY)
}
