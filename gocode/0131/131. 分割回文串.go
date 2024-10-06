package l0131

import "fmt"

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var rets [][]string

	for i := 1; i <= len(s); i++ {
		if isMatch(s[:i]) {
			subRets := partition(s[i:])
			if len(subRets) == 0 {
				rets = append(rets, append([]string{s[:i]}))
				continue
			}
			for _, subRet := range subRets {
				rets = append(rets, append([]string{s[:i]}, subRet...))
			}
		}
	}

	return rets
}

func isMatch(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	fmt.Println(s)
	return true
}
