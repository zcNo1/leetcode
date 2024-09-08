package l0300

func lengthOfLIS(nums []int) int {
	cnt := make([]int, len(nums))

	cnt[0] = 1
	ret := 1
	for i := 1; i < len(nums); i++ {
		curCnt := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && cnt[j] > curCnt {
				curCnt = cnt[j]
			}
		}
		cnt[i] = curCnt + 1
		if cnt[i] > ret {
			ret = cnt[i]
		}
	}
	return ret
}
