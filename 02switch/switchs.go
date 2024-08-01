package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	case time.Monday:
		fmt.Println("It's monday")
	default:
		fmt.Println("I don't know whereday")
	}

	t := time.Now()
	// switch 不带表达式相当于if else
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	case t.Hour() < 17:
		fmt.Println("it's after noon")
	case t.Hour() < 19:
		fmt.Println("it's after noon")
	case t.Hour() < 24:
		fmt.Println("it's after noon")
	default:
		fmt.Println("Should not be here!")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		//case any:
		//	fmt.Println("I'm an any")
		default:
			fmt.Printf("I don't know what? %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI("hey")
	whatAmI(3)
	whatAmI(func() {})
}
