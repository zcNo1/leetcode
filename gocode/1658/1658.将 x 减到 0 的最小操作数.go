package main

import "fmt"

func main() {
	fmt.Println(minOperations2([]int{3, 2, 20, 1, 1, 3}, 10))
	return
}

func minOperations(nums []int, x int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < x {
		return -1
	}

	frontSums := make([]int, len(nums)+1)
	backMap := make(map[int]int, len(nums))
	ret := -1

	// 构建front数组，每个位置存放从前往后截止当前位置的和
	for i := 0; i < len(nums); i++ {
		frontSums[i+1] = frontSums[i] + nums[i]
		if frontSums[i+1] > x {
			// 以0作为分隔符
			frontSums[i+1] = 0
			break
		}
		if frontSums[i+1] == x {
			ret = i + 1
			break
		}
	}

	// 同理，构建back数组
	for backSum, i := 0, 0; i < len(nums); i++ {
		backSum = backSum + nums[len(nums)-1-i]
		backMap[backSum] = i + 1
		if backSum > x {
			break
		}
		if backSum == x {
			if ret == -1 || ret > i+1 {
				ret = i + 1
			}
			break
		}
	}

	fmt.Println(frontSums, backMap)

	for i := 1; i < len(frontSums) && frontSums[i] != 0; i++ {
		if backIndex, ok := backMap[x-frontSums[i]]; ok {
			if ret == -1 || ret > i+backIndex {
				ret = i + backIndex
			}
		}
	}

	return ret
}

func minOperations2(nums []int, x int) int {
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < x {
		return -1
	}

	lSum := 0
	rSum := sum
	ans := n + 1

	for left, right := -1, 0; left < n; left++ {
		if left != -1 {
			lSum += nums[left]
		}
		for lSum+rSum > x && right < n {
			rSum -= nums[right]
			right++
		}
		if lSum+rSum == x {
			if ans > left+1+n-right {
				ans = left + 1 + n - right
			}
		}
	}

	if ans > n {
		return -1
	}

	return ans
}
