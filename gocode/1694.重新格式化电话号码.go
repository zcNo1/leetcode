package gocode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1694 lang=golang
 *
 * [1694] 重新格式化电话号码
 */

// @lc code=start
func reformatNumber(number string) string {
	newNumber := strings.ReplaceAll(strings.ReplaceAll(number, " ", ""), "-", "")
	n := len(newNumber)
	last := n % 3
	if last == 1 {
		last = 4
	}
	n -= last
	var ret []string
	var str []byte
	for key, val := range []byte(newNumber[:n]) {
		str = append(str, val)
		if (key+1)%3 == 0 {
			ret = append(ret, string(str))
			str = []byte{}
		}
	}
	if last == 4 {
		ret = append(ret, newNumber[n:n+2])
		n += 2
	}
	if n != len(newNumber) {
		ret = append(ret, newNumber[n:])
	}

	return strings.Join(ret, "-")
}

// @lc code=end
