package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main function")
	f := func() {
		fmt.Println("My first function expression")
	}
	f()
	y := func(u int) {
		fmt.Println("The user has started since", u)
	}
	y(55)
}
