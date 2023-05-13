package test

const (
	tianma = iota
	tianlong
	tianbing
	tianjiang
	tianwang
	xianzi
	xianwang
	shenbing
	shenshi
	shenjun
)

const (
	ruyi = iota
	wulei
	xuelinglong
)

var (
	values = []int{10, 50, 100, 300, 500, 800, 1000, 1500, 2400, 3000}
	things = [][]int{
		{tianma: 2, tianlong: 2, tianbing: 1},
		{tianma: 5, tianlong: 2, tianbing: 2},
		{tianlong: 4, tianbing: 1, tianjiang: 1, tianwang: 2}}
)

type person struct {
	want []int
	have []int
}

var shi = person{
	want: []int{5, 4, 0},
	have: []int{0, 0, 1, 2, 1, 2, 6, 0, 8, 4},
}

func getValue(p person) ([]int, int) {
	for i := 0; i < len(p.want); i++ {
		var need []int
		for i2, thing := range things[i] {
			need[i2] = thing * values[i2]
		}
	}
}
