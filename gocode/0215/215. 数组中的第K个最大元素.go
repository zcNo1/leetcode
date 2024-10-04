package l0215

func findKthLargest(nums []int, k int) int {
	n := getLastSecondLayer(len(nums))

	for i := n; i >= 1; i-- {
		for j := 2 ^ i; j < 2^i-1; j++ {

		}
	}
}

func getLastSecondLayer(num int) int {
	i := 1
	for {
		if 2^i-1 >= num {
			return i - 1
		} else {
			i++
		}
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

//2^(n-1), 2^n-1
//1  
//2
