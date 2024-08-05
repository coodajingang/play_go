package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	estr := base64.StdEncoding.EncodeToString([]byte("hello world12aweabc123!?$*&()'-=@~"))
	fmt.Println(estr)

	decodeString, err := base64.StdEncoding.DecodeString(estr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodeString))

	urlstr := base64.URLEncoding.EncodeToString([]byte("hello world12awehttp://132?asdfasd=asdf=asdf=asdabc123!?$*&()'-=@~"))
	fmt.Println(urlstr)

	bytes, err := base64.URLEncoding.DecodeString(urlstr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	toString := base64.RawStdEncoding.EncodeToString([]byte("hello world12aweabc123!?$*&()'-=@~"))
	fmt.Println(toString)

}
