package main

import "fmt"

var a string

var b int

var c = 50

func main() {
	// Declare a variable for a certain type
	// and then assign a value of that type to the variable
	fmt.Println("Printing the value of y", a, "ending")
	fmt.Printf("%T\n", a)
	a = "Jasvinder Singh Aulakh"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = 54
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	d := 50
	fmt.Printf("%T\n", c)
	fmt.Println(d)

}
