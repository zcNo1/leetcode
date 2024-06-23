package l1090

import "testing"

func Test_largestValsFromLabels(t *testing.T) {
	type args struct {
		values    []int
		labels    []int
		numWanted int
		useLimit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2}, 12},
		{"test2", args{[]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 1}, 16},
		{"test3", args{[]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValsFromLabels(tt.args.values, tt.args.labels, tt.args.numWanted, tt.args.useLimit); got != tt.want {
				t.Errorf("largestValsFromLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
