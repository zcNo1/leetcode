package l1080

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if getMySum(root, 0, limit) {
		root = nil
	}
	return root
}

func getMySum(root *TreeNode, parent int, limit int) bool {
	sum := parent + root.Val
	if root.Left == nil && root.Right == nil {
		return sum < limit
	}

	del := true
	if root.Left != nil {
		dl := getMySum(root.Left, sum, limit)
		if dl {
			root.Left = nil
		} else {
			del = false
		}
	}

	if root.Right != nil {
		dr := getMySum(root.Right, sum, limit)
		if dr {
			root.Right = nil
		} else {
			del = false
		}
	}

	return del
}
