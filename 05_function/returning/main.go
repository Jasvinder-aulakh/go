package main

import (
	"fmt"
)

func main() {
	s := foo()
	fmt.Println(s)
	t := three()
	fmt.Println(t)

	p := four()
	fmt.Printf("%T\n", p)
	i := p()
	fmt.Println(i)

}

func foo() string {
	s1 := "Hello word"
	return s1
}

func three() string {
	t3 := "This is T3 func"
	return t3

}

func four() func() int {
	return func() int {
		return 33333333
	}

}
