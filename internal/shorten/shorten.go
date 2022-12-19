package shorten

import (
	"strings"

	"github.com/dbashirov/link-shrinker/internal/utils"
)

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

	utils.Reverse(digits)

	for _, digit := range digits {
		builder.WriteString(string(alphabet[digit]))
	}

	return builder.String()
}
