package base58

import (
	"math"
)

var (
	// Bitcoint base58
	base58 = []byte{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W',
		'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	// Base58 symbol chart
	index = map[byte]int{
		'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
		'A': 9, 'B': 10, 'C': 11, 'D': 12, 'E': 13, 'F': 14, 'G': 15, 'H': 16,
		'J': 17, 'K': 18, 'L': 19, 'M': 20, 'N': 21, 'P': 22, 'Q': 23, 'R': 24,
		'S': 25, 'T': 26, 'U': 27, 'V': 28, 'W': 29, 'X': 30, 'Y': 31, 'Z': 32,
		'a': 33, 'b': 34, 'c': 35, 'd': 36, 'e': 37, 'f': 38, 'g': 39, 'h': 40,
		'i': 41, 'j': 42, 'k': 43, 'm': 44, 'n': 45, 'o': 46, 'p': 47, 'q': 48,
		'r': 49, 's': 50, 't': 51, 'u': 52, 'v': 53, 'w': 54, 'x': 55, 'y': 56,
		'z': 57}
)

// Encode an int into base58 bitcoint format
func Encode(value uint64) string {
	if value < 58 {
		return string(base58[value])
	}
	var (
		res [16]byte
		i   int
	)
	for i = len(res) - 1; value != 0; i-- {
		res[i] = base58[value%58]
		value /= 58
	}
	return string(res[i+1:])
}

// Decode decodes a base58-encoded string
func Decode(s string) uint64 {
	res := uint64(0)
	l := len(s) - 1
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		res += uint64(byteOffset) * uint64(math.Pow(58, float64(idx)))
	}
	return res
}
