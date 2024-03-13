package structs

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) Walk() string {
	return p.firstName + p.lastName + "Like Walking"
}
