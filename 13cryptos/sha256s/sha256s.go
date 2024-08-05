package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("hello world"))
	sum := hash.Sum(nil)
	fmt.Println(sum)
	fmt.Printf("%x\n", sum)

	h512 := sha512.New()
	h512.Write([]byte("hello world"))
	sum = h512.Sum(nil)
	fmt.Println(sum)
	fmt.Printf("%x\n", sum)
}
