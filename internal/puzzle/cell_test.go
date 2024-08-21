package puzzle

import (
	"fmt"
	"testing"
)

func TestCell_Conflict(t *testing.T) {
	tests := []struct {
		name string
		a    *Cell
		b    *Cell
		want bool
	}{
		{
			name: "position different; value different",
			a: &Cell{
				currentValue: 1,
				Position:     &Position{X: 1, Y: 1},
			},
			b: &Cell{
				currentValue: 2,
				Position:     &Position{X: 2, Y: 2},
			},
			want: false,
		},
		{
			name: "position same; value different",
			a: &Cell{
				currentValue: 1,
				Position:     &Position{X: 1, Y: 1},
			},
			b: &Cell{
				currentValue: 2,
				Position:     &Position{X: 1, Y: 1},
			},
			want: true,
		},
		{
			name: "position different; value same",
			a: &Cell{
				currentValue: 2,
				Position:     &Position{X: 1, Y: 1},
			},
			b: &Cell{
				currentValue: 2,
				Position:     &Position{X: 2, Y: 2},
			},
			want: true,
		},
		{
			name: "position different; value non-0, 0",
			a: &Cell{
				currentValue: 1,
				Position:     &Position{X: 1, Y: 1},
			},
			b: &Cell{
				currentValue: 0,
				Position:     &Position{X: 2, Y: 2},
			},
			want: false,
		},
		{
			name: "position different; value 0, non-0",
			a: &Cell{
				currentValue: 1,
				Position:     &Position{X: 1, Y: 1},
			},
			b: &Cell{
				currentValue: 0,
				Position:     &Position{X: 2, Y: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Conflict(tt.b); got != tt.want {
				t.Errorf("Cell.Conflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_SetCurrentValue(t *testing.T) {
	type args struct {
		newValue int
	}
	tests := []struct {
		name string
		cell *Cell
		args args
		want int
	}{
		{
			"happy path",
			&Cell{currentValue: 123},
			args{321},
			321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cell.SetCurrentValue(tt.args.newValue)
			if got := tt.cell.GetCurrentValue(); got != tt.want {
				t.Errorf("Cell.GetCurrentValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Valid(t *testing.T) {
	tests := []struct {
		name string
		cell *Cell
		want bool
	}{
		// Tested by TestPuzzle_Valid
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Valid(); got != tt.want {
				t.Errorf("Cell.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_validMacro(t *testing.T) {
	tests := []struct {
		name string
		cell *Cell
		want bool
	}{
		// Tested by TestPuzzle_Valid
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.validMacro(); got != tt.want {
				t.Errorf("Cell.validMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_validColumn(t *testing.T) {
	tests := []struct {
		name string
		cell *Cell
		want bool
	}{
		// Tested by TestPuzzle_Valid
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.validColumn(); got != tt.want {
				t.Errorf("Cell.validColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_validRow(t *testing.T) {
	tests := []struct {
		name string
		cell *Cell
		want bool
	}{
		// Tested by TestPuzzle_Valid
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.validRow(); got != tt.want {
				t.Errorf("Cell.validRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validGroup(t *testing.T) {
	type args struct {
		cell  *Cell
		group []*Cell
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"no conflicts",
			args{
				&Cell{currentValue: 1, Position: &Position{1, 1, 1, 1}},
				[]*Cell{
					{currentValue: 2, Position: &Position{1, 2, 1, 1}},
					{currentValue: 3, Position: &Position{1, 3, 1, 1}},
					{currentValue: 4, Position: &Position{1, 4, 1, 1}},
				},
			},
			true,
		},
		{
			"value conflicts",
			args{
				&Cell{currentValue: 1, Position: &Position{1, 1, 1, 1}},
				[]*Cell{
					{currentValue: 1, Position: &Position{1, 2, 1, 1}},
					{currentValue: 3, Position: &Position{1, 3, 1, 1}},
					{currentValue: 4, Position: &Position{1, 4, 1, 1}},
				},
			},
			false,
		},
		{
			"position conflicts",
			args{
				&Cell{currentValue: 1, Position: &Position{1, 1, 1, 1}},
				[]*Cell{
					{currentValue: 2, Position: &Position{1, 1, 1, 1}},
					{currentValue: 3, Position: &Position{1, 3, 1, 1}},
					{currentValue: 4, Position: &Position{1, 4, 1, 1}},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validGroup(tt.args.cell, tt.args.group); got != tt.want {
				t.Errorf("validGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_String(t *testing.T) {
	tests := []struct {
		name string
		c    *Cell
		want string
	}{
		{
			"string works",
			&Cell{currentValue: 1, trueValue: 2, initialValue: 0, potentialValues: []int{123}, Position: &Position{3, 4, 5, 6}},
			fmt.Sprintf("{initialValue: 0, currentValue: 1, trueValue: 2, potentialValue: [123], Position: %s}", (&Position{3, 4, 5, 6}).String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Cell.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_GetCurrentValue(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
		want int
	}{
		{
			"happy path",
			Cell{currentValue: 123},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.GetCurrentValue(); got != tt.want {
				t.Errorf("Cell.GetCurrentValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
