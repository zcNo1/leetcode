package _210

import "testing"

func Test_minimumMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[][]int{
				{0, 0, 0, 0, 0, 1},
				{1, 1, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 1, 0, 1, 0},
				{0, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 0, 0}}},
			want: 11,
		},
		{
			name: "test2",
			args: args{[][]int{
				{0, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1},
				{1, 1, 0, 0, 0, 1},
				{1, 1, 1, 0, 0, 1},
				{1, 1, 1, 0, 0, 1},
				{1, 1, 1, 0, 0, 0}}},
			want: 9,
		},

		{
			name: "test3",
			args: args{[][]int{{0, 0, 0, 0, 1, 1}, {1, 1, 0, 0, 0, 1}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 0, 1, 1}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 0, 0, 0}}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.grid); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
