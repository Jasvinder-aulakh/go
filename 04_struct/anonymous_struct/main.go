package main

import (
	"fmt"
)

func main() {
	p2 := struct {
		first string
		last  string
		age   int
	}{first: "Lovy",
		last: "aulakh",
		age:  27,
	}
	fmt.Println(p2)

}
