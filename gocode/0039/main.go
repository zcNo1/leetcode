package main

import (
	"fmt"
	"sort"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	cnt := make(map[byte]bool)
	bs := []byte(s)

	left := 0
	right := 1
	cnt[bs[left]] = true
	max := 1
	cur := 1

	for right < len(bs) {
		fmt.Println(cur, left, right, bs[left], bs[right], cnt)
		if val, ok := cnt[bs[right]]; !ok || val == false {
			cnt[bs[right]] = true
			cur++
			right++
		} else {
			if cur > max {
				max = cur
			}
			for bs[right] != bs[left] {
				cnt[bs[left]] = false
				left++
			}
			left++
			cur = right - left + 1
			right++
		}
		fmt.Println("---------------")
	}
	if cur > max {
		max = cur
	}

	return max
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		cur.Next = &ListNode{
			Val: (l1.Val + l2.Val + add) % 10,
		}
		add = (l1.Val + l2.Val + add) / 10
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}

	for l1 != nil {
		cur.Next = &ListNode{
			Val: (l1.Val + add) % 10,
		}
		add = (l1.Val + add) / 10
		l1 = l1.Next
		cur = cur.Next
	}

	for l2 != nil {
		cur.Next = &ListNode{
			Val: (l2.Val + add) % 10,
		}
		add = (l2.Val + add) / 10
		l2 = l2.Next
		cur = cur.Next
	}

	for add != 0 {
		cur.Next = &ListNode{
			Val: add % 10,
		}
		add = add / 10
		cur = cur.Next
	}

	return head.Next
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	var cow, col []int
	for i, ints := range matrix {
		for i2, val := range ints {
			if val == 0 {
				cow = append(cow, i)
				col = append(col, i2)
			}
		}
	}

	for _, i2 := range cow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i2][j] = 0
		}
	}

	for _, i2 := range col {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i2] = 0
		}
	}

	return
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	middle := 0
	for left <= right {
		middle = (right-left)/2 + left
		fmt.Println(left, right, middle)
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
			continue
		}
		left = middle + 1
	}

	return middle
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				cur.Next = list2
				list2 = list2.Next
			} else {
				cur.Next = list1
				list1 = list1.Next
			}
			cur = cur.Next
			continue
		}
		if list1 != nil {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
			continue
		}
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}

	return newHead.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cnt := make(map[*ListNode]bool)
	for headA != nil {
		cnt[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := cnt[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func longestConsecutive(nums []int) int {
	cnt := make(map[int]struct{})
	maxnum := int(-10e9)
	minnum := int(10e9)
	for _, num := range nums {
		cnt[num] = struct{}{}
		if num > maxnum {
			maxnum = num
		}
		if num < minnum {
			minnum = num
		}
	}
	max := 0
	cur := 0
	for i := minnum; i <= maxnum; i++ {
		if _, ok := cnt[i]; ok {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}
	if cur > max {
		max = cur
	}

	return max
}

func first(root *TreeNode) *TreeNode {
	right := root.Right
	last := root
	if root.Left != nil {
		last.Right = root.Left
		last = first(root.Left)
	}
	if right != nil {
		last.Right = right
		last = first(right)
	}
	root.Left = nil
	return last
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	first(root)
}

var cnt []int

func middle(root *TreeNode) {
	if root == nil {
		return
	}
	middle(root.Left)
	cnt = append(cnt, root.Val)
	middle(root.Right)
}

func isValidBST(root *TreeNode) bool {
	cnt = []int{}
	middle(root)
	for i := 1; i < len(cnt); i++ {
		if cnt[i] <= cnt[i-1] {
			return false
		}
	}
	return true
}

func inver(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	inver(root.Left)
	inver(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	inver(root)
	return root
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	one, two := head, head.Next
	for one != nil && two != nil && two.Next != nil {
		if one == two {
			break
		}
		one = one.Next
		two = two.Next.Next
	}
	if one != two {
		return nil
	}
	cnt := 1
	one = one.Next
	for one != two {
		one = one.Next
		cnt++
	}
	two = head
	cnt--
	for cnt != 0 {
		two = two.Next
		cnt--
	}

	one = head
	two = two.Next
	for one != two {
		one = one.Next
		two = two.Next
	}

	return one
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	one, two := head, head.Next
	for one != nil && two != nil && two.Next != nil {
		if one == two {
			return true
		}
		one = one.Next
		two = two.Next.Next
	}
	return false
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}

	cnt -= n
	if cnt == 0 {
		return head.Next
	}

	pre := head
	cur = head
	for cnt != 0 {
		pre = cur
		cur = cur.Next
		cnt--
	}
	pre.Next = cur.Next

	return head
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var com [26]int
	for _, b := range p {
		com[b-'a']++
	}

	left, right := 0, len(p)-1
	var com2 [26]int
	for idx := left; idx <= right; idx++ {
		com2[s[idx]-'a']++
	}
	var ret []int
	if com2 == com {
		ret = append(ret, left)
	}
	fmt.Println(s[left], s[right], com2)
	right++
	for right < len(s) {
		com2[s[left]-'a']--
		left++
		com2[s[right]-'a']++
		fmt.Println(s[left], s[right], com2)
		if com2 == com {
			ret = append(ret, left)
		}
		right++
	}
	return ret
}

func groupAnagrams(strs []string) [][]string {
	cnt := make(map[[26]int][]string)
	a := byte('a')

	for _, str := range strs {
		var ret [26]int
		for _, b := range []byte(str) {
			ret[b-a]++
		}
		cnt[ret] = append(cnt[ret], str)
	}

	var ret [][]string
	for _, c := range cnt {
		ret = append(ret, c)
	}

	return ret
}

func tran(candidates []int, target int) [][]int {
	var ret [][]int
	if candidates[target] == 1 {
		ret = append(ret, []int{target})
	}

	for i := 1; i <= target/2; i++ {
		ret1 := tran(candidates, i)
		ret2 := tran(candidates, target-i)
		for _, i1 := range ret1 {
			for _, i2 := range ret2 {
				ret = append(ret, append(i1, i2...))
			}
		}
	}

	return ret
}

func transfer(matrix [][]int, minN, maxN int) {
	if minN >= maxN {
		return
	}
	var idx, idy, idx2 int
	for idy = minN; idy <= maxN; idy++ {
		matrix[minN][idy], matrix[maxN][idy] = matrix[maxN][idy], matrix[minN][idy]
	}
	fmt.Println(matrix)

	for idx = minN + 1; idx < minN+(maxN-minN+1)/2; idx++ {
		idx2 = maxN - idx + minN
		matrix[idx][minN], matrix[idx2][minN] = matrix[idx2][minN], matrix[idx][minN]
		matrix[idx][maxN], matrix[idx2][maxN] = matrix[idx2][maxN], matrix[idx][maxN]
	}
	fmt.Println(matrix)

	for idy = minN + 1; idy <= maxN; idy++ {
		matrix[minN][idy], matrix[idy][minN] = matrix[idy][minN], matrix[minN][idy]
		matrix[maxN][idy], matrix[idy][maxN] = matrix[idy][maxN], matrix[maxN][idy]
	}
	fmt.Println(matrix)
	minN++
	maxN--
	transfer(matrix, minN, maxN)
}

func rotate(matrix [][]int) {
	transfer(matrix, 0, len(matrix)-1)
}
func longestCommonSubsequence(text1 string, text2 string) int {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func backTransfer(root, find *TreeNode, ret *[]*TreeNode) bool {
	if root.Val == find.Val {
		*ret = append(*ret, root)
		return true
	}

	if root.Left != nil && backTransfer(root.Left, find, ret) == true {
		*ret = append(*ret, root)
		return true
	}
	if root.Right != nil && backTransfer(root.Right, find, ret) == true {
		*ret = append(*ret, root)
		return true
	}

	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var find1, find2 []*TreeNode

	backTransfer(root, p, &find1)
	backTransfer(root, q, &find2)

	id1 := len(find1) - 1
	id2 := len(find2) - 1
	ret := root
	fmt.Println(find1)
	fmt.Println(find2)
	for id1 >= 0 && id2 >= 0 {
		if find1[id1] != find2[id2] {
			break
		}
		ret = find1[id1]
		id1--
		id2--
	}

	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	cntz := make([]*ListNode, 10001)
	cntf := make([]*ListNode, 10000)

	for i := 0; i < len(cntz); i++ {
		cntz[i] = &ListNode{}
	}
	for i := 0; i < len(cntf); i++ {
		cntf[i] = &ListNode{}
	}

	var next *ListNode
	for _, list := range lists {
		for list != nil {
			next = list.Next
			if list.Val >= 0 {
				list.Next = cntz[list.Val].Next
				cntz[list.Val].Next = list
			} else {
				list.Next = cntf[list.Val*-1].Next
				cntf[list.Val*-1].Next = list
			}
			list = next
		}
	}
	newHead := &ListNode{}
	cur := newHead
	for i := len(cntf) - 1; i >= 0; i-- {
		list := cntf[i].Next
		if list != nil {
			cur.Next = list
		}
		for list != nil {
			cur = list
			list = list.Next
		}
	}
	for i := 0; i < len(cntz); i++ {
		list := cntz[i].Next
		if list != nil {
			cur.Next = list
		}
		for list != nil {
			cur = list
			list = list.Next
		}

	}

	return newHead.Next
}

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

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode

	queue = append(queue, root)
	var ret [][]int
	var tmp []int
	pre := 1
	cur := 0
	for len(queue) != 0 {
		if queue[0] != nil {
			tmp = append(tmp, queue[0].Val)
			queue = append(queue, queue[0].Left, queue[0].Right)
			cur += 2
		}
		queue = queue[1:]
		pre--
		if pre == 0 && len(tmp) != 0 {
			ret = append(ret, tmp)
			tmp = []int{}
			pre = cur
			cur = 0
		}
	}
	return ret
}

func spiralOrder(matrix [][]int) []int {
	var ret []int
	minx, miny, maxx, maxy := 0, 0, len(matrix), len(matrix[0])
	idx, idy := 0, 0
	const (
		right = iota
		down
		left
		up
	)
	direct := right
	for idx >= minx && idx < maxx && idy >= miny && idy < maxy {
		fmt.Println(idx, idy, minx, maxx, miny, maxy)
		ret = append(ret, matrix[idx][idy])
		switch direct {
		case right:
			if idy == maxy-1 {
				idx++
				minx++
				direct = down
				continue
			}
			idy++
		case down:
			if idx == maxx-1 {
				idy--
				maxy--
				direct = left
				continue

			}
			idx++
		case left:
			if idy == miny {
				idx--
				maxx--
				direct = up
				continue
			}
			idy--

		case up:
			if idx == minx {
				idy++
				miny++
				direct = right
				continue
			}
			idx--
		}
	}
	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3

func reverseKAndCheck(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, nextHead, cnt := reverseK(head, k)
	if cnt == k {
		fmt.Println("next: ", nextHead.Val)
		//head.Next = nextHead
		head.Next = reverseKAndCheck(nextHead, k)
		return newHead
	}

	newHead, _, _ = reverseK(newHead, k)
	return newHead
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode, int) {
	if head == nil || head.Next == nil {
		return head, nil, 1
	}
	pre, next, cur := head, head.Next, head.Next
	cnt := 1
	for cur != nil && cnt < k {
		fmt.Println(pre.Val, cur.Val, next.Val, cnt)
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		cnt++
	}
	head.Next = nil
	fmt.Println("---------------")

	return pre, cur, cnt
}

func printHead(head *ListNode) {
	for head != nil {
		println(head)
		head = head.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	newHead := reverseKAndCheck(head, k)
	printHead(newHead)
	return newHead
}

func search(matrix [][]int, target int, record [][]int, idx, idy int) bool {
	if matrix[idx][idy] == target {
		return true
	}
	record[idx][idy] = 1

	if idx+1 < len(matrix) && record[idx+1][idy] != 1 && matrix[idx+1][idy] <= target {
		if search(matrix, target, record, idx+1, idy) == true {
			return true
		}
	}
	if idy+1 < len(matrix[0]) && record[idx][idy+1] != 1 && matrix[idx][idy+1] <= target {
		if search(matrix, target, record, idx, idy+1) == true {
			return true
		}
	}

	return false
}
func searchMatrix(matrix [][]int, target int) bool {
	record := make([][]int, len(matrix))
	for i := 0; i < len(record); i++ {
		record[i] = make([]int, len(matrix[0]))
	}

	return search(matrix, target, record, 0, 0)
}

func merge(intervals [][]int) [][]int {
	const (
		START = 0
		END   = 1
	)
	var ret [][]int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	ret = append(ret, intervals[0])
	cur := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][START] <= ret[cur][END] {
			if ret[cur][END] < intervals[i][END] {
				ret[cur][END] = intervals[i][END]
			}
			continue
		}
		ret = append(ret, intervals[i])
		cur++
	}
	return ret
}

func moveZeroes(nums []int) {
	cnt := 0
	idx := 0
	for num := range nums {
		if num == 0 {
			cnt++
			continue
		}
		nums[idx] = num
		idx++
	}
	for i := 0; i < cnt; i++ {
		nums[len(nums)-1-i] = 0
	}

	return
}

func combinationSum(candidates []int, target int) [][]int {
	can := make([]int, 40)
	for _, i2 := range candidates {
		can[i2] = 1
	}

	return tran(can, target)
}

func main() {

}
