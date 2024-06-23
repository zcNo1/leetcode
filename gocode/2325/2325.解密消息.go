package _325

import "strings"

const charLen = uint8(26)

func decodeMessage(key string, message string) string {
	directory := make(map[byte]byte, charLen+1)
	directory[' '] = ' '
	cur := uint8(0)

	for i := 0; i < len(key) && cur < charLen; i++ {
		if _, ok := directory[key[i]]; !ok {
			directory[key[i]] = cur + 'a'
			cur++
		}
	}

	builder := strings.Builder{}
	for i := 0; i < len(message); i++ {
		builder.WriteByte(directory[message[i]])
	}

	return builder.String()
}
