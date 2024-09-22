package l0124

import . "gocode/pkg/tree"

func maxPathSum(root *TreeNode) int {
	maxRet := root.Val

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		var left, right int
		lrMax := 0

		if root.Left != nil {
			left = dfs(root.Left)
			maxRet = getMax(maxRet, left)
			maxRet = getMax(maxRet, left+root.Val)
			lrMax = getMax(lrMax, left)
		}
		if root.Right != nil {
			right = dfs(root.Right)
			maxRet = getMax(maxRet, right)
			maxRet = getMax(maxRet, right+root.Val)
			lrMax = getMax(lrMax, right)
		}
		maxRet = getMax(maxRet, left+right+root.Val)

		return lrMax + root.Val
	}
	dfs(root)

	return maxRet
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
