package shorten

import "strings"

const alphabet = "123456789qwertyuiopasdfghjkzxcvbnmQWERTYUPASDFGHJKLZXCVBNM"

var alphabetLen = uint32(len(alphabet))

func Shorten(id uint32) string {
	var (
		digits  []uint32
		num     = id
		builder strings.Builder
	)
	for num > 0 {
		digits = append(digits, num%alphabetLen)
		num /= alphabetLen
	}

	reverse(digits)

	for _, digit := range digits {
		builder.WriteString(string(alphabet[digit]))
	}

	return builder.String()
}

func reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
