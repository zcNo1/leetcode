package l0105

import . "gocode/pkg/tree"

func buildTree_(preorder []int, left, right int, where map[int]int) *TreeNode {
	if left > right {
		return nil
	}

	root := &TreeNode{
		Val: preorder[left],
	}

	cur := where[preorder[left]]
	root.Left = buildTree_(preorder, left, cur-1, where)
	root.Right = buildTree_(preorder, cur+1, right, where)

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	where := make(map[int]int, len(inorder))

	for i, i2 := range inorder {
		where[i2] = i
	}

}
