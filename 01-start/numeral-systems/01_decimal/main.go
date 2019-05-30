package main

import "fmt"

var y = 2
var z = 4

func main() {
	fmt.Println(y)
	// %T	a Go-syntax representation of the type of the value
	fmt.Printf("%T\n", y)
	// %b base 2
	fmt.Printf("%b\n", y)
	//%x base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", y)
	// %X	base 16, with upper-case letters for A-F
	fmt.Printf("%X\n", y)
	//alternate format: add leading 0 for octal (%#o), 0x for hex (%#x)
	fmt.Printf("%#x\n", y)
	// %v the value in a default format
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\t%b\t%x\n", y, z, y)
	// %#v	a Go-syntax representation of the value
	fmt.Printf("%#v\n", y)

}
