package l0004

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{nums1: []int{1, 3}, nums2: []int{2}}, float64(2)},
		{"test2", args{nums1: []int{1, 2}, nums2: []int{3, 4}}, 2.5},
		{"test3", args{nums1: []int{1}, nums2: []int{}}, 1.0},
		{"test4", args{nums1: []int{1, 2, 3}, nums2: []int{}}, 2.0},
		{"test5", args{nums1: []int{1, 2, 3, 4, 5}, nums2: []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}}, 9.0},
		{"test6", args{nums1: []int{0, 0, 0, 0, 0}, nums2: []int{-1, 0, 0, 0, 0, 0, 1}}, float64(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
