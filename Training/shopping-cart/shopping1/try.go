package main

import (
	"errors"
	"fmt"
)

var price = make(map[string]float64)
var quantity = make(map[string]int)

func main() {
	fmt.Println(price)
	fmt.Println(quantity)
	var itemName string
	var itemPrice float64
	var itemQuantity int
	var err error
	var oldItem string
	var newItemName string
	for true {
		var choice int
		fmt.Println("Choice Menu:")
		fmt.Println("1. add item")
		fmt.Println("2. update price of an exisiting item")
		fmt.Println("3. purchase quantity of an exisiting item")
		fmt.Println("4. Sell quantity of an existing item")
		fmt.Println("5. Total worth of the item")
		fmt.Println("6. Total worth of all the item")
		fmt.Println("7. update item name")
		fmt.Println("8. updated item name")
		// fmt.Println("2. update price of an exisiting item")
		fmt.Println("Enter your choice")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter your item name")
			fmt.Scan(&itemName)
			fmt.Println("Enter your item price")
			fmt.Scan(&itemPrice)
			fmt.Println("Enter your item quantity")
			fmt.Scan(&itemQuantity)
			err = addItem(itemName, itemPrice, itemQuantity)
			if err != nil {
				fmt.Println("retry additem error", err)
			}
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("retry display item", err)
			}
		case 2:
			fmt.Println("enter your item name to update")
			fmt.Scan(&itemName)
			fmt.Println("enter your item price")
			fmt.Scan(&itemPrice)
			err := updateItemPrice(itemName, itemPrice)
			if err != nil {
				fmt.Println("Retry")
			}
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("retry")
			}
		case 3:
			fmt.Println("enter your item name")
			fmt.Scan(&itemName)
			fmt.Println("enter your purchase quantity")
			fmt.Scan(&itemQuantity)
			err := purchaseExisistingItem(itemName, itemQuantity)
			if err != nil {
				fmt.Println("retry")
			}
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("RETRY")
			}
		case 4:
			fmt.Println("enter the item name:")
			fmt.Scan(&itemName)
			fmt.Println("enter the quantity you want to sell:")
			fmt.Scan(&itemQuantity)
			err := purchaseExisistingItem(itemName, itemQuantity)
			if err != nil {
				fmt.Println("retry")
			}
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("Retry")
			}
		case 5:
			fmt.Println("Enter the item name: ")
			fmt.Scan(&itemName)
			result, err := totalWorthOfAnItem(itemName)
			fmt.Println("Worth of item is: ", result)
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("Retry")
			}
		case 6:
			err := totalWorthOAllItem()
			if err != nil {
				fmt.Println("Retry")
			}
		case 7:
			fmt.Println("Enter the item name: ")
			fmt.Scan(&itemName)
			err := updateItemName(itemName)
			if err != nil {
				fmt.Println("Retry update")
			}
			if err = dispalyItem(itemName); err != nil {
				fmt.Println("Retry update")
			}
		case 8:
			fmt.Println("Enter the old item name: ")
			fmt.Scan(&oldItem)

			fmt.Println("Enter the new item name: ")
			fmt.Scan(&newItemName)

			err := updatedItemName(newItemName, oldItem)
			if err != nil {
				fmt.Println("Failed to update item name:", err)
			}

		}
	}

}
func addItem(itemName string, itemPrice float64, itemQuantity int) error {
	if itemPrice < 0 {
		return errors.New("invalid item price")
	}
	if itemQuantity < 0 {
		return errors.New("invalid item quantity")
	}
	_, ok := price[itemName]
	if ok {
		fmt.Println(ok)
		return errors.New("item already exisit price")
	}
	_, ok = quantity[itemName]
	if ok {
		return errors.New("item already exisit, quantity")
	}
	price[itemName] = itemPrice
	quantity[itemName] = itemQuantity
	fmt.Println("Item Added Successfully")
	return nil
}
func updateItemPrice(itemNameToBeUpdated string, newItemPrice float64) error {
	if newItemPrice < 0 {
		return errors.New("invalid item price")
	}
	if _, ok := price[itemNameToBeUpdated]; !ok {
		return errors.New("item doesnot exist")
	}
	price[itemNameToBeUpdated] = newItemPrice
	fmt.Println("Item price updated Successfully")
	return nil
}
func purchaseExisistingItem(itemName string, purchaseQuantity int) error {
	if purchaseQuantity < 0 {
		return errors.New("invalid item quantity")
	}
	if _, ok := quantity[itemName]; !ok {
		return errors.New("item doesnot exist")
	}
	quantity[itemName] += purchaseQuantity
	fmt.Println("total Item quantity after purchase : ", quantity[itemName])
	return nil
}
func dispalyItem(itemName string) error {
	if _, ok := quantity[itemName]; !ok {
		return errors.New("item doesnot exist")
	}
	if _, ok := price[itemName]; !ok {
		return errors.New("item doesnot exist")
	}
	fmt.Println("Item Name: ", itemName)
	fmt.Println("Item Price: ", price[itemName])
	fmt.Println("Item quantity: ", quantity[itemName])
	return nil
}

// make sell exisiting item
func sellExisitngItem(itemName string, sellQuantity int) error {
	if sellQuantity < 0 {
		return errors.New("invalid item quantity")
	}

	if _, ok := quantity[itemName]; !ok {
		return errors.New("item doesnot exist")
	}

	if sellQuantity > quantity[itemName] {
		return errors.New("Pls input a lower value")
	}
	quantity[itemName] -= sellQuantity
	fmt.Println("sold these many items")
	fmt.Printf("the sold values are %v", quantity[itemName])
	return nil
}

// 1 item value iphone = quantity * price
func totalWorthOfAnItem(itemName string) (float64, error) {

	if _, ok := quantity[itemName]; !ok {
		return 0.0, errors.New("item does not exist!")
	}

	cost := float64(quantity[itemName]) * price[itemName]

	return cost, nil
}

// all the item .. print the totalWorth of all item .. sara item & all thier name
func totalWorthOAllItem() error {

	worth := 0.0

	for itemName, itemPrice := range price {
		itemQuantity, ok := quantity[itemName]
		if !ok {
			return errors.New("Something is wrong")
		}
		worth = float64(itemQuantity) * itemPrice
		fmt.Printf("the worth of %v is %v \n", itemName, worth)
	}
	return nil

}

func updateItemName(itemName string) error {
	if _, ok := price[itemName]; ok {
		return errors.New("Item name does not exist")
	}

	if _, ok := quantity[itemName]; ok {
		return errors.New("item name does not exist")
	}

	price[itemName] = price[itemName]
	quantity[itemName] = quantity[itemName]

	fmt.Printf("the new item name for price is %v ", price[itemName])
	fmt.Printf("the new item name for quantity is %v", quantity[itemName])
	fmt.Printf("the item name is ", itemName)

	return nil

}

func updatedItemName(newItemName, oldItem string) error {
	if _, ok := price[oldItem]; !ok {
		return errors.New("Item name does not exist")
	}

	if _, ok := quantity[oldItem]; !ok {
		return errors.New("Item name does not exist")
	}

	price[newItemName] = price[oldItem]
	quantity[newItemName] = quantity[oldItem]

	delete(price, oldItem)
	delete(quantity, oldItem)

	fmt.Printf("Updated item name from %s to %s\n", oldItem, newItemName)

	return nil

}
