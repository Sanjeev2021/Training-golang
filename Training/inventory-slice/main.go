package main

import (
	"errors"
	"fmt"
)

var inventory = make(map[string]int)

func main() {
	var item string
	var oldItem string
	var newItem string
	var qunatity int
	var choice int

	forloop := true

	for forloop {
		fmt.Println("enter your choice")
		fmt.Scan(&choice)

		fmt.Println("ENTER THE ITEM NAME: ")
		fmt.Scan(&item)

		fmt.Println("enter the qunatity of item :")
		fmt.Scan(&qunatity)

		switch choice {
		case 1:
			fmt.Println("you choosed addItems")
			addItem(item, qunatity)
		case 2:
			fmt.Println("You choosed removeItems")
			removeItems(item, qunatity)
		case 3:
			fmt.Println("you choosed get Items")
			fmt.Println("ENTER THE ITEM NAME: ")
			fmt.Scan(&item)
			getItem(item)
		case 4:
			fmt.Println("you chose update element")
			fmt.Println("ENTER THE OLD ITEM NAME: ")
			fmt.Scan(&oldItem)
			fmt.Println("ENTER THE NEW ITEM NAME: ")
			fmt.Scan(&newItem)
			updateItem(oldItem, newItem, qunatity)
		default:
			break
		}
		continue
	}
	forloop = false
}

func addItem(item string, quantity int) error {

	_, ok := inventory[item]
	if ok {
		fmt.Println("item does exist ")
		inventory[item] += quantity
		fmt.Printf("THE STORE HAS THESE IN INVENTORIES %v\n of with added  %v quantity", item, inventory[item])
	} else {
		inventory[item] = quantity
		fmt.Printf("THE STORE HAS THESE IN INVENTORIES %v\n of %v quantity", item, quantity)
		return errors.New("item does exist")
	}

	return nil
}

func removeItems(item string, quantity int) (int, error) {

	// check if item exist or not
	value, err := inventory[item]
	if !err {
		fmt.Println("item does not exist")
		return -1, errors.New("item does not exist")
	}

	if quantity <= 0 || quantity > value {
		fmt.Println("invalid quantity, please provide a valid quantity")
		return -1, errors.New("invalid quantity, please provide a valid quantity")
	}

	// code is redundant over here
	// _, ok := inventory[item]
	// if ok {
	// 	fmt.Println("item does exist ")
	// 	inventory[item] -= quantity
	// 	fmt.Printf("THE STORE HAS THESE IN INVENTORIES %v\n of with added  %v quantity", item, inventory[item])
	// } else {
	// 	return -1, errors.New("item does not exist")
	// }

	inventory[item] -= quantity
	fmt.Printf("The store now has %v of %v in inventory.\n", inventory[item], item)

	return value, nil // here return the inventory
}

func getItem(item string) (int, error) {

	value, err := inventory[item]
	if !err {
		fmt.Println("item does not exist")
	}

	return value, nil
}

func updateQuantityItem(item string, quantity int) (int, error) {
	fmt.Println("inventory before updating", inventory)

	value, ok := inventory[item]
	if !ok {
		return -1, errors.New("item does not exist") // doubt ye explain or like better way to write this.
	}

	if quantity == inventory[item] {
		return -1, errors.New("item already exist")
	}

	inventory[item] = quantity

	fmt.Println("inventory after updating", inventory)
	return value, nil

}

// had to do this ... to update the value .

// update Key in the map
func updateItem(oldItem string, newItem string, quantity int) (int, error) {
	fmt.Println("inventory before updating", inventory)

	oldValue, ok := inventory[oldItem]
	if !ok {
		return -1, errors.New("oldItem does not exist") // HERE ICANT USE .. OldItem because it should be oldValue .. which will hold the value
	}

	if _, ok = inventory[newItem]; ok {
		return -1, errors.New("NewItem does not exist")
	}

	delete(inventory, oldItem)

	inventory[newItem] = quantity

	fmt.Println("The new value is updating", inventory)

	return oldValue, nil

}
