package l1110

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

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	cntDelete := make(map[int]int)
	for _, val := range to_delete {
		cntDelete[val]++
	}

	var ret []*TreeNode

	transferNodes(root, cntDelete, true, &ret)

	return ret
}

func transferNodes(root *TreeNode, cntDelete map[int]int, isFatherDelete bool, ret *[]*TreeNode) bool {
	if root == nil {
		return true
	}

	deleteMe := false
	if _, ok := cntDelete[root.Val]; ok {
		deleteMe = true
	}

	if isFatherDelete && !deleteMe {
		*ret = append(*ret, root)
	}

	if transferNodes(root.Left, cntDelete, deleteMe, ret) == true {
		root.Left = nil
	}

	if transferNodes(root.Right, cntDelete, deleteMe, ret) == true {
		root.Right = nil
	}

	return deleteMe
}
