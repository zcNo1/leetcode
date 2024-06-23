package _331

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

const (
	FALSE = iota
	TRUE
	OR
	AND
)

func evaluateTree(root *TreeNode) bool {
	if root.Val == FALSE {
		return false
	}
	if root.Val == TRUE {
		return true
	}
	if root.Val == OR {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	if root.Val == AND {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

	return false
}
