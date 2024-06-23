package _129

import (
	"reflect"
	"testing"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	type args struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				n:         2,
				redEdges:  [][]int{{0, 1}, {1, 0}},
				blueEdges: [][]int{{0, 1}, {1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "test2",
			args: args{
				n:         3,
				redEdges:  [][]int{{0, 1}},
				blueEdges: [][]int{{2, 1}},
			},
			want: []int{0, 1, -1},
		},
		{
			name: "test3",
			args: args{
				n:         3,
				redEdges:  [][]int{{1, 0}},
				blueEdges: [][]int{{2, 1}},
			},
			want: []int{0, -1, -1},
		},
		{
			name: "test4",
			args: args{
				n:         3,
				redEdges:  [][]int{{0, 1}},
				blueEdges: [][]int{{1, 2}},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "test5",
			args: args{
				n:         3,
				redEdges:  [][]int{{0, 1}, {0, 2}},
				blueEdges: [][]int{{1, 0}},
			},
			want: []int{0, 1, 1},
		},
		{
			name: "test6",
			args: args{
				n:         5,
				redEdges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
				blueEdges: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: []int{0, 1, 2, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestAlternatingPaths(tt.args.n, tt.args.redEdges, tt.args.blueEdges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestAlternatingPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
