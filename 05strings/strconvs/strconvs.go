package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ParseFloat 第二个参数表示解析的位数 32/64 ， 第一个参数是符合[floating-point literals]的浮点数
	f1, _ := strconv.ParseFloat("1.23456789", 64)
	fmt.Println(f1)
	f2, _ := strconv.ParseFloat("1.23456789", 32)
	fmt.Println(f2)
	f3, err := strconv.ParseFloat("0x123456789", 32)
	fmt.Println(err) //parsing "0x123456789": invalid syntax
	fmt.Println(f3)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	i2, _ := strconv.ParseInt("0x123ba", 0, 64)
	fmt.Println(i2)

	u1, _ := strconv.ParseUint("0x123ba", 0, 64)
	fmt.Println(u1)

	fmt.Println(strconv.Quote("abc"))
	fmt.Println(strconv.QuotedPrefix(`"abc"''`))
	fmt.Println(strconv.QuoteToGraphic(`abc"''`))
	fmt.Println(strconv.QuoteRune('A'))
	fmt.Println(strconv.QuoteRune('菊'))
	fmt.Println(strconv.QuoteRuneToASCII('菊'))

	fmt.Println(strconv.FormatInt(0xabc, 10))
	fmt.Println(strconv.FormatInt(0xabc, 16))
	fmt.Println(strconv.FormatInt(0xabc, 2))
	fmt.Println(strconv.FormatInt(0xabc, 8))
	fmt.Println(strconv.FormatInt(0xabc, 20))
}
