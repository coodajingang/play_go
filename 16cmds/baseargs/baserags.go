package main

import (
	"fmt"
	"os"
)

func main() {
	argsAll := os.Args

	args := os.Args[1:]

	aa := os.Args[3]

	fmt.Println(argsAll)
	fmt.Println(args)
	fmt.Println(aa)
}
