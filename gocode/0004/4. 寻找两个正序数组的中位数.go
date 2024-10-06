package l0004

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m+n == 1 {
		if m == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	k := (m + n) / 2

	start1, end1, start2, end2 := -1, 0, -1, 0
	var cur, pre int
	for k > 1 && start1 < m-1 && start2 < n-1 {
		fmt.Println(k, start1, start2)
		end1 = start1 + k/2
		end2 = start2 + k/2
		if end1 >= m {
			end1 = m - 1
		}
		if end2 >= n {
			end2 = n - 1
		}
		if nums1[end1] > nums2[end2] {
			k -= end2 - start2
			start2 = end2
		} else {
			k -= end1 - start1
			start1 = end1
		}
	}

	if start1 == m-1 {
		pre = nums2[start2+k]
		cur = nums2[start2+k+1]
	} else if start2 == n-1 {
		pre = nums1[start1+k]
		cur = nums1[start1+k+1]
	} else {
		pre = nums1[start1+1]
		cur = nums2[start2+1]
		if pre > cur {
			pre, cur = cur, pre
			if start2+2 < n && cur > nums2[start2+2] {
				cur = nums2[start2+2]
			}
		} else {
			if start1+2 < m && cur > nums1[start1+2] {
				cur = nums1[start1+2]
			}
		}
	}

	if (m+n)%2 == 0 {
		return float64(cur+pre) / 2.0
	}
	return float64(cur)
}
