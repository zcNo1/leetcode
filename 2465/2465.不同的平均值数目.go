package l2465

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	cnt := make(map[int]int)

	for i := 0; i < len(nums)/2; i++ {
		cnt[nums[i]+nums[len(nums)-1-i]]++
	}

	return len(cnt)
}
