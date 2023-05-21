package LCP33

import "testing"

func Test_storeWater(t *testing.T) {
	type args struct {
		bucket []int
		vat    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3}, []int{6, 8}}, 4},
		{"test2", args{[]int{9, 0, 1}, []int{0, 2, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storeWater(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("storeWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
