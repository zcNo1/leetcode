package l0020

func isValid(s string) bool {
	var record []byte

	sByte := []byte(s)

	for _, b := range sByte {
		if b == '(' || b == '[' || b == '{' {
			record = append(record, b)
		} else {
			switch b {
			case ')':
				if len(record) == 0 || record[len(record)-1] != '(' {
					return false
				}
				record = record[0 : len(record)-1]
			case ']':
				if len(record) == 0 || record[len(record)-1] != '[' {
					return false
				}
				record = record[0 : len(record)-1]
			case '}':
				if len(record) == 0 || record[len(record)-1] != '{' {
					return false
				}
				record = record[0 : len(record)-1]
			}
		}
	}
	if len(record) != 0 {
		return false
	}

	return true
}
