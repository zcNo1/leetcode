package l1373

import "testing"

func Test_maxSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nil}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumBST(tt.args.root); got != tt.want {
				t.Errorf("maxSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
