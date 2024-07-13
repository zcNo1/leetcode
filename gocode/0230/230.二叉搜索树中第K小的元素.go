package l0230

import . "gocode/pkg/tree"

func middleTransfer(root *TreeNode, cnt *int, k int) *TreeNode {
	if root == nil {
		return nil
	}

	ret := middleTransfer(root.Left, cnt, k)
	if ret != nil {
		return ret
	}
	*cnt++
	if *cnt == k {
		return root
	}

	return middleTransfer(root.Right, cnt, k)
}

func kthSmallest(root *TreeNode, k int) int {
	var cnt int

	ret := middleTransfer(root, &cnt, k)
	if ret != nil {
		return ret.Val
	}

	return -1
}
