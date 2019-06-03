package main

import (
	"fmt"
)

const (
	x = 4
	y = x >> 1
	z = y << 1
)

const (
	k = 1024
	m = 1024 * k
	g = 1024 * m
)

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t%b\n", x, x)
	fmt.Printf("%d\t\t\t%b\n", x, x)
	fmt.Printf("%d\t\t\t%b\n", y, y)
	fmt.Printf("%d\t\t%b\n", z, z)

	fmt.Printf("%d\t\t\t%b\n", k, k)
	fmt.Printf("%d\t\t\t%b\n", m, m)
	fmt.Printf("%d\t\t%b\n", g, g)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
