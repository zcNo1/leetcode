package gocode

import (
	"fmt"
	"sort"
)

// 11. 盛最多水的容器
func maxArea(height []int) int {
	var left, right int
	var max, cur int
	left = 0
	right = len(height) - 1
	for left < right {
		if height[left] <= height[right] {
			cur = (right - left) * height[left]
			if max < cur {
				max = cur
			}
			left++
		} else {
			cur = (right - left) * height[right]
			if max < cur {
				max = cur
			}
			right--
		}
	}

	return max
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int

	for i := 0; i < len(nums)-2; {
		left, right, ans := i+1, len(nums)-1, 0-nums[i]
		for left < right {
			if nums[left]+nums[right] < ans {
				left++
			} else if nums[left]+nums[right] > ans {
				right--
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < len(nums)-1 {
					if nums[left] != nums[left+1] {
						left++
						break
					}
					left++
				}
				right--
			}
		}
		for i < len(nums)-2 {
			if nums[i] != nums[i+1] {
				i++
				break
			}
			i++
		}
	}

	return ret
}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	const (
		GROUPNUM = 100000
		GROUPLEN = 100000
	)

	const (
		MINIDX = 0
		MAXIDX = 1
		MAXNUM = 1000000000
		MINNUM = -1000000000
	)

	fmt.Println(MAXNUM)

	var minmaxnumRecord [GROUPNUM][2]int
	var cntRecord [GROUPNUM]map[int]struct{}

	for i := 0; i < GROUPNUM; i++ {
		minmaxnumRecord[i][MINIDX] = MAXNUM + MAXNUM
		minmaxnumRecord[i][MAXIDX] = MINNUM + MAXNUM
	}

	for _, num := range nums {
		num += MAXNUM
		groupIdx := num / GROUPLEN
		if cntRecord[groupIdx] == nil {
			cntRecord[groupIdx] = make(map[int]struct{})
		}

		cntRecord[groupIdx][num] = struct{}{}
		if num > minmaxnumRecord[groupIdx][MAXIDX] {
			minmaxnumRecord[groupIdx][MAXIDX] = num
		}
		if num < minmaxnumRecord[groupIdx][MINIDX] {
			minmaxnumRecord[groupIdx][MINIDX] = num
		}
	}
	max := 0
	cur := 0
	lastMax := MINNUM - 1
	for i := 0; i < GROUPNUM; i++ {
		if cntRecord[i] == nil {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}
		if lastMax != minmaxnumRecord[i][MINIDX]-1 {
			if cur > max {
				max = cur
			}
			cur = 0
		}
		lastMax = minmaxnumRecord[i][MAXIDX]

		for j := minmaxnumRecord[i][MINIDX]; j <= minmaxnumRecord[i][MAXIDX]; j++ {
			if _, ok := cntRecord[i][j]; ok {
				cur++
			} else {
				if cur > max {
					max = cur
				}
				cur = 0
			}
		}
	}
	if cur > max {
		max = cur
	}

	return max
}

// 560. 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	ret := 0
	cnt := make([]int, len(nums))
	last := 0
	for i := 0; i < len(nums); i++ {
		cnt[i] = last + nums[i]
		last = cnt[i]
		if cnt[i] == k {
			ret++
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			cnt[j] -= nums[i]
			if cnt[j] == k {
				ret++
			}
		}
	}

	return ret
}

// 53. 最大子数组和
func maxSubArray(nums []int) int {
	cnt := make([]int, len(nums))
	cnt[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if cnt[i-1] < 0 {
			cnt[i] = nums[i]
		} else {
			cnt[i] = cnt[i-1] + nums[i]
		}
		if cnt[i] > max {
			max = cnt[i]
		}
	}
	return max
}
