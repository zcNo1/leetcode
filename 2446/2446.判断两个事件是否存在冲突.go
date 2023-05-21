package l2446

import (
	"log"
	"strconv"
	"strings"
)

func haveConflict(event1 []string, event2 []string) bool {
	startTime1 := getMinute(event1[0])
	endTime1 := getMinute(event1[1])
	startTime2 := getMinute(event2[0])
	endTime2 := getMinute(event2[1])

	log.Println(startTime1, endTime1, startTime2, endTime2)

	if startTime1 == startTime2 {
		return true
	} else if startTime1 < startTime2 {
		return endTime1 >= startTime2
	} else {
		return endTime2 >= startTime1
	}
}

func getMinute(t string) int64 {
	strs := strings.Split(t, ":")
	hours, _ := strconv.ParseInt(strs[0], 10, 64)
	minutes, _ := strconv.ParseInt(strs[1], 10, 64)
	return hours*60 + minutes
}
