package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p("Contains:", strings.Contains("test", "es"))
	p("ContainsAny:", strings.ContainsAny("test", "es"))
	p("ContainsRune:", strings.ContainsRune("这是一菊花", rune('菊')))
	p("ContainsFunc", strings.ContainsFunc("test sencond", func(r rune) bool {
		return r == rune('a')
	}))

	p("Count", strings.Count("test sencond sentence", "sen"))
	p("HasPrefix", strings.HasPrefix("pre123.log", "pre"))
	p("HasSuffix", strings.HasSuffix("pre123.log", ".log"))

	p("Index:", strings.Index("test sestes", "es"))
	p("IndexAny:", strings.IndexAny("test sestes", "es"))
	p("IndexRune:", strings.IndexRune("这是一菊花,请来一杯菊花", rune('菊')))
	p("IndexRune:", strings.Index("这是一菊花,请来一杯菊花", "菊"))
	p("IndexRune:", strings.IndexRune("这是一菊花,请来一杯菊花", rune('a')))
	p("IndexFunc", strings.IndexFunc("test 这是一菊花，请来一杯菊花", func(r rune) bool {
		return r == rune('菊')
	}))

	p("Join", strings.Join([]string{"a", "b", "杯菊", "花"}, "-"))
	p("Split", strings.Split("a-b-杯菊-花", "-"))

	p("ToUpper", strings.ToUpper("a-b-杯菊-花"))
	p("ToLower", strings.ToLower("A-B-杯菊-花"))
	p("ToTitle", strings.ToTitle("a-b-杯菊-花"))
	p("ToTitle", strings.ToTitle("test sestes"))

	p("Repeat", strings.Repeat("a@!2", 3))

	p("ReplaceAll", strings.ReplaceAll("test sestes", " ", "_"))
	p("Replace", strings.Replace("test sestes", "es", "_", -1))
	p("Replace", strings.Replace("test sestes", "es", "_", 1))

	p("Map", strings.Map(func(r rune) rune {
		switch r {
		case rune('A'):
			return '9'
		case rune('菊'):
			return '蓝'
		default:
			return r
		}
	}, "A-B-杯菊-花"))

	af, cut := strings.CutPrefix("pre123.log", "pre")
	p("CutPrefix", af, cut)
	af, cut = strings.CutPrefix("pre123.log", "preo")
	p("CutPrefix", af, cut)
	af, cut = strings.CutSuffix("pre123.log", ".log")
	p("CutSuffix", af, cut)
	af, cut = strings.CutSuffix("pre123.log", ".log2")
	p("CutSuffix", af, cut)

	be, af, cut := strings.Cut("test_pre123.log", "e")
	p("Cut", be, af, cut)
	be, af, cut = strings.Cut("test_pre123.log", "100")
	p("Cut", be, af, cut)

	p("EqualFold", strings.EqualFold("golang", "GOLang"))
	p("EqualFold", strings.EqualFold("golang", ""))
	var s string
	p("EqualFold", strings.EqualFold("golang", s))

	p("Fields", strings.Fields("test essda a233"))
	p("Fields", strings.Fields(""))
	p("Fields", strings.Fields(s))
}
