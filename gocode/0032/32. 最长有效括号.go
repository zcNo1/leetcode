package l0032

func longestValidParentheses(s string) int {
	var stack []int

	ret := 0

	var popUtil1Func func() bool
	var mergeStack func()
	var cnt, popNum int

	popUtil1Func = func() bool {
		cnt = 0
		for len(stack) > 0 {
			popNum = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if popNum == 1 {
				return true
			}
			cnt += popNum
		}

		return false
	}

	mergeStack = func() {
		cnt = 0
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] == 1 {
				if cnt != 0 {
					stack[i+1] = cnt
					stack = stack[:i+2]
					if cnt > ret {
						ret = cnt
					}
				}
				return
			}
			cnt += stack[i]
		}
		if cnt != 0 {
			stack[0] = cnt
			stack = stack[:1]
			if cnt > ret {
				ret = cnt
			}
		}
	}

	for _, char := range s {
		//fmt.Println(stack)
		if char == '(' {
			stack = append(stack, 1)
		} else {
			if popUtil1Func() {
				stack = append(stack, cnt+2)
				//fmt.Println(stack)
				mergeStack()
			}
		}
	}
	//fmt.Println(stack)

	return ret
}
