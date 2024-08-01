package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruit"`
}
type response2 struct {
	Page   int
	Fruits []string
}

func main() {
	// 基本类型
	i, _ := json.Marshal(10)
	fmt.Println(string(i))
	b, _ := json.Marshal(true)
	fmt.Println(string(b))
	s, _ := json.Marshal("abc")
	fmt.Println(string(s))
	f, _ := json.Marshal(10.88)
	fmt.Println(string(f))
	s2, _ := json.Marshal("abc这是以菊花残")
	fmt.Println(string(s2))

	slic := []string{"a", "b", "c", "菊花"}
	sl, _ := json.Marshal(slic)
	fmt.Println(string(sl))

	arr := [3]int{1, 2, 3}
	ar, _ := json.Marshal(arr)
	fmt.Println(string(ar))
	mm := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "菊花": 5}
	mmm, _ := json.Marshal(mm)
	fmt.Println(string(mmm))

	// struct encode
	r1 := &response1{10, []string{"a", "b", "c"}}
	r2 := &response2{10, []string{"a", "b"}}

	re1, _ := json.Marshal(r1)
	re2, _ := json.Marshal(r2)
	fmt.Println(string(re1))
	fmt.Println(string(re2))

	// custom encoder
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(r1)
	enc.Encode(r2)
}
