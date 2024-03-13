package main

import "fmt"

func main() {
	var result int
	var n1, n2 int

	forloop := true

	for forloop {
		fmt.Println("enter your choice :")
		var choice int
		fmt.Scan(&choice)

		if choice < 1 || choice > 5 {
			fmt.Println("enter your input between 1 and 5 ")
			break
		}

		fmt.Println("Enter the value of 1st number")
		fmt.Scan(&n1)

		fmt.Println("enter the value of 2nd number")
		fmt.Scan(&n2)

		switch choice {
		case 1:
			fmt.Println("YOU CHOSE ADDITION")
			result = addition(n1, n2)
			fmt.Println(result)
			//	continue // ifi add continue then it will keep running the for loop even after adding break
			break
		case 2:
			fmt.Println("YOU SELECTED DIVISION")
			result = divide(n1, n2)
			fmt.Println(result)
			break
		case 3:
			fmt.Println("you selected multiply")
			result = multiply(n1, n2)
			fmt.Println(result)
			break
		case 4:
			fmt.Println("you selected subtract")
			result = subtract(n1, n2)
			fmt.Println(result)
			break
		case 5:
			fmt.Println("you selected modulous")
			result = modulos(n1, n2)
			fmt.Println(result)
			break
		default:
			break
		}
		continue
	}
	forloop = false
}

// let the use give number then do the operation.
func addition(n1, n2 int) int {
	return n1 + n2
}

func divide(n1, n2 int) int {

	divide := n1 / n2
	return divide
}

func multiply(n1, n2 int) int {
	multiply := n1 * n2
	return multiply
}

func subtract(n1, n2 int) int {
	subtract := n2 - n1
	return subtract
}

func modulos(n1, n2 int) int {
	modulos := n1 % n2
	return modulos
}
