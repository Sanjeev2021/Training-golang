package main

import (
	"errors"
	"fmt"
)

var cart = make(map[string]item)

type item struct {
	Name     string
	Price    int
	Quantity int
}

func main() {

	var choice int
	var name string
	var price int
	var quantity int

	forloop := true

	fmt.Println("Enter the name of the item:")
	fmt.Scan(&name)

	fmt.Println("Enter the price of the item:")
	fmt.Scan(&price)

	fmt.Println("Enter the quantity of the item:")
	fmt.Scan(&quantity)

	for forloop {
		fmt.Println("\n enter your choice : ")
		fmt.Scan(&choice)

		if choice < 1 || choice > 5 {
			fmt.Println("Please give your input between 1 and 5")
			break
		}

		things := "phone"
		newItem := item{
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		switch choice {
		case 1:
			addItem(cart, things, newItem)
		case 2:
			getItem(things)
		case 3:
			removeItem(things, &newItem)
		case 4:
			UpdateItem(things, &newItem)
		case 5:
			Totalcost(things, &newItem)
		default:
			forloop = false
		}
		continue
	}
}

// things is my key , newItem is my Value
func addItem(cart map[string]item, things string, newItem item) {

	cart[things] = newItem

	fmt.Printf("the cart has %v item of price %v of quantiy %v which are of type %v  in the cart", cart[things].Name, cart[things].Price, cart[things].Quantity, things)

}

func getItem(things string) (item, error) {

	value, ok := cart[things]
	if !ok {
		fmt.Println("item does not exist")
		return item{}, errors.New("item does not exist")
	}

	// return value, int { can't directly access value inside a struct .. if it was not a struct then this was right instead of that do this }
	fmt.Printf("Item: %s, Price: %d, Quantity: %d\n", value.Name, value.Price, value.Quantity)
	return value, nil

}

func removeItem(things string, newItem *item) (item, error) {
	value, err := cart[things]
	if !err {
		return value, errors.New("item does not exist")
	}

	fmt.Println("\n enter the quantity you want to remove")
	fmt.Scan(&newItem.Quantity)

	if value.Quantity < newItem.Quantity {
		fmt.Println("The value is less than the quantity")
		return value, errors.New("the value is less than the quantity , please enter a lower value")
	}

	value.Quantity -= newItem.Quantity
	cart[things] = value

	fmt.Printf("Removed %v of %v from the cart. Now, there are %v left.\n", newItem.Quantity, things, value.Quantity)
	return cart[things], nil
}

func UpdateItem(things string, newItem *item) (item, error) {
	fmt.Println("Cart before updating", cart)

	value, ok := cart[things]
	if !ok {
		return value, errors.New("item does not exist")
	}

	fmt.Println("\n enter the quantity you want to Update")
	fmt.Scan(&newItem.Quantity)

	value.Quantity = newItem.Quantity
	cart[things] = value

	fmt.Println("inventory after updating", cart)
	return cart[things], nil

}

func Totalcost(things string, newItem *item) error {
	value, ok := cart[things]
	if !ok {
		return errors.New("item does not exist")
	}

	cost := value.Price * value.Quantity
	fmt.Println("\nthe cost of the product is ", cost)

	return nil

}
