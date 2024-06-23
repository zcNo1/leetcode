package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(prefixCount2([]string{"a", "attention", "practice", "attend"}, "at"))
	return
}

func prefixCount1(words []string, pref string) int {
	cnt := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			cnt++
		}
	}
	return cnt
}

func prefixCount2(words []string, pref string) int {
	cnt := 0
	pBytes := []byte(pref)
	for _, word := range words {
		wBytes := []byte(word)
		hasPre := true
		for i, pByte := range pBytes {
			if i >= len(wBytes) || pByte != wBytes[i] {
				hasPre = false
				break
			}
		}
		if hasPre {
			cnt++
		}
	}
	return cnt
}

func prefixCount3(words []string, pref string) int {
	cnt := 0
	lenPref := len(pref)
	for _, word := range words {
		if len(word) >= lenPref && word[0:lenPref] == pref {
			cnt++
		}
	}
	return cnt
}
