package l0105

import . "gocode/pkg/tree"

func buildTree_(preorder []int, inorder []int, where map[int]int) *TreeNode {
	if preorder == nil || len(preorder) == 0 || inorder == nil || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	leftLen := where[preorder[0]] - where[inorder[0]]
	root.Left = buildTree_(preorder[1:leftLen+1], inorder[:leftLen], where)
	root.Right = buildTree_(preorder[leftLen+1:], inorder[leftLen+1:], where)

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	where := make(map[int]int, len(inorder))

	for i, i2 := range inorder {
		where[i2] = i
	}

	return buildTree_(preorder, inorder, where)
}
