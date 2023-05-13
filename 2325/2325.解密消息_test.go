package _325

import "testing"

func Test_decodeMessage(t *testing.T) {
	type args struct {
		key     string
		message string
	}
	tests := []struct {
		name  string
		input args
		want  string
	}{
		{
			name: "test1",
			input: args{
				key:     "the quick brown fox jumps over the lazy dog",
				message: "vkbs bs t suepuv",
			},
			want: "this is a secret",
		},
		{
			name: "test2",
			input: args{
				key:     "eljuxhpwnyrdgtqkviszcfmabo",
				message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			},
			want: "the five boxing wizards jump quickly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeMessage(tt.input.key, tt.input.message); got != tt.want {
				t.Errorf("decodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
