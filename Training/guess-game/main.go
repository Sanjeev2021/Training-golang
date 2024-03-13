package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100)
	fmt.Println(random)
	for true {
		input := userInput()
		guess(random, input)
		if input == random {
			break // Exit the inner loop if the guess is correct
		}
	}

}

func userInput() int {
	var input int
	fmt.Printf("\n enter your input : ")
	fmt.Scan(&input)
	return input
}

func guess(random, input int) {

	if input > random {
		if input >= random-5 {
			fmt.Println("you are just there , try guessing lower")
		}

		fmt.Println("too high , guess low")
	} else if input < random {
		if input <= random+5 {
			fmt.Println("you are just there , try guessing higher")
		}
		fmt.Println("TOO LOW , GUESS HIGH")
	} else if input == random {
		fmt.Println("YOU GUESSED IT RIGHT")
	}
}

// func randome() {
// 	random := rand.Intn(100)
// 	fmt.Println(random)
// 	fmt.Println("the random number is generated guess it !")
// }
