package l1335

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{6, 5, 4, 3, 2, 1}, 2}, 7},
		{"test2", args{[]int{9, 9, 9}, 4}, -1},
		{"test3", args{[]int{1, 1, 1}, 3}, 3},
		{"test4", args{[]int{7, 1, 7, 1, 7, 1}, 3}, 15},
		{"test5", args{[]int{11, 111, 22, 222, 33, 333, 44, 444}, 6}, 843},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
