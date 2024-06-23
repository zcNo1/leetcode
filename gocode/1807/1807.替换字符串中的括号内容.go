package main

import (
	"log"
)

func main() {
	log.Println(evaluate("(name)is(age)yearsold(name)", [][]string{{"name", "bob"}, {"age", "two"}}))
}

func evaluate(s string, knowledge [][]string) string {
	replaces := make(map[string]string, len(knowledge))

	for _, val := range knowledge {
		replaces[val[0]] = val[1]
	}

	strNew := []byte(nil)
	replaceKey := []byte(nil)
	start := false
	for _, val := range []byte(s) {
		if val == '(' {
			start = true
			replaceKey = nil
			continue
		}
		if val == ')' {
			start = false
			if replaceVal, ok := replaces[string(replaceKey)]; ok {
				strNew = append(strNew, replaceVal...)
			} else {
				strNew = append(strNew, '?')
			}
			continue
		}
		if !start {
			strNew = append(strNew, val)
			continue
		}
		replaceKey = append(replaceKey, val)
	}

	return string(strNew)
}
