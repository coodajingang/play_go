package main

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.DateTime)) //2024-08-04 13:20:33
	fmt.Println(now.Format(time.TimeOnly)) // 13:20:33
	fmt.Println(now.Format(time.DateOnly)) // 2024-08-04
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(now.Format(time.Kitchen)) //1:20PM
	fmt.Println(now.Format(time.Layout))
	fmt.Println(now.Format(time.RFC822Z))
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))
	fmt.Println(now.Format(time.RFC1123Z))
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC3339Nano))
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format(time.Stamp))
	fmt.Println(now.Format(time.StampMilli))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02-15:04:05"))
	fmt.Println(now.Format("2006/01/02-15:04:05pm"))

	fmt.Println(time.Parse("03:04PM", "10:34PM"))

	const timestr = "2024-08-04 13:32:29"
	fmt.Println(time.ParseInLocation(time.DateTime, timestr, time.Local))
	fmt.Println(time.ParseInLocation(time.DateTime, timestr, time.UTC))
	losangle, _ := time.LoadLocation("America/Los_Angeles")
	newyork, _ := time.LoadLocation("America/New_York")
	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.ParseInLocation(time.DateTime, timestr, shanghai))
	fmt.Println(time.ParseInLocation(time.DateTime, timestr, newyork))
	fmt.Println(time.ParseInLocation(time.DateTime, timestr, losangle))

	fmt.Println(now.Location())
	fmt.Println(now.In(time.UTC))
	fmt.Println(now.In(losangle))
}
