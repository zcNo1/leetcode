package l0543

import (
	"fmt"
	. "gocode/pkg/tree"
)

func getMaxDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	cntLeft := getMaxDepth(root.Left, diameter)
	cntRight := getMaxDepth(root.Right, diameter)

	fmt.Println(root.Val, cntLeft, cntRight)

	if *diameter < cntRight+cntLeft+1 {
		*diameter = cntRight + cntLeft + 1
	}

	if cntLeft > cntRight {
		return cntLeft + 1
	}
	return cntRight + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int

	getMaxDepth(root, &diameter)
	return diameter - 1
}
