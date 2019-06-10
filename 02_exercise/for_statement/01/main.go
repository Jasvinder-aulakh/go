package main

import (
	"fmt"
)

func main() {
	i := 2
	for i <= 9 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 9; j <= 13; j++ {
		fmt.Printf("%T\n\t%d\t%#x\n", j, j, j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}

}
