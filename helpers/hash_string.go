package helpers

import (
	"math/big"
	"strings"
)

const base62Chars = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789"

func Encrypt(input string) string {
	base := big.NewInt(62)
	zero := big.NewInt(0)
	result := make([]byte, 0)

	n := new(big.Int)
	n.SetString(input, 10)

	for n.Cmp(zero) > 0 {
		_, rem := new(big.Int), new(big.Int)
		n.DivMod(n, base, rem)
		result = append(result, base62Chars[rem.Int64()])
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func Decrypt(input string) string {
	base := big.NewInt(62)
	result := big.NewInt(0)

	for _, char := range input {
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(strings.IndexByte(base62Chars, byte(char)))))
	}

	return result.String()
}
