package l0152

func maxProduct(nums []int) int {
	products := make([]int, len(nums))

	ret := nums[0]
	products[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {

		}
	}

	return ret
}
