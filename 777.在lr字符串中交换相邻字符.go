package main

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
func canTransform(start string, end string) bool {
	var index int
	for index < len(start) {
		if start[index] == end[index] {
			index++
			continue
		}
		strings.Map()
		bytes.Index([]byte(start[index:]), end[index])
	}

	return false
}

//func main() {
//	start := "RXXLRXRXL"
//	end := "XRLXXRRLX"
//	fmt.Println(canTransform(start, end))
//}

// @lc code=end
