package helpers

import (
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(number uint64) string {
	var builder strings.Builder

	builder.Grow(10)

	for ; number > 0; number = number / uint64(len(alphabet)) {
		builder.WriteByte(alphabet[number%uint64(len(alphabet))])
	}

	return builder.String()
}
