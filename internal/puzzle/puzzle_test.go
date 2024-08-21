package puzzle

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    *Puzzle
		wantErr bool
	}{
		{
			name: "4x4",
			args: args{1, 4},
			want: &Puzzle{
				Cells: []*Cell{
					{Position: &Position{1, 1, 1, 1}},
					{Position: &Position{1, 2, 1, 1}},
					{Position: &Position{1, 3, 1, 2}},
					{Position: &Position{1, 4, 1, 2}},
					{Position: &Position{2, 1, 1, 1}},
					{Position: &Position{2, 2, 1, 1}},
					{Position: &Position{2, 3, 1, 2}},
					{Position: &Position{2, 4, 1, 2}},
					{Position: &Position{3, 1, 2, 1}},
					{Position: &Position{3, 2, 2, 1}},
					{Position: &Position{3, 3, 2, 2}},
					{Position: &Position{3, 4, 2, 2}},
					{Position: &Position{4, 1, 2, 1}},
					{Position: &Position{4, 2, 2, 1}},
					{Position: &Position{4, 3, 2, 2}},
					{Position: &Position{4, 4, 2, 2}},
				},
				Min: 1,
				Max: 4,
			},
			wantErr: false,
		},
		{
			name: "9x9; only one row checked",
			args: args{1, 9},
			want: &Puzzle{
				Cells: []*Cell{
					{Position: &Position{1, 1, 1, 1}},
					{Position: &Position{1, 2, 1, 1}},
					{Position: &Position{1, 3, 1, 1}},
					{Position: &Position{1, 4, 1, 2}},
					{Position: &Position{1, 5, 1, 2}},
					{Position: &Position{1, 6, 1, 2}},
					{Position: &Position{1, 7, 1, 3}},
					{Position: &Position{1, 8, 1, 3}},
					{Position: &Position{1, 9, 1, 3}},
				},
				Min: 1,
				Max: 4,
			},
			wantErr: false,
		},
		{
			name:    "non-Square integer errors",
			args:    args{1, 5},
			want:    &Puzzle{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, c := range tt.want.Cells {
				gotCell := got.GetCell(*c.Position)
				if !gotCell.Position.Equal(c.Position) {
					t.Errorf("New() = %v, want %v\n", gotCell.Position, c.Position)
				}
			}
		})
	}
}

func Test_determineMacro(t *testing.T) {
	type args struct {
		value     int
		macroSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1:1", args{1, 3}, 1},
		{"2:1", args{2, 3}, 1},
		{"3:1", args{3, 3}, 1},
		{"4:2", args{4, 3}, 2},
		{"5:2", args{5, 3}, 2},
		{"6:2", args{6, 3}, 2},
		{"7:3", args{7, 3}, 3},
		{"8:3", args{8, 3}, 3},
		{"9:3", args{9, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineMacro(tt.args.value, tt.args.macroSize); got != tt.want {
				t.Errorf("determineMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetCell(t *testing.T) {
	testPuzzle, _ := New(1, 4)
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   *Cell
	}{
		{
			"get's Cell",
			testPuzzle,
			args{Position{3, 3, 2, 2}},
			testPuzzle.Cells[10],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.puzzle.GetCell(tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.GetCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_GetMacro(t *testing.T) {
	testPuzzle, _ := New(1, 4)
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []*Cell
	}{
		{
			"get's Macro properly",
			testPuzzle,
			args{Position{3, 3, 2, 2}},
			[]*Cell{
				{Position: &Position{3, 3, 2, 2}},
				{Position: &Position{3, 4, 2, 2}},
				{Position: &Position{4, 3, 2, 2}},
				{Position: &Position{4, 4, 2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.puzzle.GetMacro(tt.args.position)
			for _, c := range tt.want {
				gotHas := false

				for _, g := range got {
					if g.Position.Equal(c.Position) {
						gotHas = true
					}
				}

				if !gotHas {
					t.Errorf("got:%v doesn't contain %v\n", got, c.Position)
				}
			}
		})
	}
}

func TestPuzzle_GetColumn(t *testing.T) {
	testPuzzle, _ := New(1, 4)
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []*Cell
	}{
		{
			"get's Column properly",
			testPuzzle,
			args{Position{3, 3, 2, 2}},
			[]*Cell{
				{Position: &Position{1, 3, 1, 2}},
				{Position: &Position{2, 3, 1, 2}},
				{Position: &Position{3, 3, 2, 2}},
				{Position: &Position{4, 3, 2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.puzzle.GetColumn(tt.args.position)
			for _, c := range tt.want {
				gotHas := false

				for _, g := range got {
					if g.Position.Equal(c.Position) {
						gotHas = true
					}
				}

				if !gotHas {
					t.Errorf("got:%v doesn't contain %v\n", got, c.Position)
				}
			}
		})
	}
}

func TestPuzzle_GetRow(t *testing.T) {
	testPuzzle, _ := New(1, 4)
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []*Cell
	}{
		{
			"get's Row properly",
			testPuzzle,
			args{Position{3, 3, 2, 2}},
			[]*Cell{
				{Position: &Position{3, 1, 2, 1}},
				{Position: &Position{3, 2, 2, 1}},
				{Position: &Position{3, 3, 2, 2}},
				{Position: &Position{3, 4, 2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.puzzle.GetRow(tt.args.position)
			for _, c := range tt.want {
				gotHas := false

				for _, g := range got {
					if g.Position.Equal(c.Position) {
						gotHas = true
					}
				}

				if !gotHas {
					t.Errorf("got:%v doesn't contain %v\n", got, c.Position)
				}
			}
		})
	}
}

func TestPuzzle_GetLine(t *testing.T) {
	testPuzzle, _ := New(1, 4)
	type args struct {
		position Position
		row      bool
	}

	tests := []struct {
		name   string
		puzzle *Puzzle
		args   args
		want   []*Cell
	}{
		{
			"get's row line properly",
			testPuzzle,
			args{Position{3, 3, 2, 2}, true},
			[]*Cell{
				{Position: &Position{3, 1, 2, 1}},
				{Position: &Position{3, 2, 2, 1}},
				{Position: &Position{3, 3, 2, 2}},
				{Position: &Position{3, 4, 2, 2}},
			},
		},
		{
			"get's Column properly",
			testPuzzle,
			args{Position{3, 3, 2, 2}, false},
			[]*Cell{
				{Position: &Position{1, 3, 1, 2}},
				{Position: &Position{2, 3, 1, 2}},
				{Position: &Position{3, 3, 2, 2}},
				{Position: &Position{4, 3, 2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.puzzle.GetLine(tt.args.position, tt.args.row)
			for _, c := range tt.want {
				gotHas := false

				for _, g := range got {
					if g.Position.Equal(c.Position) {
						gotHas = true
					}
				}

				if !gotHas {
					t.Errorf("got:%v doesn't contain %v\n", got, c.Position)
				}
			}
		})
	}
}

func TestPuzzle_Valid(t *testing.T) {
	testPuzzle, _ := New(1, 4)

	tests := []struct {
		name  string
		cells []*Cell
		want  bool
		want1 map[Position]bool
	}{
		{
			"no conflicts",
			[]*Cell{
				{currentValue: 1, Position: &Position{1, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 2, Position: &Position{1, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 3, Position: &Position{1, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 4, Position: &Position{1, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 3, Position: &Position{2, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 4, Position: &Position{2, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 1, Position: &Position{2, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 2, Position: &Position{2, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 2, Position: &Position{3, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 1, Position: &Position{3, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 4, Position: &Position{3, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 3, Position: &Position{3, 4, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 4, Position: &Position{4, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 3, Position: &Position{4, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 2, Position: &Position{4, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 1, Position: &Position{4, 4, 2, 2}, Puzzle: testPuzzle},
			},
			true,
			map[Position]bool{
				{1, 1, 1, 1}: true,
				{1, 2, 1, 1}: true,
				{1, 3, 1, 2}: true,
				{1, 4, 1, 2}: true,
				{2, 1, 1, 1}: true,
				{2, 2, 1, 1}: true,
				{2, 3, 1, 2}: true,
				{2, 4, 1, 2}: true,
				{3, 1, 2, 1}: true,
				{3, 2, 2, 1}: true,
				{3, 3, 2, 2}: true,
				{3, 4, 2, 2}: true,
				{4, 1, 2, 1}: true,
				{4, 2, 2, 1}: true,
				{4, 3, 2, 2}: true,
				{4, 4, 2, 2}: true,
			},
		},
		{
			"conflicts",
			[]*Cell{
				{currentValue: 1, trueValue: 1, initialValue: 0, Position: &Position{1, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 2, initialValue: 0, Position: &Position{1, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 3, initialValue: 0, Position: &Position{1, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{1, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 3, initialValue: 0, Position: &Position{2, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 4, trueValue: 4, initialValue: 0, Position: &Position{2, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 1, initialValue: 0, Position: &Position{2, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 2, initialValue: 0, Position: &Position{2, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 2, initialValue: 0, Position: &Position{3, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 1, initialValue: 0, Position: &Position{3, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 4, initialValue: 0, Position: &Position{3, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 3, initialValue: 0, Position: &Position{3, 4, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 4, trueValue: 4, initialValue: 0, Position: &Position{4, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 3, initialValue: 0, Position: &Position{4, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 2, initialValue: 0, Position: &Position{4, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 1, initialValue: 0, Position: &Position{4, 4, 2, 2}, Puzzle: testPuzzle},
			},
			false,
			map[Position]bool{
				{1, 1, 1, 1}: true,
				{1, 2, 1, 1}: true,
				{1, 3, 1, 2}: false,
				{1, 4, 1, 2}: false,
				{2, 1, 1, 1}: false,
				{2, 2, 1, 1}: true,
				{2, 3, 1, 2}: false,
				{2, 4, 1, 2}: false,
				{3, 1, 2, 1}: true,
				{3, 2, 2, 1}: false,
				{3, 3, 2, 2}: false,
				{3, 4, 2, 2}: false,
				{4, 1, 2, 1}: true,
				{4, 2, 2, 1}: true,
				{4, 3, 2, 2}: false,
				{4, 4, 2, 2}: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPuzzle.Cells = tt.cells
			got, got1 := testPuzzle.Valid()
			if got != tt.want {
				t.Errorf("Puzzle.Valid() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Puzzle.Valid() got1 = %v,\n want %v", got1, tt.want1)
			}
		})
	}
}

func Test_setPotentialValues(t *testing.T) {
	testPuzzle, _ := New(1, 4)

	tests := []struct {
		name     string
		cells    []*Cell
		position *Position
		want     []int
	}{
		{
			"empty puzzle",
			[]*Cell{
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{1, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{1, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{1, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{1, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{2, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{2, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{2, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{2, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{3, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{3, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{3, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{3, 4, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{4, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{4, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{4, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{4, 4, 2, 2}, Puzzle: testPuzzle},
			},
			&Position{1, 1, 1, 1},
			[]int{1, 2, 3, 4},
		},
		{
			"cell 1,1 certain",
			[]*Cell{
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{1, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 2, initialValue: 0, Position: &Position{1, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{1, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 4, trueValue: 4, initialValue: 0, Position: &Position{1, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{2, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{2, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{2, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{2, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 2, initialValue: 0, Position: &Position{3, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{3, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{3, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 3, trueValue: 3, initialValue: 0, Position: &Position{3, 4, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 4, trueValue: 4, initialValue: 0, Position: &Position{4, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{4, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{4, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{4, 4, 2, 2}, Puzzle: testPuzzle},
			},
			&Position{1, 1, 1, 1},
			[]int{1},
		},
		{
			"cell 1,1 certain but wrong",
			[]*Cell{
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{1, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 2, trueValue: 2, initialValue: 0, Position: &Position{1, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 3, initialValue: 0, Position: &Position{1, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 4, trueValue: 4, initialValue: 0, Position: &Position{1, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{2, 1, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{2, 2, 1, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{2, 3, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{2, 4, 1, 2}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 2, initialValue: 0, Position: &Position{3, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{3, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 4, initialValue: 0, Position: &Position{3, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{3, 4, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 1, trueValue: 4, initialValue: 0, Position: &Position{4, 1, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 3, initialValue: 0, Position: &Position{4, 2, 2, 1}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 2, initialValue: 0, Position: &Position{4, 3, 2, 2}, Puzzle: testPuzzle},
				{currentValue: 0, trueValue: 1, initialValue: 0, Position: &Position{4, 4, 2, 2}, Puzzle: testPuzzle},
			},
			&Position{1, 1, 1, 1},
			[]int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPuzzle.Cells = tt.cells
			testPuzzle.setPotentialValues(tt.position)
			got := testPuzzle.GetCell(*tt.position).potentialValues
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Puzzle.Valid() got1 = %v,\n want %v", got, tt.want)
			}
		})
	}
}
