package l0076

func minWindow(s string, t string) string {
	cnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		cnt[t[i]]++
	}

	var left, right int

	// 先找到第一个值
	for i := 0; i < len(s); i++ {
		if checkRightByte(cnt, s[i]) {
			left = i
			cnt[s[i]]--
			break
		}
	}

	// 目标长度是1，特殊处理
	if len(t) == 1 {
		if checkOk(cnt) {
			return t
		} else {
			return ""
		}
	}

	// 找到第一个满足的情况
	for i := left + 1; i < len(s); i++ {
		if checkRightByte(cnt, s[i]) {
			cnt[s[i]]--
			if checkOk(cnt) {
				right = i + 1
				break
			}
		}
	}
	if !checkOk(cnt) {
		return ""
	}

	ret := s[left:right]
	// 如果与左侧第一个值相等，可以弹出左侧，囊括右侧
	for cnt[s[left]] < 0 {
		cnt[s[left]]++
		left++
		for !checkRightByte(cnt, s[left]) {
			left++
		}
	}

	if len(ret) > right-left {
		ret = s[left:right]
	}
	//fmt.Println(ret)

	for left < right && right < len(s) {
		if checkRightByte(cnt, s[right]) {
			cnt[s[right]]--
		}
		right++

		// 如果与左侧第一个值相等，可以弹出左侧，囊括右侧
		for cnt[s[left]] < 0 {
			cnt[s[left]]++
			left++
			for !checkRightByte(cnt, s[left]) {
				left++
			}
		}

		if len(ret) > right-left {
			ret = s[left:right]
		}
	}

	return ret
}

func checkRightByte(cnt map[byte]int, a byte) bool {
	_, ok := cnt[a]
	return ok
}

func checkOk(cnt map[byte]int) bool {
	for _, i := range cnt {
		if i > 0 {
			return false
		}
	}

	return true
}
