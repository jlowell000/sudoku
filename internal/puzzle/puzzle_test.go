package puzzle

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    Puzzle
		wantErr bool
	}{
		{
			"square works",
			args{4},
			Puzzle{
				Max:       4,
				macroSize: 2,
				InitialValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				TrueValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				PotentialValues: map[Position][]int{
					{1, 1}: {0}, {2, 1}: {0}, {3, 1}: {0}, {4, 1}: {0},
					{1, 2}: {0}, {2, 2}: {0}, {3, 2}: {0}, {4, 2}: {0},
					{1, 3}: {0}, {2, 3}: {0}, {3, 3}: {0}, {4, 3}: {0},
					{1, 4}: {0}, {2, 4}: {0}, {3, 4}: {0}, {4, 4}: {0},
				},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1},
					{1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3},
					{1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
			},
			false,
		},
		{
			"non-square fails",
			args{5},
			Puzzle{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetMacro(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []Position
	}{
		{
			"gets macro",
			&Puzzle{
				Max:       4,
				macroSize: 2,
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]Position{{1, 3}, {2, 3}, {1, 4}, {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.GetMacro(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.GetMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetColumn(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []Position
	}{
		{
			"gets col",
			&Puzzle{
				Max: 4,
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]Position{{1, 3}, {2, 3}, {3, 3}, {4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.GetColumn(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.GetColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetRow(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []Position
	}{
		{
			"gets row",
			&Puzzle{
				Max: 4,
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]Position{{2, 1}, {2, 2}, {2, 3}, {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.GetRow(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetLine(t *testing.T) {
	type args struct {
		position Position
		row      bool
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []Position
	}{
		{
			"gets col",
			&Puzzle{
				Max: 4,
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}, false},
			[]Position{{1, 3}, {2, 3}, {3, 3}, {4, 3}},
		},
		{
			"gets row",
			&Puzzle{
				Max: 4,
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}, true},
			[]Position{{2, 1}, {2, 2}, {2, 3}, {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.GetLine(tt.args.position, tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.GetLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_SetPotentialValues(t *testing.T) {
	tests := []struct {
		name   string
		puzzle *Puzzle
		want   map[Position][]int
	}{
		{
			"single missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 0, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 0, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			map[Position][]int{
				{1, 1}: {1}, {2, 1}: {2}, {3, 1}: {3}, {4, 1}: {4},
				{1, 2}: {2}, {2, 2}: {3}, {3, 2}: {1}, {4, 2}: {2},
				{1, 3}: {2}, {2, 3}: {1}, {3, 3}: {4}, {4, 3}: {3},
				{1, 4}: {4}, {2, 4}: {3}, {3, 4}: {2}, {4, 4}: {1},
			},
		},
		{
			"all missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			map[Position][]int{
				{1, 1}: {1}, {2, 1}: {2}, {3, 1}: {3}, {4, 1}: {4},
				{1, 2}: {2}, {2, 2}: {3}, {3, 2}: {1}, {4, 2}: {2},
				{1, 3}: {2}, {2, 3}: {1}, {3, 3}: {4}, {4, 3}: {3},
				{1, 4}: {4}, {2, 4}: {3}, {3, 4}: {2}, {4, 4}: {1},
			},
		},
		{
			"mostly missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 2, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 2, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			map[Position][]int{
				{1, 1}: {1}, {2, 1}: {2}, {3, 1}: {3}, {4, 1}: {4},
				{1, 2}: {2}, {2, 2}: {3}, {3, 2}: {1}, {4, 2}: {2},
				{1, 3}: {2}, {2, 3}: {1}, {3, 3}: {4}, {4, 3}: {3},
				{1, 4}: {4}, {2, 4}: {3}, {3, 4}: {2}, {4, 4}: {1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.puzzle.SetPotentialValues()
		})
	}
}

func TestPuzzle_setPotentialValues(t *testing.T) {
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []int
	}{
		{
			"single missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 0, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 0, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]int{1},
		},
		{
			"all missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 0, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 0,
					{1, 4}: 0, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]int{1, 2, 3, 4},
		},
		{
			"mostly missing",
			&Puzzle{
				Max:             4,
				macroSize:       2,
				PotentialValues: map[Position][]int{},
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				InitialValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 2, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				CurrentValues: map[Position]int{
					{1, 1}: 0, {2, 1}: 2, {3, 1}: 0, {4, 1}: 0,
					{1, 2}: 0, {2, 2}: 0, {3, 2}: 0, {4, 2}: 0,
					{1, 3}: 0, {2, 3}: 0, {3, 3}: 0, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 0, {3, 4}: 0, {4, 4}: 0,
				},
				TrueValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			args{Position{2, 3}},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.puzzle.setPotentialValues(tt.args.position)
			got := tt.puzzle.PotentialValues[tt.args.position]
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.setPotentialValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_Valid(t *testing.T) {
	tests := []struct {
		name   string
		puzzle *Puzzle
		want   bool
	}{
		{
			"valid",
			&Puzzle{
				Max:       4,
				macroSize: 2,
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				CurrentValues: map[Position]int{
					{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			true},
		{
			"invalid",
			&Puzzle{
				Max:       4,
				macroSize: 2,
				Positions: []Position{
					{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
					{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
				},
				CurrentValues: map[Position]int{
					{1, 1}: 2, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
					{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
					{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
					{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 1,
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.Valid(); got != tt.want {
				t.Errorf("Puzzle.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_valid(t *testing.T) {
	puzzle := Puzzle{
		Max:       4,
		macroSize: 2,
		Positions: []Position{
			{1, 1}, {2, 1}, {3, 1}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
			{1, 3}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 4},
		},
		CurrentValues: map[Position]int{
			{1, 1}: 1, {2, 1}: 2, {3, 1}: 3, {4, 1}: 4,
			{1, 2}: 2, {2, 2}: 3, {3, 2}: 1, {4, 2}: 2,
			{1, 3}: 2, {2, 3}: 1, {3, 3}: 4, {4, 3}: 3,
			{1, 4}: 4, {2, 4}: 3, {3, 4}: 2, {4, 4}: 2,
		},
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   bool
	}{
		{"{1.1} valid", &puzzle, args{Position{1, 1}}, true},
		{"{4.4} invalid", &puzzle, args{Position{4, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.valid(tt.args.position); got != tt.want {
				t.Errorf("Puzzle.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_validGroup(t *testing.T) {
	type args struct {
		position Position
		group    []Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   bool
	}{
		{
			"diff pos; diff value; valid",
			&Puzzle{
				CurrentValues: map[Position]int{{0, 0}: 1, {0, 1}: 2, {1, 0}: 3, {1, 1}: 4},
			},
			args{Position{0, 0}, []Position{{0, 1}, {1, 0}, {1, 1}}},
			true,
		},
		{
			"shared pos; diff value; map handles this",
			&Puzzle{
				CurrentValues: map[Position]int{{0, 0}: 1, {0, 0}: 2, {1, 0}: 3, {1, 1}: 4},
			},
			args{Position{0, 0}, []Position{{0, 0}, {1, 0}, {1, 1}}},
			true,
		},
		{
			"diff pos; shared value; valid",
			&Puzzle{
				CurrentValues: map[Position]int{{0, 0}: 1, {0, 1}: 2, {1, 0}: 1, {1, 1}: 4},
			},
			args{Position{0, 0}, []Position{{0, 1}, {1, 0}, {1, 1}}},
			false,
		},
		{
			"shared pos; diff value; valid",
			&Puzzle{
				CurrentValues: map[Position]int{{0, 0}: 1, {0, 1}: 2, {1, 0}: 3, {0, 0}: 1},
			},
			args{Position{0, 0}, []Position{{0, 1}, {1, 0}, {1, 1}}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.validGroup(tt.args.position, tt.args.group); got != tt.want {
				t.Errorf("Puzzle.validGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_Conflict(t *testing.T) {
	type args struct {
		a Position
		b Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   bool
	}{
		{
			"diff pos; diff value; no conflict",
			&Puzzle{CurrentValues: map[Position]int{{0, 0}: 1, {0, 1}: 2}},
			args{Position{0, 0}, Position{0, 1}},
			false,
		},
		{
			"diff pos; same value; conflict",
			&Puzzle{CurrentValues: map[Position]int{{0, 0}: 1, {0, 1}: 1}},
			args{Position{0, 0}, Position{0, 1}},
			true,
		},
		{
			"same pos; diff value; conflict",
			&Puzzle{CurrentValues: map[Position]int{{0, 0}: 1, {0, 0}: 2}},
			args{Position{0, 0}, Position{0, 0}},
			false,
		},
		{
			"same pos; same value; no conflict",
			&Puzzle{CurrentValues: map[Position]int{{0, 0}: 1, {0, 0}: 1}},
			args{Position{0, 0}, Position{0, 0}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.Conflict(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Puzzle.Conflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_CurrentValueString(t *testing.T) {
	tests := []struct {
		name   string
		puzzle *Puzzle
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.CurrentValueString(); got != tt.want {
				t.Errorf("Puzzle.CurrentValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_ReadableCurrentValueString(t *testing.T) {
	tests := []struct {
		name   string
		puzzle *Puzzle
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.ReadableCurrentValueString(); got != tt.want {
				t.Errorf("Puzzle.ReadableCurrentValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
