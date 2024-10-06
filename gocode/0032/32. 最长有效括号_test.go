package l0032

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{s: "(()"}, 2},
		{"test2", args{s: ")()())"}, 4},
		{"test3", args{s: ""}, 0},
		{"test4", args{s: ")(()())"}, 6},
		{"test5", args{s: ")(()())()"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
