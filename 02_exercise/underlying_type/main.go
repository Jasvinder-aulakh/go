package main

import (
	"fmt"
)

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
