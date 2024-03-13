package main

import (
	"errors"
	"fmt"
	"time"
)

type Student struct {
	firstName              string
	lastName               string
	fullName               string
	DOB                    int
	age                    int
	semesterCgpa           []float64
	finalCgpa              float64
	semesterGrades         []string
	finalGrade             string
	yearofEnrollement      int
	Yearofpassing          int
	NumberofYearToGraduate int
	AdmissionYear          int
	CGPA                   float64
}

func newStudent(firstname, lastname string, yearofenrollement, yearofpassing, dob int, semestercgpa []float64) (*Student, error) {

	if yearofenrollement > yearofpassing {
		return nil, errors.New("Year of enrollment cant be greater than year of passing")
	}

	for _, cgpa := range semestercgpa {
		if cgpa < 0 {
			return nil, errors.New("Cgpa cant be less than 0")
		}
	}

	if firstname == "" {
		return nil, errors.New("Firstname missing")
	}
	if lastname == "" {
		return nil, errors.New("last Name missing")
	}
	fullName := firstname + lastname
	// 	return s.fullName

	// AGE OF STUDENT
	currentYear := time.Now().Year()
	studentage := currentYear - dob

	// NUMBER OF YEARS TO GRADUATE
	yeartograduate := yearofpassing - yearofenrollement

	// FINAL CGPA OF A STUDENT
	for _, cgpa := range semestercgpa {
		if cgpa < 0.0 {
			return nil, errors.New("The cgpa can't be less than 0")
		}
	}
	total := 0.0
	for _, cgpaValue := range semestercgpa {
		total += cgpaValue
	}
	averagefinalcgpa := total / float64(len(semestercgpa))

	semesterGrades := []string{}
	for _, cgpa := range semestercgpa {
		if cgpa >= 9 {
			semesterGrades = append(semesterGrades, "A")
		} else if cgpa >= 8 {
			semesterGrades = append(semesterGrades, "B")
		} else if cgpa >= 7 {
			semesterGrades = append(semesterGrades, "C")
		} else if cgpa >= 6 {
			semesterGrades = append(semesterGrades, "D")
		} else if cgpa >= 5 {
			semesterGrades = append(semesterGrades, "E")
		} else {
			semesterGrades = append(semesterGrades, "F")
		}
	}

	// final grade of student
	var finalgrade string
	if averagefinalcgpa >= 9 {
		finalgrade = "A"
	} else if averagefinalcgpa >= 8 {
		finalgrade = "B"
	} else if averagefinalcgpa >= 7 {
		finalgrade = "C"
	} else if averagefinalcgpa >= 6 {
		finalgrade = "D"
	} else if averagefinalcgpa >= 5 {
		finalgrade = "E"
	} else {
		finalgrade = "F"
	}

	return &Student{
		firstName:              firstname,
		lastName:               lastname,
		fullName:               fullName,
		DOB:                    dob,
		age:                    studentage,
		semesterCgpa:           semestercgpa,
		finalCgpa:              averagefinalcgpa,
		yearofEnrollement:      yearofenrollement,
		Yearofpassing:          yearofpassing,
		NumberofYearToGraduate: yeartograduate,
		semesterGrades:         semesterGrades,
		finalGrade:             finalgrade,
	}, nil
}

// func (s *Student) fullNameofStudent(firstName, lastName string) string {
// 	s.fullName = s.firstName + s.lastName
// 	return s.fullName
// }

// func (s *Student) getAge(dob int) {
// 	currentYear := time.Now().Year()
// 	studentage := currentYear - dob
// 	fmt.Println("The age od the student is", studentage)
// }

// func (s *Student) NumberofYearTobeGraduated(admissionyear int) (int, error) {
// 	graudateyear := 2030
// 	yeartograduate := graudateyear - admissionyear
// 	return yeartograduate, nil
// }

// func (s *Student) Yearofgraduaterequired(yearofenrollement, yearofpassing int) (int, error) {
// 	graudateyear := yearofpassing - yearofenrollement
// 	fmt.Println("The passing year is", graudateyear)
// 	return graudateyear, nil
// }

// func (s *Student) FinalCgpaofStudent(semesterCgpa []float64) (float64, error) {
// 	for _, cgpa := range semesterCgpa {
// 		if cgpa < 0.0 {
// 			return 0.0, errors.New("The cgpa can't be less than 0")
// 		}
// 	}

// 	total := 0.0
// 	for _, cgpaValue := range semesterCgpa {
// 		total += cgpaValue
// 	}
// 	average := total / float64(len(semesterCgpa))

// 	return average, nil
// }

// func (s *Student) SemGrades(cgpa float64) {

// 	var grade string

// 	if cgpa >= 9 {
// 		grade = "A"
// 	} else if cgpa >= 8 {
// 		grade = "B"
// 	} else if cgpa >= 7 {
// 		grade = "C"
// 	} else if cgpa >= 6 {
// 		grade = "D"
// 	} else if cgpa >= 5 {
// 		grade = "E"
// 	} else {
// 		grade = "F"
// 	}

// 	fmt.Println(grade)
// }

// func (s *Student) finalGradeofStudent(finalcgpa string) {
// 	var finalgrade string

// 	if s.finalCgpa >= 9 {
// 		finalgrade = "A"
// 	} else if s.finalCgpa >= 8 {
// 		finalgrade = "B"
// 	} else if s.finalCgpa >= 7 {
// 		finalgrade = "C"
// 	} else if s.finalCgpa >= 6 {
// 		finalgrade = "D"
// 	} else if s.finalCgpa >= 5 {
// 		finalgrade = "E"
// 	} else {
// 		finalgrade = "F"
// 	}

// 	fmt.Println(finalgrade)
// }

func main() {
	semcgpa := []float64{1, 2, 3, 9}
	yash, err := newStudent("yash", "shah", 2013, 2023, 1993, semcgpa)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Println("the grades of yash", yash.semesterGrades)
	fmt.Println("THE FINAL grades of yash", yash.finalGrade)
	fmt.Println("The current cgpa", yash.semesterCgpa)
	fmt.Println("THE FINAL CGPA OF YASH", yash.finalCgpa)

	new_cgpa := []float64{9, 9, 9, 9}
	updatedcgpa, err := yash.updateSemesterCgpa(new_cgpa)
	if err != nil {
		fmt.Println("Error updating CGPA:", err)
	}
	fmt.Println("THE NEW CGPA IS", updatedcgpa)
	fmt.Println("THE new after updated cgpa OF YASH", yash.finalCgpa)
	fmt.Println("The grades after updating", yash.semesterGrades)
	fmt.Println("The final grades after updating", yash.finalGrade)
	fmt.Println("THE AGE OF YASH IS", yash.age)
	yash.updateFirstName("Sanjeev")
	fmt.Println("first name of yash changed to", yash.firstName)

	fmt.Println("The full name of yash is", yash.fullName)

	yash.updateDob(2221)
	fmt.Println("The dob of yash is", yash.DOB)
	fmt.Println("THE age of yash is", yash.age)

	fmt.Println(yash.semesterGrades)
	fmt.Println(yash.finalGrade)
	fmt.Println(yash.finalCgpa)

}

// make a function to update name
func (s *Student) updateFirstName(firstname string) (string, error) {
	if firstname == " " {
		return "", errors.New("fIRST NAME CANT BE EMPTY")
	}

	s.firstName = firstname
	s.fullName = firstname + s.lastName
	return firstname, nil
}

// function to update last name
func (s *Student) updatelastName(lastname string) (string, error) {
	if lastname == " " {
		return "", errors.New("Last name cant be empty")
	}

	s.lastName = lastname
	s.fullName = s.firstName + lastname // handling full name
	return lastname, nil
}

// function to update date of birth
func (s *Student) updateDob(dob int) (int, error) {
	if dob < 0 {
		return -1, errors.New("date of birth cant be less than 0.")
	}
	s.DOB = dob
	currentYear := time.Now().Year()
	s.age = currentYear - s.DOB // handling current age
	return s.DOB, nil
}

// function to update year of enrollement
func (s *Student) updateYearofEnrollement(yearofenrollement int) (int, error) {
	if yearofenrollement < 0 {
		return -1, errors.New("Year of enrollement cant be less than 0.")
	}

	s.yearofEnrollement = yearofenrollement
	s.NumberofYearToGraduate = yearofenrollement - s.Yearofpassing // handling year to graduate
	return yearofenrollement, nil

}

// function to update year of passing
func (s *Student) updateYearofPassing(yearofpassing int) (int, error) {
	if yearofpassing < 0 {
		return -1, errors.New("year of passing cant be less than 0.")
	}

	s.Yearofpassing = yearofpassing
	s.NumberofYearToGraduate = s.yearofEnrollement - yearofpassing // handling year to graduate
	return yearofpassing, nil
}

// function to update semester cgpa
func (s *Student) updateSemesterCgpa(semestercgpa []float64) ([]float64, error) {
	for _, value := range semestercgpa {
		if value < 0 {
			return nil, errors.New("value cant be less than 0")
		}
	}

	s.semesterCgpa = semestercgpa

	// final cgpa recalculation
	for _, cgpa := range semestercgpa {
		if cgpa < 0.0 {
			return nil, errors.New("The cgpa can't be less than 0")
		}
	}
	total := 0.0
	for _, cgpaValue := range semestercgpa {
		total += cgpaValue
	}
	s.finalCgpa = total / float64(len(semestercgpa))

	// function to update sem grades
	for i, cgpa := range semestercgpa {
		var grade string
		if cgpa >= 9 {
			grade = "A"
		} else if cgpa >= 8 {
			grade = "B"
		} else if cgpa >= 7 {
			grade = "C"
		} else if cgpa >= 6 {
			grade = "D"
		} else if cgpa >= 5 {
			grade = "E"
		} else {
			grade = "F"
		}
		s.semesterGrades[i] = grade
	}

	// final grade of student
	var finalgrade string
	if s.finalCgpa >= 9 {
		finalgrade = "A"
	} else if s.finalCgpa >= 8 {
		finalgrade = "B"
	} else if s.finalCgpa >= 7 {
		finalgrade = "C"
	} else if s.finalCgpa >= 6 {
		finalgrade = "D"
	} else if s.finalCgpa >= 5 {
		finalgrade = "E"
	} else {
		finalgrade = "F"
	}
	s.finalGrade = finalgrade

	return s.semesterCgpa, nil

}
