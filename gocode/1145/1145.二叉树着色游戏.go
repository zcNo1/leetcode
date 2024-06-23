package _145

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

func BuildTreeNode(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	var root *TreeNode
	var nodeList []*TreeNode
	nodeList = append(nodeList, root)
	for i := 0; i < len(vals) && len(nodeList) > 0; i++ {
		tmp := nodeList
		nodeList = nil
		for _, node := range tmp {
			node = &TreeNode{
				Val: vals[i],
			}
			nodeList = append(nodeList, node.Left, node.Right)
		}
	}

	return root
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	xNode := getXNode(root, x)
	left := getNum(xNode.Left)
	if left > n/2 {
		return true
	}
	right := getNum(xNode.Right)
	if right > n/2 {
		return true
	}
	if n-right-left-1 > n/2 {
		return true
	}

	return false
}

func getXNode(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}

	node1 := getXNode(root.Left, x)
	if node1 != nil {
		return node1
	}

	return getXNode(root.Right, x)
}

func getNum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return getNum(root.Left) + getNum(root.Right) + 1
}
