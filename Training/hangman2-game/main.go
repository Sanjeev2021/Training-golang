package main

import "fmt"

func main() {
	// guess := 0
	life := 5

	words := [3]string{"sanjeev", "yadav", "go"}

	for wordinarray := 0; wordinarray < len(words); wordinarray++ {
		word1 := words[wordinarray]
		word := make([]string, len(word1))
		for i := 0; i < len(word1); i++ {
			word[i] = "_"
		}
		for true {
			fmt.Println(word)
			var input string
			fmt.Println("Guess the letter ")
			fmt.Scanln(&input)

			var flagForCorrectInput bool = false
			for i := 0; i < len(word1); i++ {
				if input == string(word1[i]) {
					word[i] = input
					flagForCorrectInput = true
				}
			}

			if flagForCorrectInput == true {
				fmt.Println("your guess was right , keep trying")
			} else {
				life--
				fmt.Println("you are wrong ... you have ", life, " live left \n Thank you for playing")
			}

			if life == 0 {
				break
			}

			var wordGuessed bool
			wordGuessed = true
			for i := 0; i < len(word1); i++ {
				if word[i] != string(word1[i]) {
					wordGuessed = false
				}
			}
			if wordGuessed {
				fmt.Println("congrats you guessed it")

				break
			}
		}

		if wordinarray < len(word)-1 {
			var play string
			fmt.Println("do you want to play next word ? (yes/no)")
			fmt.Scanln(&play)

			if play != "yes" {
				break
			}

		}

	}

	// now check for life and guess

	// chances()
}

func Randomword() {
	words := [3]string{"sanjeev", "yadav", "go"} // made an array over here of size 3
	fmt.Println("the value in array", words)

}

func UserInput() int {
	var input int
	fmt.Println("Guess the word ")
	fmt.Scanln(&input)
	return input
}

func chances(life, guess int) {
	if life == guess {
		fmt.Println("you have %v chances", life)
	}

	if guess > life {
		fmt.Println("sorry you lost the game")
	}

	if life > guess {
		fmt.Println("you have chances keep on playing.")
	}
}

// func printCharacters(characters []string) {

// 	reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter a character: ")
//     char, _, err := reader.ReadRune() // Rune -
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     fmt.Println("You entered the character:", char)

// 	wordFounded := false
// 	for  i =:0 ; i<len(word) ; i ++ {
// 		if word[i] == char
// 	}
// }
