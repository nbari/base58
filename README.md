# base58 encode/decode

base58 using the Bitcoin addresses format:

    123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz

Max integer `uint64`: 18446744073709551615

Usage:

```go
package main

import (
	"fmt"

	"github.com/nbari/base58"
)

func main() {
	x := base58.Encode(1024)
	fmt.Printf("Base58 of 1024: %s\n", x)

	num := base58.Decode(x)
	fmt.Printf("Num of %q: %d\n", x, num)
}
```
