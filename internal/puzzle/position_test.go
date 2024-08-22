package puzzle

import (
	"fmt"
	"testing"
)

func TestPosition_Equal(t *testing.T) {

	tests := []struct {
		name string
		a    *Position
		b    Position
		want bool
	}{
		{"same", &Position{0, 0}, Position{0, 0}, true},
		{"x diff", &Position{0, 0}, Position{1, 0}, false},
		{"y diff", &Position{0, 0}, Position{0, 1}, false},
		{"all diff", &Position{0, 0}, Position{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("Position.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_SameMacro(t *testing.T) {
	size := 3
	macroMap := map[int]int{
		1: 1, 2: 1, 3: 1,
		4: 2, 5: 2, 6: 2,
		7: 3, 8: 3, 9: 3,
	}
	type args struct {
		macroSize int
		position  Position
	}
	type test struct {
		name string
		p    *Position
		args args
		want bool
	}

	tests := []test{}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			for x := 1; x <= 9; x++ {
				for y := 1; y <= 9; y++ {
					want := macroMap[i] == macroMap[x] && macroMap[j] == macroMap[y]
					name := fmt.Sprintf("{%d,%d}; {%d,%d}; %d; %v", i, j, x, y, size, want)
					test := test{name, &Position{i, j}, args{size, Position{x, y}}, want}
					tests = append(tests, test)
				}
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SameMacro(tt.args.macroSize, tt.args.position); got != tt.want {
				t.Errorf("Position.SameMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_DetermineMacro(t *testing.T) {
	size := 3
	macroMap := map[int]int{
		1: 1, 2: 1, 3: 1,
		4: 2, 5: 2, 6: 2,
		7: 3, 8: 3, 9: 3,
	}
	type test struct {
		name      string
		p         *Position
		macroSize int
		want      int
		want1     int
	}

	tests := []test{}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			name := fmt.Sprintf("{%d,%d}%d", i, j, size)
			test := test{name, &Position{i, j}, size, macroMap[i], macroMap[j]}
			tests = append(tests, test)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.DetermineMacro(tt.macroSize)
			if got != tt.want {
				t.Errorf("Position.DetermineMacro() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Position.DetermineMacro() got1 = %v, want %v", got1, tt.want1)
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
			&Position{X: 1, Y: 2},
			"{X: 1, Y: 2}",
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
