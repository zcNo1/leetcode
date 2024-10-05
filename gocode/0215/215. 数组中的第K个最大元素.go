package l0215

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	nLayer := getLastSecondLayer(n)

	// get max heap
	// transfer layer
	for i := nLayer; i >= 1; i-- {
		// transfer node of this layer
		for j := 1<<(i-1) - 1; j < 1<<i-1; j++ {
			//fmt.Println(i, j)
			transferChild(nums, j)
		}
	}
	//fmt.Println(nums)

	// get k nums
	for i := 1; i < k; i++ {
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
		nums = nums[:len(nums)-1]
		transferChild(nums, 0)
		//fmt.Println(nums)
	}

	return nums[0]
}

func transferChild(nums []int, parent int) {
	n := len(nums)
	if parent >= n || parent < 0 {
		return
	}

	left := parent*2 + 1
	right := parent*2 + 2
	var maxChild int

	// no left child, mean no child, return
	if left >= n {
		return
	}

	// no right child, compare parent with left child
	if right >= n {
		maxChild = left
	} else {
		// has left and right child, select max child
		if nums[left] > nums[right] {
			maxChild = left
		} else {
			maxChild = right
		}
	}
	// if max child bigger than parent, swap them
	if nums[maxChild] > nums[parent] {
		nums[maxChild], nums[parent] = nums[parent], nums[maxChild]
		// parent might smaller than child of child, continue transfer
		transferChild(nums, maxChild)
	}
}

func getLastSecondLayer(num int) int {
	if num == 0 {
		return 0
	}
	i := 0
	for {
		if 1<<i-1 >= num {
			return i - 1
		}
		i++
	}
}

//2^n-1
//1 1
//2 3
//3 7
//
//2^(n-1)
//1 1
//2 2
//3 4

//2^(n-1)-1, 2^n-1
//1 0 1
//2 1 3
