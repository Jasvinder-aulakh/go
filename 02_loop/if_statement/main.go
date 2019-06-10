package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This is true statement")
	}
	if false {
		fmt.Println("This is false statement")

	}
	if !true {
		fmt.Println("This is not true")
	}
	if !false {
		fmt.Println("This is not false statement")
	}
	if 2 == 4/2 {
		fmt.Println("This is true")
	}
	if 2 != 3 {
		fmt.Println("This is false")
	}
	if !(2 == 4) {
		fmt.Println("This is false")
	}
	if !(3 != 3) {
		fmt.Println("True")
	}
}
