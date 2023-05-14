package l1054

import (
	"reflect"
	"testing"
)

func Test_rearrangeBarcodes(t *testing.T) {
	type args struct {
		barcodes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 1, 1, 2, 2, 2}}, []int{}},
		{"test1", args{[]int{1, 1, 1, 1, 2, 2, 3, 3}}, []int{}},
		{"test1", args{[]int{7, 7, 7, 8, 5, 7, 5, 5, 5, 8}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeBarcodes(tt.args.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
