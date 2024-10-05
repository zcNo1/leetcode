package l0215

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
		}, 5},
		{"test2", args{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastSecondLayer(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{6}, 2},
		{"test2", args{7}, 2},
		{"test3", args{8}, 3},
		{"test4", args{0}, 0},
		{"test5", args{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastSecondLayer(tt.args.num); got != tt.want {
				t.Errorf("getLastSecondLayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
