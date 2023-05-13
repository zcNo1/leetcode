package l2441

import "testing"

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nums: []int{-1, 2, -3, 3}}, 3},
		{"test2", args{nums: []int{-1, 10, 6, 7, -7, 1}}, 7},
		{"test3", args{nums: []int{-10, 8, 6, 7, -2, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxK(tt.args.nums); got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}
