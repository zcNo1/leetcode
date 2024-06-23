package gocode

/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */

// @lc code=start
func minAddToMakeValid(s string) int {
	var l, r, ret int
	for _, val := range s {
		if val == '(' {
			l++
		} else {
			r++
		}
		if r > l {
			ret += r - l
			r = 0
			l = 0
		}
	}

	if l-r > 0 {
		return ret + l - r
	}

	return ret + r - l
}

//func main() {
//	fmt.Println(minAddToMakeValid("()))(("))
//}

// @lc code=end
