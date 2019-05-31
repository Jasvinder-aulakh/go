package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)
const (
	x = iota + 1
	y = iota + 1
	z = iota + 1
)
const (
	d = iota
	e
	f
	g = iota
	h = iota
	i
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Printf("%T\n%T\n%T\n", a, b, c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n%T\n%T\n", x, y, z)
}
