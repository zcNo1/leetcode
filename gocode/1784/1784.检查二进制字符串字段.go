package gocode

/*
 * @lc app=leetcode.cn id=1784 lang=golang
 *
 * [1784] 检查二进制字符串字段
 */

// @lc code=start
func checkOnesSegment(s string) bool {
	flag := true
	for _, num := range []byte(s[1:]) {
		if num == '0' {
			flag = false
		} else {
			if flag == false {
				return false
			}
		}
	}

	return true
}

// @lc code=end

//func main() {
//	fmt.Println(checkOnesSegment("1001"))
//}
