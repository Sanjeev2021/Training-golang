package main

import "fmt"

func main() {
	var num1 int
	fmt.Println("enter your num1")
	fmt.Scan(&num1)

	var num2 int
	fmt.Println("enter your num2")
	fmt.Scan(&num2)

	fmt.Println("the value before swapping", num1, num2)

	//now swapping begin

	if num1 == 0 {
		num1 = num2
		num2 = 0
	} else if num2 == 0 {
		num2 = num1
		num1 = 0
	} else {
		num1 = num1 * num2
		num2 = num1 / num2
		num1 = num1 / num2
		num1 = num1 + num2
		num2 = num1 - num2
		num1 = num1 - num2
	}

	fmt.Println("the value after swapping", num1, num2)
}
