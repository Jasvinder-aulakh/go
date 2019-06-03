package main

import (
	"fmt"
)

const (
	j = 10
	k = 10 + iota
	l = 10 + iota
	m = 10 + (iota * 2)
	n = 10 + (iota * 2)
	o = 10 + (iota * 2)
	p = 10 + iota
	q = 10 + iota
)

func main() {
	a := 1
	b := 2
	c := 3
	d := 4
	e := 5
	f := 6
	g := 7
	h := 8
	i := 9

	fmt.Printf("%d\\t%b\\t%#x\n", a, a, a)
	fmt.Printf("%d\\t%b\\t%#x\n", b, b, b)
	fmt.Printf("%d\\t%b\\t%#x\n", c, c, c)
	fmt.Printf("%d\\t%b\\t%#x\n", d, d, d)
	fmt.Printf("%d\\t%b\\t%#x\n", e, e, e)
	fmt.Printf("%d\\t%b\\t%#x\n", f, f, f)
	fmt.Printf("%d\\t%b\\t%#x\n", g, g, g)
	fmt.Printf("%d\\t%b\\t%#x\n", h, h, h)
	fmt.Printf("%d\\t%b\\t%#x\n", i, i, i)

	fmt.Printf("%d\\t%b\\t%#x\n", j, j, j)
	fmt.Printf("%d\\t%b\\t%#x\n", k, k, k)
	fmt.Printf("%d\\t%b\\t%#x\n", l, l, l)
	fmt.Printf("%d\\t%b\\t%#x\n", m, m, m)
	fmt.Printf("%d\\t%b\\t%#x\n", n, n, n)
	fmt.Printf("%d\\t%b\\t%#x\n", o, o, o)
	fmt.Printf("%d\\t%b\\t%#x\n", p, p, p)

}
