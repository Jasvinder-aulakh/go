package main

import (
	"fmt"
)

func main() {
	foo()
	jass(1, 2, 3, 4, 5, 6, 7, 8, 9)
	singh("Hello", "This", "is", "testing", "String")

}

func foo() {
	fmt.Println("Testing")

}

func jass(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
func singh(y ...string) {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
