package main

import (
	"math/rand"
	"time"
)

func main() {
	println(rand.Int())
	println(rand.Intn(100))
	println(rand.Intn(100))

	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)
	println(r1.Intn(100))
	println(r1.Intn(100))
	println(r1.Int())
	println(r1.Float32())
	println(r1.Float32()*5 + 5)
	println(r1.Float32()*5 + 5)

}
