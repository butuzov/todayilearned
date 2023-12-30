# Hashing

## Example

```golang
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"lukechampine.com/blake3"
)

func main() {
	s := "foobar"

	{
		hash := sha1.New()
		hash.Write([]byte(s))
		fmt.Println("sha1  ", hex.EncodeToString(hash.Sum(nil)))
	}

	{
		hash := blake3.New(16, nil)
		hash.Write([]byte(s))
		fmt.Println("blake3", hex.EncodeToString(hash.Sum(nil)))
	}

	{
		has := md5.New()
		has.Write([]byte(s))
		fmt.Println("md5   ", hex.EncodeToString(has.Sum(nil)))
	}
}
```

## Algorithmes

### Blacke3

- https://golangexample.com/pure-go-implementation-of-blake3-with-avx2-and-sse4-1-acceleration/
