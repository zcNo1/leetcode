package l1156

import "log"

func maxRepOpt1(text string) int {
	if len(text) <= 1 {
		return len(text)
	}

	textByte := []byte(text)

	cnt := make(map[byte]int, 26)
	for _, r := range textByte {
		cnt[r]++
	}
	log.Println(cnt)

	ret := 1
	last := textByte[0]
	cnt_now := 1

	for i := 1; i < len(textByte); i++ {
		log.Println(i, " ", textByte[i])
		if textByte[i] == last {
			cnt_now++
			continue
		}

		if i+1 >= len(textByte)-1 {
			break
		}

		last_i := i
		if textByte[i+1] == last {
			i++
			for ; i < len(textByte); i++ {
				log.Println("inner: dx", i, " ", textByte[i])
				if textByte[i] != last {
					break
				}
				cnt_now++
			}
		}

		log.Println(cnt_now)
		if cnt_now < cnt[last] {
			cnt_now++
		}
		ret = getMax(ret, cnt_now)

		last = textByte[last_i]
		cnt_now = 1
		i = last_i
	}

	if cnt_now < cnt[last] {
		cnt_now++
	}
	ret = getMax(ret, cnt_now)

	return ret
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
