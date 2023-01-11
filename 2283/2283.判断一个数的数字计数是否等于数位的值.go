package main

import "log"

func main() {
	log.Println(digitCount("030"))
}

func digitCount(num string) bool {
	numByte := []byte(num)
	n := len(numByte)
	cnt := make([]int, 10)

	for _, b := range numByte {
		cnt[int(b-'0')]++
	}

	for i := 0; i < n; i++ {
		if int(numByte[i]-'0') != cnt[i] {
			return false
		}
	}

	return true
}
