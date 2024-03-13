package main

import "fmt"

func main() {
	var result = [2][2]int{}

	// var a1 = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 2, 6},
	// 	{7, 8, 9},
	// }

	// var a2 = [3][3]int{
	// 	{11, 12, 13},
	// 	{14, 12, 16},
	// 	{17, 18, 19},
	// }

	var a1 = [2][2]int{}
	var a2 = [2][2]int{}

	fmt.Println("Enter the value for a1")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&a1[i][j])
		}
	}

	fmt.Println("Enter the value for a2")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&a2[i][j])
		}
	}

	for i := 0; i < 2; i++ { // row iteration
		for j := 0; j < 2; j++ { // column iteration
			result[i][j] = 0
			for k := 0; k < 2; k++ { //
				result[i][j] += a1[i][k] * a2[j][k]
				fmt.Println("the result is", result[i][j])
			}
			fmt.Println()
		}

		// after multplication
		// this will iterate over every row and column after multiplication to make it look more presentable
		fmt.Println("The result is ")
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				fmt.Printf(" %d", result[i][j])
			}
			fmt.Println()
		}

	}

}
