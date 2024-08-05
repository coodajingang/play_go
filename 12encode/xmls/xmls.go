package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func main() {
	//
	p := &Plant{Id: "id123", Name: "name123", Origin: []string{"1", "2", "3"}}
	out, _ := xml.Marshal(p)
	fmt.Println(string(out))
	out2, _ := xml.MarshalIndent(p, "", "  ")
	fmt.Println(string(out2))
	fmt.Println(xml.Header + string(out2))

	var p2 Plant
	if err := xml.Unmarshal(out2, &p2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	tomato := &Plant{Id: "id1", Name: "tomato"}
	tomato.Origin = []string{"1", "2", "3"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"plants>plant,omitempty"`
	}
	emp := &Plant{}
	n := &Nesting{}
	n.Plants = []*Plant{p, tomato, emp}

	out, _ = xml.MarshalIndent(n, "", "  ")
	fmt.Println(string(out))
}
