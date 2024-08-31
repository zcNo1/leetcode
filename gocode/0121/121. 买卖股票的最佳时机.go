package l0121

func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}
	ret := 0

	pre := prices[1] - prices[0]
	if pre > ret {
		ret = pre
	}
	for i := 2; i < len(prices); i++ {
		if pre < 0 {
			pre = prices[i] - prices[i-1]
		} else {
			pre += prices[i] - prices[i-1]
		}
		if pre > ret {
			ret = pre
		}
	}

	return ret
}
