package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hash = flag.String("hash", "sha256", "ハッシュ関数を指定します。sha256かsha512が使えます.")
var msg = flag.String("msg", "abc", "ハッシュする文字列を指定します。")

func main() {
	flag.Parse()

	if *hash == "sha256" {
		var hashed_msg [32]byte
		hashed_msg = sha256.Sum256([]byte(*msg))
		fmt.Printf("%x\n", hashed_msg)
	} else if *hash == "sha512" {
		var hashed_msg [64]byte
		hashed_msg = sha512.Sum512([]byte(*msg))
		fmt.Printf("%x\n", hashed_msg)
	} else {
		var hashed_msg [32]byte = sha256.Sum256([]byte("x"))
		fmt.Printf("%x\n", hashed_msg)
	}
}
