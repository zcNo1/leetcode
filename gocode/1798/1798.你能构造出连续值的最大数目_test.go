package _798

import "testing"

func Test_getMaximumConsecutive(t *testing.T) {
	type args struct {
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{coins: []int{1, 3}},
			want: 2,
		},
		{
			name: "test2",
			args: args{coins: []int{1, 1, 1, 4}},
			want: 8,
		},
		{
			name: "test3",
			args: args{coins: []int{1, 4, 10, 3, 1}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumConsecutive(tt.args.coins); got != tt.want {
				t.Errorf("getMaximumConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
