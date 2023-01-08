package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1441 lang=golang
 *
 * [1441] 用栈操作构建数组
 */

var notSameNum = []string{"Push", "Pop"}
var sameNum = "Push"

// @lc code=start
func buildArray(target []int, n int) []string {
	var ret []string
	it := 0
	il := 1
	for il <= n && it < len(target) {
		if target[it] != il {
			ret = append(ret, notSameNum...)
			il++
			continue
		}
		ret = append(ret, sameNum)
		it++
		il++
	}

	return ret
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 4))
}

// @lc code=end
