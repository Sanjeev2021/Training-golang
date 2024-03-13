package main

func main() {
	breakOuterLoop()
}

func breakOuterLoop() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break // the value will break at 5 but continue for further iteration .
			}
		}
	}
}
