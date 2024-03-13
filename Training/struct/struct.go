package main

import (
	"errors"
	"fmt"
)

// outside driver
type Person struct {
	name   string
	age    int
	gender string
}

func newPerson(name, gender string, age int) (*Person, error) {
	if gender != "M" && gender != "F" {
		return nil, errors.New("Gender invalid: should be M or F  ; value given was " + gender)
	}
	if age <= 0 {
		return nil, errors.New("Age invalid: should be > 0 ; value given was " + string(age))
	}
	return &Person{
		name:   name,
		age:    age,
		gender: gender,
	}, nil
}

//	func isAdult(p1 Person) bool {
//		return p1.age >= 18
//	}
func (p Person) isAdult() bool {
	return p.age >= 18
}

func (p *Person) increaseAge(increaseAgeBy int) {
	p.age += increaseAgeBy
	fmt.Println("From method", p.age) //40 50

}

func main() {
	//driver code - written by user

	person1, err := newPerson("yash", "M", 30)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	person2, err := newPerson("zxasa", "F", 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	person2.increaseAge(10)
	fmt.Println("From Main", person2.age)
	person1.increaseAge(10)
	fmt.Println("From Main", person1.age)
	person2.increaseAge(10)
	fmt.Println("From Main", person2.age)
	person1.increaseAge(10)
	fmt.Println("From Main", person1.age)
	// fmt.Println("the value of person1 ", person1.age, person1.gender, person1.name)
	// fmt.Println("the value of person2 ", person2.age, person2.name, person2.gender)
	// fmt.Println(person1.isAdult()) //true
	// fmt.Println(person2.isAdult()) //false
	// fmt.Println(person1.name)

	// person1 = person2

	// person1.name = "paul"

	// fmt.Println(person2.name)
	// fmt.Println(person1.name)

}
