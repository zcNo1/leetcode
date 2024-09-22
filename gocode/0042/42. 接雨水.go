package l0042

import "fmt"

func trap(height []int) int {
	left, right := 0, 0

	// 去掉左边的0
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			left = i
			break
		}
	}
	right = left + 1
	rightMax := right

	ret := 0

	for right < len(height) {
		// 如果右边高于等于左边，则是一个凹型，计算中间水量
		if height[right] >= height[left] {
			ret += cal(height, left, right)
			left = right
			right++
			rightMax = right
			continue
		}

		// 记录右边的最大值，防止右边都比左边小
		if height[right] > height[rightMax] {
			rightMax = right
		}

		if right == len(height)-1 {
			// 如果到最后了，右边的最大值都是0，证明后续无法蓄水，直接返回
			if height[rightMax] == 0 {
				return ret
			}
			// 计算左边和右边最高直接的水量
			ret += cal(height, left, rightMax)

			// 递归
			left = rightMax
			right = left + 1
			rightMax = right
			continue
		}

		right++
	}

	return ret
}

func cal(height []int, left, right int) int {
	minArea := getMin(height[left], height[right])

	ret := 0
	for i := left + 1; i < right; i++ {
		ret += minArea - height[i]
	}

	fmt.Println(left, right, ret)

	return ret
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
