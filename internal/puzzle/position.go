package puzzle

import "fmt"

type Position struct {
	X int
	Y int
}

func (a *Position) Equal(position Position) bool {
	return a.X == position.X && a.Y == position.Y
}

func (a *Position) SameMacro(macroSize int, position Position) bool {
	amx, amy := a.DetermineMacro(macroSize)
	bmx, bmy := position.DetermineMacro(macroSize)
	return amx == bmx && amy == bmy
}

func (p *Position) DetermineMacro(macroSize int) (int, int) {
	mx := p.X / macroSize
	my := p.Y / macroSize
	if p.X%macroSize != 0 {
		mx++
	}
	if p.Y%macroSize != 0 {
		my++
	}
	return mx, my
}

func (p *Position) String() string {
	return fmt.Sprintf("{X: %d, Y: %d}", p.X, p.Y)
}
