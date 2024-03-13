package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	randomWord()
}

func randomWord() {

	// using get method
	url := "https:api.api-ninjas.com/v1/randomword"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err)
	}

	//using http request to take the response
	response := http.DefaultClient
	resp, err := response.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // ioutil is deprecated so use io.
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("the word is :", string(body))

}

// https://api.api-ninjas.com/v1/randomword

// this happened because my memory was pointing at a nil value and it went into panic mode .. to resolve panic we always use resolve . with defer

func hangMan() {
	fmt.Println(
		"  +---+\n  |   |\n      |\n      |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n      |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n  |   |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|   |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n /    |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n / \\  |\n      |\n=========",
	)
}
