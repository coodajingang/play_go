package main

import (
	"encoding/json"
	"fmt"
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

	b1 := []byte(`true`)
	b2 := []byte(`"abc"`)
	b3 := []byte(`100.08`)
	b4 := []byte(`{"name":"John","age":30}`)
	b5 := []byte(`["a","b","c","菊花"]`)
	b6 := []byte(`{"a":1,"b":2,"c":3,"d":4,"菊花":5}`)
	b7 := []byte(`{"page":10,"fruit":["a","b","c"]}`)
	b8 := []byte(`{"Page":10,"Fruits":["a","b"]}`)
	b9 := []byte(`{"Page":10,"Fruits":["a","b"], "em1":{"e1":1,"e2":3.8, "e3":[1,2,3],"e4":{"a":"b"}}}`)

	var (
		bb bool
		ss string
		ff float64
		m4 map[string]interface{}
		s5 []string
		m6 map[string]int
		r7 response1
		r8 response2
		m9 map[string]interface{}
	)
	json.Unmarshal(b1, &bb)
	fmt.Println(bb)
	json.Unmarshal(b2, &ss)
	fmt.Println(ss)
	json.Unmarshal(b3, &ff)
	fmt.Println(ff)
	json.Unmarshal(b4, &m4)
	fmt.Println(m4)
	json.Unmarshal(b5, &s5)
	fmt.Println(s5)
	json.Unmarshal(b6, &m6)
	fmt.Println(m6)
	json.Unmarshal(b7, &r7)
	fmt.Println(r7)
	json.Unmarshal(b8, &r8)
	fmt.Println(r8)
	json.Unmarshal(b9, &m9)
	fmt.Println(m9)

}
