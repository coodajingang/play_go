package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	foo := flag.NewFlagSet("foo", flag.ExitOnError)
	fname := foo.String("name", "foo name", "this is foo name")
	fage := foo.Int("age", 18, "this is foo age")

	bar := flag.NewFlagSet("bar", flag.ContinueOnError)
	bname := bar.String("name", "bar name", "this is bar name")
	bage := bar.Int("age", 18, "this is bar age")

	if len(os.Args) < 2 {
		fmt.Println("usage: subflags foo bar")
		flag.Usage()
	}
	switch os.Args[1] {
	case "foo": //./subflags foo -name abc -age 22 bar  123 454
		foo.Parse(os.Args[2:])
		fmt.Println(*foo)
		fmt.Println(*fname)
		fmt.Println(*fage)
		fmt.Println(foo.Args())
	case "bar": // ./subflags  bar -age 99  123 454
		bar.Parse(os.Args[2:])
		fmt.Println(*bar)
		fmt.Println(*bname)
		fmt.Println(*bage)
		fmt.Println(bar.Args())
	default:
		fmt.Println("show default usage")
		flag.Usage()
	}
}
