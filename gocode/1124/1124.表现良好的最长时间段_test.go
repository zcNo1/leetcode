package l1224

import "testing"

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{9, 9, 6, 0, 6, 6, 9}}, 3},
		{"test2", args{[]int{6, 6, 6}}, 0},
		{"test2", args{[]int{6, 6, 9}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
