package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Jass")
	pop(911)
}

// func ()  {}
func foo() {
	fmt.Println("hello this is foo")
}
func bar(s string) {
	// bar(s string)
	fmt.Println("Hello,", s)
}

// everything in go is pass by values
func pop(i int) {
	fmt.Println("hello,", i)

}
