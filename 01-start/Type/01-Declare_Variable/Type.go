package main

import "fmt"

var a = 10
var b int
var c = "This is Jasvinder Testing book"

func main() {
	// declare x variable
	x := 34
	fmt.Println(x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	b = x + a
	fmt.Printf("%T\n", b)

}
