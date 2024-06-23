package gocode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, 0)

	for _, val := range cpdomains {
		log.Printf("domain: %s", val)
		domainInfo := strings.Split(val, " ")
		cnt, err := strconv.Atoi(domainInfo[0])
		if err != nil {
			log.Printf("string %s to num failed, err: %s",
				domainInfo[1], err)
			continue
		}

		domainAll := strings.Split(domainInfo[1], ".")
		domain := domainAll[len(domainAll)-1]
		m[domain] += cnt
		for i := len(domainAll) - 2; i >= 0; i-- {
			domain = domainAll[i] + "." + domain
			m[domain] += cnt
		}
	}

	var ret []string
	for childDomain, cnt := range m {
		ret = append(ret, fmt.Sprintf("%d %s", cnt, childDomain))
	}

	return ret
}

//func main() {
//	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
//}

// @lc code=end
