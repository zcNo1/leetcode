package l0022

func generateParenthesis(n int) []string {
	var ret []string
	dfx(n, 0, "", &ret)
	return ret
}

func dfx(n int, leftCount int, cur string, ret *[]string) {
	if n == 0 {
		for leftCount > 0 {
			cur += ")"
			leftCount--
		}
		*ret = append(*ret, cur)
		return
	}

	dfx(n-1, leftCount+1, cur+"(", ret)

	if leftCount > 0 {
		dfx(n, leftCount-1, cur+")", ret)
	}
}
