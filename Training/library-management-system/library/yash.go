package main

import (
	"errors"
	"fmt"
	"sort"
)

var bookStatus = make(map[string]string)
var quantity = make(map[string]int)

const neverIssued = "Never issued"
const returned = "returned"
const borrowed = "borrowed"

func main() {
	var bookName string
	var oldbookName string
	var newbookName string
	var quantityValue int

	var forloop = true
	for forloop {
		var choice int
		fmt.Println("\nEnter your choice ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the book name you want to add")
			fmt.Scan(&bookName)
			fmt.Println("Enter the quantity of the book")
			fmt.Scan(&quantityValue)
			err := addBook(bookName, quantityValue)
			if err != nil {
				fmt.Println("Error", err)
			}
			if err = display(bookName); err != nil {
				fmt.Println("Error", err)
			}

		case 2:
			fmt.Println("Write the name of the book you want to remove")
			fmt.Scan(&bookName)
			err := removeBook(bookName)
			if err != nil {
				fmt.Println("Error", err)
			}
			if err = display(bookName); err != nil {
				fmt.Println("Error", err)
			}
		case 3:
			fmt.Println("Enter the name of the book you want to update")
			fmt.Scan(&oldbookName)
			fmt.Println("Enter the name you want to update")
			fmt.Scan(&newbookName)
			err := updateBookName(oldbookName, newbookName)
			if err != nil {
				fmt.Println("Error", err)
			}
			if err = display(bookName); err != nil {
				fmt.Println("Error", err)
			}
		case 4:
			fmt.Println("Enter the bookname to check it's status")
			fmt.Scan(&bookName)
			err := issueBook(bookName, quantityValue)
			if err != nil {
				fmt.Println("Error", err)
			}
			if err = display(bookName); err != nil {
				fmt.Println("Error", err)
			}
		case 5:
			fmt.Println("Enter the bookname to check it's status")
			fmt.Scan(&bookName)
			err := returnBook(bookName)
			if err != nil {
				fmt.Println("Error", err)
			}
			if err = display(bookName); err != nil {
				fmt.Println("Error", err)
			}
		case 6:
			fmt.Println("Enter the name of the book you want to display")
			err := display(bookName)
			if err != nil {
				fmt.Println("Error", err)
			}
		case 7:
			err := displayAllBook()
			if err != nil {
				fmt.Println("Error", err)
			}
		default:
			forloop = false
		}

	}
}
func addBook(bookName string, quantityValue int) error {
	//boookSatus main add karna hai and initial status uska never issued karna  hai
	if _, ok := bookStatus[bookName]; ok {
		return errors.New("Book already exist!")
	}

	if _, ok := quantity[bookName]; ok {
		return errors.New("quantity already exist!")
	}

	bookStatus[bookName] = neverIssued
	quantity[bookName] = quantityValue

	fmt.Printf("The book %v added successfully with status %v and quantity %v \n", bookName, bookStatus[bookName], quantity[bookName])
	return nil

}
func removeBook(bookName string) error {
	//delete book
	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist!")
	}

	delete(bookStatus, bookName)
	fmt.Printf("Removed the book successfully %v\n", bookName)

	return nil
}
func updateBookName(oldbookName, newbookName string) error {
	// update book name
	if _, ok := bookStatus[oldbookName]; !ok {
		return errors.New("Book does not exist!")
	}

	bookStatus[newbookName] = bookStatus[oldbookName]

	delete(bookStatus, oldbookName)

	fmt.Printf("Updated item name from %v to %v\n", oldbookName, newbookName)

	return nil

}
func issueBook(bookName string, quantityValue int) error {
	//isme tu bookStatus ko boorowed karega agar status never  isssued ya fir reurned hai
	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist !")
	}

	if quantityValue <= 0 {
		return errors.New("The quantity is 0 , so cant issue the book")
	}

	fmt.Println("It is going for check")
	if bookStatus[bookName] != neverIssued || bookStatus[bookName] != returned {
		bookStatus[bookName] = borrowed
		fmt.Printf("The bookName %v is borrowed", bookName)
	}

	return nil

}
func returnBook(bookName string) error {
	//isme tu bookStatus ko returned karega agar status borrowed hai toh
	if _, ok := bookStatus[bookName]; !ok {
		return errors.New("Book does not exist !")
	}

	if bookStatus[bookName] == borrowed {
		bookStatus[bookName] = returned
		fmt.Printf("The returned book name is %v", bookName)
	}
	return nil
}

func display(bookName string) error {
	fmt.Printf("Book name %v of status %v", bookName, bookStatus[bookName])
	return nil
}

func displayAllBook() error {

	for title, details := range bookStatus {
		fmt.Printf("\nThe book names present are %v  of status type %v", title, details)
	}

	keys := make([]string, 0, len(quantity))

	for k := range quantity {
		fmt.Println("quantity", k)
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return quantity[keys[i]] > quantity[keys[j]]
	})

	fmt.Print("the name of the book and quantity: ")
	for _, key := range keys {
		fmt.Printf("%s (%v) ", key, quantity[key])
	}
	fmt.Println()
	return nil
}
