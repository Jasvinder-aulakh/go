package main

import (
	"fmt"
)

func main() {
	foo()
	// anonymous func does not have name
	func() {
		fmt.Println("Anonymous func run")
	}()
	func(x int) {
		fmt.Println("the number is", x)
	}(42)
}
func foo() {
	fmt.Println("hello from foo")

}
