package l2446

import "testing"

func Test_haveConflict(t *testing.T) {
	type args struct {
		event1 []string
		event2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]string{"01:15", "02:00"}, []string{"02:00", "03:00"}}, true},
		{"test1", args{[]string{"01:00", "02:00"}, []string{"01:20", "03:00"}}, true},
		{"test1", args{[]string{"10:00", "11:00"}, []string{"14:00", "15:00"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haveConflict(tt.args.event1, tt.args.event2); got != tt.want {
				t.Errorf("haveConflict() = %v, want %v", got, tt.want)
			}
		})
	}
}
