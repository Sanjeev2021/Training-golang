package main

import (
	"errors"
	"fmt"
	"strings"
)

var bookStatus = make(map[string]string)

func main() {
	var bookName string
	var author string

	var forloop = true

	for forloop {
		var choice int

		fmt.Println("Enter your choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the name of book")
			fmt.Scan(&bookName)
			fmt.Println("Enter the name of the author")
			fmt.Scan(&author)
			if err := addBook(bookName, author); err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 2:
			fmt.Println("Enter the name of book you want to remove")
			fmt.Scan(&bookName)
			if err := borrowedBook(bookName); err != nil {
				fmt.Println("Error:", err)
				return
			}
			DisplayBook(bookName)

		case 3:
			fmt.Println("Enter the name of book")
			fmt.Scan(&bookName)
			if err := returnBook(bookName); err != nil {
				fmt.Println("Error:", err)
				return
			}
			DisplayBook(bookName)

		case 4:
			fmt.Println("Enter the name of book")
			fmt.Scan(&bookName)
			if err := removeBook(bookName); err != nil {
				fmt.Println("Error:", err)
				return
			}
			DisplayBook(bookName)
		default:
			forloop = false
		}
	}

}

func addBook(bookName string, author string) error {

	if _, ok := bookStatus[bookName]; ok {
		return errors.New("Book already exist")
	}

	bookStatus[bookName] = fmt.Sprintf("Author: %s, Status: Never issued, Borrowed: false, Returned: false", author)

	bookStatus[bookName] = "Never issued" // never issued that's why false
	fmt.Printf("Added book: %v of the author %v\n", bookName, author)

	return nil

}

func removeBook(bookName string) error {

	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist")
	}

	delete(bookStatus, bookName)
	fmt.Printf("Removed the book successfully %v\n", bookName)

	return nil
}

func borrowedBook(bookName string) error {

	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist")
	}

	if isReturned(bookStatus[bookName]) {
		return errors.New("Book is returned")
	}

	bookStatus[bookName] = "Borrowed"
	bookStatus[bookName] = setBookStatus(bookStatus[bookName], "Borrowed")

	return nil
}

func returnBook(bookName string) error {

	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist")
	}

	if !isBorrowed(bookStatus[bookName]) {
		return errors.New("Book not borrowed")
	}

	bookStatus[bookName] = setBookStatus(bookStatus[bookName], "Book Returned")

	return nil

}

// func DisplayBook(bookName string) error {
// 	if _, ok := booklist[bookName]; !ok {
// 		return errors.New("Book does not exist")
// 	}

// 	if _, ok := bookReturned[bookName]; !ok {
// 		return errors.New("Book not returned ")
// 	}

// 	if _, ok := bookborrowed[bookName]; !ok {
// 		return errors.New("Book not borrowed")
// 	}

// 	if _, ok := bookStatus[bookName]; !ok {
// 		return errors.New("The status of the book is ...")
// 	}
// 	fmt.Println("BOOK Name", bookName)
// 	fmt.Println("name of the book borrowed", bookborrowed[bookName])
// 	fmt.Println("name of the book returned", bookReturned[bookName])
// 	fmt.Println("Status of the book", bookStatus[bookName])

// 	return nil
// }

func DisplayBook(bookName string) {
	fmt.Println("BOOK Name", bookName)
	fmt.Println(bookStatus[bookName])
}

func isBorrowed(bookStatus string) bool {
	return bookStatusContains(bookStatus, "Borrowed: true")
}

func isReturned(bookStatus string) bool {
	return bookStatusContains(bookStatus, "Returned: true")
}

func setBookStatus(bookStatus, newStatus string) string {
	return strings.ReplaceAll(bookStatus, "Status: "+getStatus(bookStatus), "Status: "+newStatus)
}

func getStatus(bookStatus string) string {
	statusIndex := strings.Index(bookStatus, "Status: ") + len("Status: ")
	if statusIndex == -1 {
		return ""
	}
	return bookStatus[statusIndex:]
}

func bookStatusContains(bookStatus, substring string) bool {
	return strings.Contains(bookStatus, substring)
}
