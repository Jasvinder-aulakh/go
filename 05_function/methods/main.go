package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r reciever) identifier(parameters) (return(s))

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)

}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			"Monneypenny",
			"Don",
		},
		ltk: true,
	}
	fmt.Println(sa2)
	sa2.speak()

	sa3 := secretAgent{
		person: person{
			"Shoot",
			"Don",
		},
		ltk: true,
	}
	fmt.Println(sa3)
	sa3.speak()

	sa4 := secretAgent{
		person: person{
			"Soldier",
			"Hitlar",
		},
		ltk: true,
	}
	fmt.Println(sa4)
	sa4.speak()

}
