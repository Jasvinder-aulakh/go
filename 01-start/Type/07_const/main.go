package main

import (
	"fmt"
)

const m = 34
const n = 34

const (
	a = 4
	b = 4.5
	c = "Jasvinder Singh"
)
const (
	x int     = 5
	y float32 = 45
	z string  = "Jasvinder Singh Aulakh"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	fmt.Println(m == n)
	fmt.Println(n == m)
	fmt.Printf("%T\n", m)

}
