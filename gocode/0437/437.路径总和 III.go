package l0437

import (
	"fmt"
	. "gocode/pkg/tree"
)

func pathSum(root *TreeNode, targetSum int) int {
	return pathSumContain(root, []int{0}, targetSum)
}

func pathSumContain(root *TreeNode, parentsSum []int, target int) int {
	if root == nil {
		return 0
	}

	ret := 0
	sum := parentsSum[len(parentsSum)-1] + root.Val
	for _, parentSum := range parentsSum {
		if sum-parentSum == target {
			fmt.Println(root.Val, sum, parentSum)
			ret++
		}
	}

	parentsSum = append(parentsSum, sum)
	left := pathSumContain(root.Left, parentsSum, target)
	right := pathSumContain(root.Right, append([]int{}, parentsSum...), target)

	return ret + left + right
}
