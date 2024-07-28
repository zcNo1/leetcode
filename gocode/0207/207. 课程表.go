package l0207

func dfs(course map[int][]int, record map[int]bool, cur int) bool {
	pres, ok := course[cur]
	// 如果这门课，没有依赖，则认为正常
	if !ok || pres == nil || len(pres) <= 0 {
		return true
	}

	// 如果这门课已经被遍历过，则存在相互依赖，返回异常
	if record[cur] == true {
		return false
	}
	record[cur] = true

	for _, pre := range pres {
		if dfs(course, record, pre) == false {
			return false
		}
	}
	// 如果这门课正常，则将这门课的依赖设置为nil，这样后续再次遍历过来，可以直接认为正常返回
	course[cur] = nil

	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	course := make(map[int][]int, numCourses)
	record := make(map[int]bool, numCourses)

	for _, prerequisite := range prerequisites {
		course[prerequisite[0]] = append(course[prerequisite[0]], prerequisite[1])
		record[prerequisite[0]] = false
		record[prerequisite[1]] = false
	}

	for cur := range course {
		if dfs(course, record, cur) == false {
			return false
		}
	}

	return true
}
