package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	surname string
	age     int
}

func main() {
	p1 := person{
		first:   "Jasvinder",
		surname: "Aulakh",
		last:    "Singh",
		age:     22,
	}

	p2 := person{
		first: "Happy",
		last:  "mand",
		age:   20,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.surname, p1.age)
	fmt.Println(p2.first, p2.last, p2.surname, p2.age)

}
