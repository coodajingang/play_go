package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	println(s)
	fmt.Printf("%T\n", s)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
