package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	1. 当前时间
	2. 创建时间
	3. 时间比较
	4. 时间属性
	5. 时间加减
	*/
	now := time.Now()
	fmt.Println(now)

	date := time.Date(2024, 6, 28, 12, 30, 59, 432300, time.UTC)
	fmt.Println(date)

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	fmt.Println(date.Hour())
	fmt.Println(date.Minute())
	fmt.Println(date.Second())
	fmt.Println(date.Nanosecond())
	fmt.Println(date.Location())
	fmt.Println(date.Weekday())
	fmt.Println(date.YearDay())

	fmt.Println(date.Before(now))
	fmt.Println(date.After(now))
	fmt.Println(date.Equal(now))

	sub := now.Sub(date)
	fmt.Println(sub)
	fmt.Println(sub.Hours())
	fmt.Println(sub.Minutes())
	fmt.Println(sub.Seconds())
	fmt.Println(sub.Milliseconds())
	fmt.Println(sub.Nanoseconds())

	add := now.Add(8 * time.Hour)
	fmt.Println(add)
}
