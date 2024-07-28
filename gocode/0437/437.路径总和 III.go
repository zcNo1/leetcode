package l0437

import . "gocode/pkg/tree"

var ori int

func pathSumContain(root *TreeNode, t int) int {
	return pathSumContain(root.Left, t-root.Val) + pathSumContain(root.Right, t-root.Val)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var cnt int
	if root.Val == targetSum {
		cnt++
	}0

	pathSum(root.Left, targetSum-root.Val)
	pathSum(root.Right, targetSum-root.Val)

	cnt += pathSum(root.Left, ori) + pathSum(root.Left, targetSum-root.Val) + pathSum(root.Right, targetSum) + pathSum(root.Right, targetSum-root.Val)
	return cnt
}
