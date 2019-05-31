package main

import "fmt"

func main() {
	a := "This is Jasvinder Singh Aulakh"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := []byte(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%#U ", a[i])
	}

}
