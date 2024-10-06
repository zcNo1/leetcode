package l0084

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{heights: []int{2, 1, 5, 6, 2, 3}}, 10},
		{"test2", args{heights: []int{2, 4}}, 4},
		{"test3", args{heights: []int{1000, 1, 5, 6, 2, 3}}, 1000},
		{"test4", args{heights: []int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
