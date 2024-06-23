package l2465

import "testing"

func Test_distinctAverages(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 1, 4, 0, 3, 5}}, 2},
		{"test1", args{[]int{1, 100}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctAverages(tt.args.nums); got != tt.want {
				t.Errorf("distinctAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
