package l0108

import (
	. "gocode/pkg/tree"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}
	middle := len(nums) / 2
	root.Val = nums[middle]
	if middle > 0 {
		root.Left = sortedArrayToBST(nums[:middle])
	}
	if middle+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[middle+1:])
	}

	return root
}
