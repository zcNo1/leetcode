package l0394

import "fmt"

func decodeString(s string) string {
	var ret []byte
	sBytes := []byte(s)

	num := 0
	var nums []int
	var cur []byte
	var repeats [][]byte
	for _, sByte := range sBytes {
		if sByte >= '0' && sByte <= '9' {
			num = num*10 + int(sByte-'0')
			fmt.Println(num)
		} else if sByte == '[' {
			// 记录nums
			nums = append(nums, num)
			num = 0
			// nums为空，证明是第一次，将之前的tmp记录
			if len(nums) == 1 {
				ret = append(ret, cur...)
			} else {
				repeats = append(repeats, cur)
			}
			cur = []byte{}
		} else if sByte == ']' {
			var tmp []byte
			for i := 0; i < nums[len(nums)-1]; i++ {
				tmp = append(tmp, cur...)
			}
			nums = nums[:len(nums)-1]
			if len(nums) == 0 {
				ret = append(ret, tmp...)
				cur = []byte{}
			} else {
				cur = repeats[len(repeats)-1]
				cur = append(cur, tmp...)
				repeats = repeats[:len(repeats)-1]
			}
		} else {
			cur = append(cur, sByte)
		}
	}

	ret = append(ret, cur...)

	return string(ret)
}
