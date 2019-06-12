package main

import (
	"fmt"
)

func main() {
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
func abc() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}

}
