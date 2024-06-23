package l1335

import "math"

var res int

func minDifficulty(jobDifficulty []int, d int) int {
	res = math.MaxInt
	if len(jobDifficulty) < d {
		return -1
	}
	if d == 1 {
		var ret int
		for _, i2 := range jobDifficulty {
			if ret < i2 {
				ret = i2
			}
		}
		return ret
	}

	if len(jobDifficulty) == d {
		var ret int
		for _, i2 := range jobDifficulty {
			ret += i2
		}
		return ret
	}

	if jobDifficulty[0] <= jobDifficulty[len(jobDifficulty)-1] {
		return jobDifficulty[0] + minDifficulty(jobDifficulty[1:], d-1)
	} else {
		return jobDifficulty[len(jobDifficulty)-1] + minDifficulty(jobDifficulty[:len(jobDifficulty)-1], d-1)
	}
}

func dsf(jobDifficulty []int) {

}
