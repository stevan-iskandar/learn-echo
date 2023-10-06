package helpers

import (
	"math/big"
)

func HashString(value string) string {
	// Convert the hexadecimal string to a decimal integer
	decimalInt := new(big.Int)
	decimalInt.SetString(value, 16)

	// Encode the decimal integer as Basbig
	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base62 := ""
	base := big.NewInt(62)
	zero := big.NewInt(0)

	for decimalInt.Cmp(zero) > 0 {
		mod := new(big.Int)
		decimalInt.DivMod(decimalInt, base, mod)
		base62 = string(base62Chars[mod.Int64()]) + base62
	}

	return base62
}
