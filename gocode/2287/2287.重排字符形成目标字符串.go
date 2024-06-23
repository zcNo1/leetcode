package main

import "log"

func main() {
	log.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}

func rearrangeCharacters(s string, target string) int {
	var cntS, cntTar [26]int
	for _, i2 := range s {
		cntS[i2-'a']++
	}
	for _, i2 := range target {
		cntTar[i2-'a']++
	}

	ret := len(s)
	for i, cnt := range cntTar {
		if cnt == 0 {
			continue
		}
		if cntS[i] < cntTar[i] {
			return 0
		}
		if cntS[i]/cntTar[i] < ret {
			ret = cntS[i] / cntTar[i]
		}
	}

	return ret
}
