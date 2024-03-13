// package main

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// )

// var bookInfo = make(map[string]string)

// func main() {
// 	for {
// 		var choice int

// 		fmt.Println("Enter your choice")
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			var bookName, author string
// 			fmt.Println("Enter the name of book")
// 			fmt.Scan(&bookName)
// 			fmt.Println("Enter the name of the author")
// 			fmt.Scan(&author)
// 			if err := addBook(bookName, author); err != nil {
// 				fmt.Println("Error:", err)
// 			}
// 			DisplayBook(bookName)

// 		case 2:
// 			var bookName string
// 			fmt.Println("Enter the name of book you want to remove")
// 			fmt.Scan(&bookName)
// 			if err := borrowedBook(bookName); err != nil {
// 				fmt.Println("Error:", err)
// 			}
// 			DisplayBook(bookName)

// 		case 3:
// 			var bookName string
// 			fmt.Println("Enter the name of book")
// 			fmt.Scan(&bookName)
// 			if err := returnBook(bookName); err != nil {
// 				fmt.Println("Error:", err)
// 			}
// 			DisplayBook(bookName)

// 		case 4:
// 			var bookName string
// 			fmt.Println("Enter the name of book")
// 			fmt.Scan(&bookName)
// 			if err := removeBook(bookName); err != nil {
// 				fmt.Println("Error:", err)
// 			}
// 			DisplayBook(bookName)

// 		default:
// 			return
// 		}
// 	}
// }

// func addBook(bookName, author string) error {
// 	if _, ok := bookInfo[bookName]; ok {
// 		return errors.New("Book already exists")
// 	}

// 	bookInfo[bookName] = fmt.Sprintf("Author: %s, Status: Never issued, Borrowed: false, Returned: false", author)

// 	fmt.Printf("Added book: %v of the author %v\n", bookName, author)
// 	return nil
// }

// func removeBook(bookName string) error {
// 	if _, ok := bookInfo[bookName]; !ok {
// 		return errors.New("Book does not exist")
// 	}

// 	delete(bookInfo, bookName)
// 	fmt.Printf("Removed the book successfully %v\n", bookName)
// 	return nil
// }

// func borrowedBook(bookName string) error {
// 	if _, ok := bookInfo[bookName]; !ok {
// 		return errors.New("Book does not exist")
// 	}

// 	if isReturned(bookInfo[bookName]) {
// 		return errors.New("Book is returned")
// 	}

// 	bookInfo[bookName] = setBookStatus(bookInfo[bookName], "Borrowed")

// 	return nil
// }

// func returnBook(bookName string) error {
// 	if _, ok := bookInfo[bookName]; !ok {
// 		return errors.New("Book does not exist")
// 	}

// 	if !isBorrowed(bookInfo[bookName]) {
// 		return errors.New("Book not borrowed")
// 	}

// 	bookInfo[bookName] = setBookStatus(bookInfo[bookName], "Book Returned")

// 	return nil
// }

// func DisplayBook(bookName string) {
// 	fmt.Println("BOOK Name", bookName)
// 	fmt.Println(bookInfo[bookName])
// }

// func isBorrowed(bookStatus string) bool {
// 	return bookStatusContains(bookStatus, "Borrowed: true")
// }

// func isReturned(bookStatus string) bool {
// 	return bookStatusContains(bookStatus, "Returned: true")
// }

// func bookStatusContains(bookStatus, substring string) bool {
// 	return strings.Contains(bookStatus, substring)
// }

// func setBookStatus(bookStatus, newStatus string) string {
// 	return strings.ReplaceAll(bookStatus, "Status: "+getStatus(bookStatus), "Status: "+newStatus)
// }

// func getStatus(bookStatus string) string {
// 	statusIndex := strings.Index(bookStatus, "Status: ") + len("Status: ")
// 	if statusIndex == -1 {
// 		return ""
// 	}
// 	return bookStatus[statusIndex:]
// }
