package main

import (
	"errors"
	"fmt"
)

type Student struct {
	name          string
	age           int
	rollNumber    int
	AdmissionYear int
	CGPA          float64
}

func (s *Student) isAdult() bool {
	return s.age >= 18
}

func (s *Student) DecrementAge(decreaseageby int) {
	s.age -= decreaseageby
	fmt.Println("From method", s.age)
}

func (s *Student) UpdateStudentInfo(newName string, newrollNumber int) error {
	if newName != " " {
		s.name = newName
	}

	if newrollNumber > 0 {
		s.rollNumber = newrollNumber
	}

	return nil
}

func (s *Student) NumberofYearToGraduate(admissionyear int) (int, error) {
	graudateyear := 2030
	yeartograduate := graudateyear - admissionyear
	// add an error check over here
	return yeartograduate, nil
}

func (s *Student) Yearofpassing() {

}

func main() {
	student1, err := newStudent("SANJEEV", 21, 1, 2021, 8.0)
	if err != nil {
		fmt.Println("There is some error", err.Error())
	}

	student2, err := newStudent("Virat", 15, 2, 2022, 9.2)
	if err != nil {
		fmt.Println("There is some error", err.Error())
	}

	// student1.age = 12
	fmt.Println("The age of student 1", student1.age)
	fmt.Println("The name of student 1", student1.name)
	fmt.Println("The roll number of student1 is", student1.rollNumber)
	fmt.Println("The roll number of student2", student2.rollNumber)
	fmt.Println("Is the student an adult", student1.isAdult())
	fmt.Println("Is the student an adult", student2.isAdult())

	student1.DecrementAge(2)
	fmt.Println("THE AGE OF STUDENT 1 FUNC CALL", student1.age)

	student1.UpdateStudentInfo("maxie", 88)
	fmt.Println("The new name of student1", student1.name)
	fmt.Println("The cgpa of student1", student1.CGPA)
	fmt.Println("The cgpa of student2", student2.CGPA)
	fmt.Println("The new rollNumber of student1", student1.rollNumber)
	fmt.Println("The number of years to graduate")
	fmt.Println(student1.NumberofYearToGraduate(2021))
}

func newStudent(name string, age, rollnumber int, admissionyear int, cgpa float64) (*Student, error) {
	if rollnumber < 1 {
		return nil, errors.New("roll number cant be less than 1")
	}

	return &Student{
		name:          name,
		age:           age,
		rollNumber:    rollnumber,
		AdmissionYear: admissionyear,
		CGPA:          cgpa,
	}, nil
}
