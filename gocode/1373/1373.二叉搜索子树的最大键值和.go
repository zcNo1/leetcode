package l1373

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

var maxSum int

func maxSumBST(root *TreeNode) int {
	maxSum = 0
	getBst(root)

	return maxSum

}

func getBst(root *TreeNode) (bool, int, int, int) {
	ret := true
	sumBST := root.Val
	min := root.Val
	max := root.Val
	if root.Left != nil {
		isBst, sumVal, minVal, maxVal := getBst(root.Left)
		if !isBst {
			ret = false
		} else {
			if maxVal >= root.Val {
				ret = false
			} else {
				min = minVal
			}
		}

		sumBST += sumVal
	}
	if root.Right != nil {
		isBst, sumVal, minVal, maxVal := getBst(root.Right)
		if !isBst {
			ret = false
		} else {
			if minVal <= root.Val {
				ret = false
			} else {
				max = maxVal
			}
		}
		sumBST += sumVal
	}
	if ret {
		if sumBST > maxSum {
			maxSum = sumBST
		}
	} else {
		sumBST = 0
	}

	return ret, sumBST, min, max
}
