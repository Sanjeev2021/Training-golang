package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i + j)
		}
		fmt.Println()
	}
}
