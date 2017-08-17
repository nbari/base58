[![Build Status](https://travis-ci.org/nbari/base58.svg?branch=master)](https://travis-ci.org/nbari/base58)
[![Coverage Status](https://coveralls.io/repos/github/nbari/base58/badge.svg?branch=master)](https://coveralls.io/github/nbari/base58?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nbari/base58)](https://goreportcard.com/report/github.com/nbari/base58)

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
