package gocode

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nums: []int{100, 4, 200, 1, 3, 2}}, 4},
		{"test2", args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, 9},
		{"test2", args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, -1}}, 10},
		{"test3", args{nums: []int{100, 4, 999999999, 1, 3, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
