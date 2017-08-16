package base58

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var i uint64
	for i = 0; i <= (1<<64 - 1); i++ {
		b58 := Encode(i)
		num := Decode(b58)
		if num != i {
			t.Fatalf("Expecting %d for %v", i, b58)
		}
		//fmt.Printf("b58 = %v num = %d i = %d\n", b58, num, i)
	}
}