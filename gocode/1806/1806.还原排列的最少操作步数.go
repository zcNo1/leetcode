package main

import "log"

func main() {
	var ret []int
	for i := 2; i <= 1000; i += 2 {
		ret = append(ret, reinitializePermutation(i))
	}

	log.Println(ret)
}

func reinitializePermutation(n int) int {
	origin := make([]int, n)
	first := make([]int, n)
	second := make([]int, n)
	for i := 0; i < n; i++ {
		origin[i] = i
		first[i] = i
	}
	cnt := 1
	for {
		//log.Printf("num[%d]", cnt)
		if cnt%2 == 1 {
			change(second, first)
			//log.Printf("arr: %v", second)
			if checkSame(second, origin) {
				break
			}
		} else {
			change(first, second)
			//log.Printf("arr: %v", first)
			if checkSame(first, origin) {
				break
			}
		}

		cnt++
	}

	return cnt
}

func change(arr1, arr2 []int) {
	halfN := len(arr1) / 2
	for i := 0; i < len(arr1); i++ {
		if i%2 == 0 {
			arr1[i] = arr2[i/2]
		} else {
			arr1[i] = arr2[halfN+(i-1)/2]
		}
	}
}

func checkSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
