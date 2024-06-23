package l1156

import "testing"

func Test_maxRepOpt1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"ababa"}, 3},
		{"test2", args{"aaabaaa"}, 6},
		{"test3", args{"aaabbaaa"}, 4},
		{"test4", args{"aaaaa"}, 5},
		{"test5", args{"abcdef"}, 1},
		{"test6", args{"aabaaabaaaba"}, 7},
		{"test6", args{"aababbb"}, 4},
		{"test7", args{"babbaaabbbbbaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepOpt1(tt.args.text); got != tt.want {
				t.Errorf("maxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
