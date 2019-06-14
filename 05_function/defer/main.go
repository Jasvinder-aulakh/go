package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
	club()

}

func foo() {
	fmt.Println("foo")

}
func bar() {
	fmt.Println("bar")

}
func club() {
	fmt.Println("club")

}
