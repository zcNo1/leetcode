package _234

func balancedString(s string) int {
	if len(s)%4 != 0 {
		return -1
	}
	cnt := len(s) / 4
	sByte := []byte(s)
	cnts := make(map[byte]int)
	for _, i2 := range sByte {
		cnts[i2]++
	}

	for b, i := range cnts {
		cnts[b] = i - cnt
		if cnts[b] <= 0 {
			delete(cnts, b)
		}
	}
	if len(cnts) == 0 {
		return 0
	}

	var left, right, max int
	left = -1
	for right < len(s) && left < right {
		if sByte[right]
	}

	return max
}
