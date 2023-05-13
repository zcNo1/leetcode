package _138

import (
	"fmt"
	"strings"
)

func alphabetBoardPath(target string) string {
	const LETTERLEN = 26
	const (
		UP    = 'U'
		DOWN  = 'D'
		LEFT  = 'L'
		RIGHT = 'R'
		GET   = '!'
	)

	var xl, yl [LETTERLEN]int
	for i := 0; i < LETTERLEN; i++ {
		xl[i] = i % 5
		yl[i] = i / 5
	}
	fmt.Println(xl)
	fmt.Println(yl)

	ret := strings.Builder{}
	cur := 'a'
	for _, t := range target {
		if t == cur {
			ret.WriteByte(GET)
			continue
		}
		t2 := t
		cur2 := cur
		if t == 'z' {
			t2 = 'u'
		}
		if cur == 'z' {
			cur2 = 'u'
			ret.WriteByte(UP)
		}
		x := xl[cur2-'a'] - xl[t2-'a']
		y := yl[cur2-'a'] - yl[t2-'a']
		for y < 0 {
			ret.WriteByte(DOWN)
			y++
		}
		for y > 0 {
			ret.WriteByte(UP)
			y--
		}
		for x < 0 {
			ret.WriteByte(RIGHT)
			x++
		}
		for x > 0 {
			ret.WriteByte(LEFT)
			x--
		}
		if t == 'z' {
			ret.WriteByte(DOWN)
		}
		ret.WriteByte(GET)
		cur = t
	}

	return ret.String()
}
