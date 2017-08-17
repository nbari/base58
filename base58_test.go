package base58

import (
	"testing"
)

var testTable = []struct {
	in  uint64
	out string
}{
	{0, "1"},
	{1, "2"},
	{32, "Z"},
	{58, "21"},
	{59, "22"},
	{64, "27"},
	{1024, "Jf"},
	{4096, "2Dd"},
	{8192, "3SF"},
	{16384, "5sV"},
	{65536, "LUw"},
	{4294967296, "7YXq9H"},
	{18446744073709551615, "jpXCZedGfVQ"},
}

// how to make this fast
func TestEncodeDecode(t *testing.T) {
	var i uint64
	for i = 0; i <= (1 << 16); i++ {
		b58 := Encode(i)
		num := Decode(b58)
		if num != i {
			t.Fatalf("Expecting %d for %s", i, b58)
		}
	}
}

func TestEncode(t *testing.T) {
	for _, tt := range testTable {
		b58 := Encode(tt.in)
		if b58 != tt.out {
			t.Fatalf("Expecting %v for %s", b58, tt.out)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tt := range testTable {
		num := Decode(tt.out)
		if num != tt.in {
			t.Fatalf("Expecting %d for %d", tt.in, num)
		}
	}
}
