package puzzle

import (
	"testing"
)

func TestPosition_SameMacro(t *testing.T) {
	tests := []struct {
		name string
		a    *Position
		b    *Position
		want bool
	}{
		{
			"everything the same",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			true,
		},
		{
			"y different",
			&Position{X: 0, Y: 1, MacroX: 0, MacroY: 0},
			&Position{X: 0, Y: 1, MacroX: 0, MacroY: 1},
			false,
		},
		{
			"x different",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 1, Y: 0, MacroX: 1, MacroY: 0},
			false,
		},
		{
			"everything different",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 1, Y: 1, MacroX: 1, MacroY: 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SameMacro(tt.b); got != tt.want {
				t.Errorf("Position.sameMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_Equal(t *testing.T) {
	tests := []struct {
		name string
		a    *Position
		b    *Position
		want bool
	}{
		{
			"everything the same",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			true,
		},
		{
			"y different",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 0, Y: 1, MacroX: 0, MacroY: 0},
			false,
		},
		{
			"x different",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 1, Y: 0, MacroX: 0, MacroY: 0},
			false,
		},
		{
			"everything different",
			&Position{X: 0, Y: 0, MacroX: 0, MacroY: 0},
			&Position{X: 1, Y: 1, MacroX: 1, MacroY: 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("Position.sameMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Position
		want string
	}{
		{
			"normal string",
			&Position{X: 1, Y: 2, MacroX: 3, MacroY: 4},
			"{X: 1, Y: 2, MacroX: 3, MacroY: 4}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
