package main

/*
 * @lc app=leetcode.cn id=927 lang=golang
 *
 * [927] 三等分
 */

// @lc code=start
func threeEqualParts(arr []int) []int {
	fail := []int{-1, -1}
	var indexOfOne []int

	// 统计1的个数
	for key, v := range arr {
		if v == 1 {
			indexOfOne = append(indexOfOne, key)
		}
	}
	// 1的个数要求可以被三等分
	if len(indexOfOne)%3 != 0 {
		return fail
	}

	// 获取三等分后每一段最后一个1的位置
	i11 := indexOfOne[len(indexOfOne)/3-1]
	i12 := indexOfOne[len(indexOfOne)/3*2-1]
	i13 := indexOfOne[len(indexOfOne)-1]

	// 先检查最后一段的1后面的0是否可以被其他段匹配
	i03 := len(arr) - 1
	d3 := i03 - i13

	if i12-i11 < d3 || i13-i12 < d3 {
		return fail
	}
	ret := fail
	ret[0] = i11 + d3
	ret[1] = i12 + d3

	// 从最后一个1往前一起推动，检查每个1的位置是否相同

	return ret
}

func main() {
	threeEqualParts([]int{})
}

// @lc code=end
