package main

import "fmt"

// i need to make a board - 3 x 3
// ask for position from user
// check if position is empty
// check if draw
// check if win

func main() {
	board := initalizeBoard()
	Xturn := true
	for {

		var playerSymbol string
		position := askPosition(playerSymbol)

		// validate input
		if board[position] != " " {
			fmt.Println("The Position is already taken !")
			continue
		}

		if Xturn {
			fmt.Println("player1 Chance")
			playerSymbol = "X"
		} else {
			fmt.Println("player2 Chance")
			playerSymbol = "O"
		}

		board[position] = playerSymbol
		// if Xturn {
		// 	board[position] = "X"
		// } else {
		// 	board[position] = "O"
		// }

		drawBoard(board)
		if Win(board) {
			fmt.Println("Thank you for playing")
			break
		}

		if Draw(board) {
			break
		}
		Xturn = !Xturn

	}
}

func drawBoard(board map[int]string) { // board is an argument - the variable that we take inside a func is argument
	fmt.Printf("\n %v | %v | %v \n", board[0], board[1], board[2])
	fmt.Println("-----------")
	fmt.Printf(" %v | %v | %v \n", board[3], board[4], board[5])
	fmt.Println("-----------")
	fmt.Printf(" %v | %v | %v \n", board[6], board[7], board[8])
}

func askPosition(playerSymbol string) int {

	var position int
	fmt.Printf("Player %s, enter position (1-9): ", playerSymbol)
	fmt.Scan(&position)

	if position < 0 || position > 9 {
		fmt.Println("Pls enter your position between 0 and 8")
		return askPosition(playerSymbol)
	}
	return position - 1
}

func initalizeBoard() map[int]string {
	board := make(map[int]string)
	for i := 0; i <= 8; i++ {
		board[i] = " "
	}
	return board
}

func Win(board map[int]string) bool {

	if board[0] != " " && board[1] == board[2] && board[2] == board[0] {
		fmt.Println("YOU WON")
		return true
	} else if board[0] != " " && board[3] == board[6] && board[6] == board[0] {
		fmt.Println("YOU WON")
		return true
	} else if board[1] != " " && board[4] == board[7] && board[7] == board[1] {
		fmt.Println("YOU WON")
		return true
	} else if board[2] != " " && board[5] == board[8] && board[8] == board[2] {
		fmt.Println("YOU WON")
		return true
	} else if board[3] != " " && board[4] == board[5] && board[5] == board[3] {
		fmt.Println("YOU WON")
		return true
	} else if board[6] != " " && board[7] == board[8] && board[8] == board[6] {
		fmt.Println("YOU WON")
		return true
	} else if board[0] != " " && board[4] == board[8] && board[8] == board[0] {
		fmt.Println("YOU WON")
		return true
	} else if board[2] != " " && board[4] == board[6] && board[6] == board[2] {
		fmt.Println("YOU WON")
		return true
	}
	return false
}

func Draw(board map[int]string) bool {
	for i := 0; i <= 8; i++ {
		if board[i] == " " {
			return false
		}
	}
	fmt.Println("It is a Draw.")
	return true
}
