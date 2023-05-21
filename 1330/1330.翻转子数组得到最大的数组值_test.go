package l1330

import "testing"

func Test_maxValueAfterReverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 3, 1, 5, 4}}, 10},
		{"test1", args{[]int{2, 4, 9, 24, 2, 1, 10}}, 68},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValueAfterReverse(tt.args.nums); got != tt.want {
				t.Errorf("maxValueAfterReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
