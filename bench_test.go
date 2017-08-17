package base58

import (
	"testing"
)

func BenchmarkBase58EncodeDecode(b *testing.B) {
	for i, val := uint64(0), uint64(1<<16); i <= val; i++ {
		Decode(Encode(i))
	}
}
