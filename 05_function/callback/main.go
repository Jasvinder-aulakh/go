package mian

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

}

func sum(x1 ...int) int {
	fmt.Printf("%T\n", x1)
	total := 0
	for _, v := range x1 {
		total += v

	}
	return total

}
