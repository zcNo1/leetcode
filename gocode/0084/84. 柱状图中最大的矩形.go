package l0084

// 单调栈
func largestRectangleArea(heights []int) int {
	ret := 0
	newHeights := append([]int{0}, heights...)
	newHeights = append(newHeights, 0)
	var stack []int
	stack = append(stack, 0)

	for i := 1; i < len(newHeights); {
		if newHeights[i] >= newHeights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
			continue
		}
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = getMax((i-stack[len(stack)-1]-1)*newHeights[popIndex], ret)
	}

	return ret
}

// 暴力，固定高，求宽
func largestRectangleArea2(heights []int) int {
	ret := 0
	newHeights := append([]int{0}, heights...)
	newHeights = append(newHeights, 0)
	var left, right int
	for i := 1; i < len(newHeights)-1; i++ {
		left = i - 1
		for newHeights[left] >= newHeights[i] && left > 0 {
			left--
		}
		right = i + 1
		for newHeights[right] >= newHeights[i] && right < len(newHeights)-1 {
			right++
		}
		ret = getMax(ret, (right-left-1)*newHeights[i])
	}

	return ret
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力，遍历所有节点，宽和高都不固定
func largestRectangleArea1(heights []int) int {
	ret := 0
	for i := 0; i < len(heights); i++ {
		minNum := heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] < minNum {
				minNum = heights[j]
			}
			if (j-i+1)*minNum > ret {
				ret = (j - i + 1) * minNum
			}
		}
	}

	return ret
}
