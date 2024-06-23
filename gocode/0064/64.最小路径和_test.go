package l0064

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 7},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum2(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 7},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum2(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
