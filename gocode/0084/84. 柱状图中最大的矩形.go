package l0084

//

func largestRectangleArea(heights []int) int {

}

// 暴力
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
