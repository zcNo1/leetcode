package _233

import "sort"

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	var ret []string
	ret = append(ret, folder[0])
	for i := 1; i < len(folder); i++ {
		cur := ret[len(ret)-1] + "/"
		if len(cur) > len(folder[i]) || cur != folder[i][:len(cur)] {
			ret = append(ret, folder[i])
		}
	}

	return ret
}
