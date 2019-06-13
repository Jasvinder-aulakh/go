package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type rawAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := rawAgent{
		person: person{
			first: "james",
			age:   32,
			last:  "bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

}
