package main

import "fmt"

//capital names means public field it can be accessed anywhere in the package
// if it will be small name like FirstName - firstname it will become private field

// encapsulation - it says that a struct field is private it can be accessed
type Person struct {
	firstName string
	lastName  string
	Age       int
	FullName  func() string
}

func (p Person) Walk() string {
	return p.firstName + p.lastName + "likes Walking"
}

func main() {
	p := Person{
		firstName: "Sanjeev",
		lastName:  "Yadav",
		Age:       23,
	}
	s := p.Walk()
	p1 := Person{
		firstName: "virat",
		lastName:  "kohli",
		Age:       23,
	}

	fmt.Println(p, p1, s)
}
