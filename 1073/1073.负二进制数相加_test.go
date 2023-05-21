package l1073

import (
	"reflect"
	"testing"
)

func Test_addNegabinary(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 1, 1, 1, 1}, []int{1, 0, 1}}, []int{1, 0, 0, 0, 0}},
		{"test1", args{[]int{0}, []int{0}}, []int{0}},
		{"test1", args{[]int{0}, []int{1}}, []int{1}},
		{"test1", args{[]int{1}, []int{1}}, []int{1, 1, 0}},
		{"test1", args{[]int{1}, []int{1, 1}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addNegabinary(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addNegabinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
