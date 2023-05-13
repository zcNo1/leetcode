package l1814

import "testing"

func Test_countNicePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{42, 11, 1, 97}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNicePairs(tt.args.nums); got != tt.want {
				t.Errorf("countNicePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_getNumSub(t *testing.T) {
//	type args struct {
//		num int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getNumSub(tt.args.num); got != tt.want {
//				t.Errorf("getNumSub() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_getRevereNum(t *testing.T) {
//	type args struct {
//		num int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getRevereNum(tt.args.num); got != tt.want {
//				t.Errorf("getRevereNum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
