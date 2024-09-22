package l0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	x := len(nums1) + len(nums2)
	if x%2 == 1 {
		return float64(findXNum(nums1, nums2, x/2))
	} else {
		return float64(findXNum(nums1, nums2, x/2)+findXNum(nums1, nums2, x/2-1)) / 2
	}
}

func findXNum(nums1, nums2 []int, x int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 == 0 {
		return nums2[x]
	}

	if n2 == 0 {
		return nums1[x]
	}

	// 不相交
	if nums2[0] > nums1[n1-1] {
		if x < n1 {
			return nums1[x]
		} else {
			return nums2[x-n1]
		}
	}

	if nums1[0] > nums2[n2-1] {
		if x < n2 {
			return nums2[x]
		} else {
			return nums1[x-n2]
		}
	}

	// 相交
	m1 := n1 / 2
	m2 := n2 / 2
	if nums1[m1] == nums2[m2] {
		if x == m1+m2 {
			return nums1[m1]
		} else if x < m1+m2 {
			return findXNum(nums1[:m1], nums2[:m2], x)
		} else {
			return findXNum(nums1[m1+1:], nums2[m2+1:], x-m1-m2)
		}
	} else if nums1[m1] > nums2[m2] {
		if x > m1+m2 {
			return findXNum(nums1, nums2[m2+1:], x-m2)
		} else {
			return findXNum(nums1[:m1], nums2, x-m1)
		}
	} else {
		if x > m1+m2 {
			return findXNum(nums1[m1+1:], nums2, x-m1)
		} else {
			return findXNum(nums1, nums2[:m2], x-m2)
		}
	}
}
