package l0101

import . "gocode/pkg/tree"

func compare(root, root2 *TreeNode) bool {
	if root == nil && root2 == nil {
		return true
	}

	if root == nil || root2 == nil {
		return false
	}

	if root.Val != root2.Val {
		return false
	}

	if !compare(root.Left, root2.Right) {
		return false
	}

	if !compare(root.Right, root2.Left) {
		return false
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}
