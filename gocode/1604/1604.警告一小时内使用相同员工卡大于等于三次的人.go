package _604

import (
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	record := make(map[string][]int)
	for i, s := range keyName {
		record[s] = append(record[s], getTime(keyTime[i]))
	}

	var ret []string
	var left, right int
	for s, ints := range record {
		if len(ints) < 3 {
			continue
		}
		sort.Ints(ints)
		left = 0
		right = 2
		for right < len(ints) {
			if ints[right]-ints[left] <= 60 {
				ret = append(ret, s)
				break
			}
			left++
			right++
		}
	}
	sort.Strings(ret)
	return ret
}

func getTime(hours string) int {
	minutes := 0
	ret := strings.Split(hours, ":")
	if len(ret) != 2 {
		return -1
	}
	num, err := strconv.Atoi(ret[0])
	if err != nil {
		return -1
	}
	minutes += 60 * num

	num, err = strconv.Atoi(ret[1])
	if err != nil {
		return -1
	}
	minutes += num

	return minutes
}
