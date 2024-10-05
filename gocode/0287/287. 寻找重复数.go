package l0287

func findDuplicate(nums []int) int {
	ret := nums[0]
	for ret != nums[ret] {
		nums[ret], ret = ret, nums[ret]
	}
	return ret
}
