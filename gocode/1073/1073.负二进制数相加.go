package l1073

import "log"

func addNegabinary(arr1 []int, arr2 []int) []int {
	cntOne := make(map[int]int)
	var ret []int
	idx1 := len(arr1) - 1
	idx2 := len(arr2) - 1
	idx := 0

	cntOne[0] = 0

	for idx1 >= 0 || idx2 >= 0 {
		log.Println("idx1: ", idx1, ", idx2: ", idx2, ", idx: ", idx)
		cnt := 0
		if idx1 >= 0 {
			cnt += arr1[idx1]
			idx1--
		}
		if idx2 >= 0 {
			cnt += arr2[idx2]
			idx2--
		}
		cnt += cntOne[idx]
		yu := cnt % 2
		shang := cnt / 2
		cntOne[idx+1] += shang
		cntOne[idx+2] += shang
		ret = append(ret, yu)
		idx++
	}

	ret = append(ret, cntOne[idx]%2)
	ret = append(ret, (cntOne[idx+1]+cntOne[idx]/2)%2)

	log.Println(ret)
	idxRet := len(ret) - 1
	for ; idxRet >= 0; idxRet-- {
		if ret[idxRet] != 0 {
			break
		}
	}

	log.Println(idxRet)

	var res []int
	for i := idxRet; i >= 0; i-- {
		res = append(res, ret[i])
	}

	if len(res) == 0 {
		res = append(res, 0)
	}

	return res
}
