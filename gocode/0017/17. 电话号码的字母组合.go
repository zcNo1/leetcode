package l0017

var phoneNum = [][]string{
	{},                   // 0
	{},                   // 1
	{"a", "b", "c"},      //2
	{"d", "e", "f"},      //3
	{"g", "h", "i"},      //4
	{"j", "k", "l"},      //5
	{"m", "n", "o"},      //6
	{"p", "q", "r", "s"}, //7
	{"t", "u", "v"},      //8
	{"w", "x", "y", "z"}, //9
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var rets, lastRets []string

	lastRets = append(lastRets, "")

	digitsByte := []byte(digits)
	for _, b := range digitsByte {
		num := b - '0'
		for _, ret := range lastRets {
			for _, s := range phoneNum[num] {
				rets = append(rets, ret+s)
			}
		}
		lastRets = rets
		rets = []string{}
	}

	return lastRets
}
