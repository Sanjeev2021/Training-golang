// package main

// import (
// 	"errors"
// 	"fmt"
// )

// var cart = make(map[string]map[int]float64)

// func main() {

// 	// var choice int
// 	for {
// 		var name string
// 		var price float64
// 		var quantity int
// 		var item string
// 		var newQuantity int

// 		// forloop := true

// 		fmt.Println("enter the item you want to add:")
// 		fmt.Scan(&item)

// 		fmt.Println("Enter the name of the item:")
// 		fmt.Scan(&name)

// 		fmt.Println("Enter the price of the item:")
// 		fmt.Scan(&price)

// 		fmt.Println("Enter the quantity of the item:")
// 		fmt.Scan(&quantity)

// 		addItem(item, float64(price), quantity, name)

// 		result, err := getItem(item)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Printf("Item details: %v\n", result)
// 		}

// 		removeItem(item, quantity, price)

// 		UpdateItem(item, newQuantity, quantity, price)
// 	}
// }

// func addItem(item string, price float64, quantity int, name string) error {

// 	_, ok := cart[item]
// 	if !ok {
// 		cart[item] = make(map[int]float64)
// 		cart[item][quantity] = float64(price)
// 		fmt.Printf("The store now has %v of %v in cart %v and %v quantity %v.\n", cart[item], item, name, price, quantity)
// 	} else {
// 		cart[item][quantity] += float64(price) // here the value was beign mismatched
// 		return errors.New("item does exist")
// 	}

// 	return nil
// }

// func getItem(item string) (map[int]float64, error) {
// 	value, err := cart[item]
// 	if !err {
// 		fmt.Println("item does not exist")
// 	}

// 	return value, nil
// }

// func removeItem(item string, quantity int, price float64) (map[int]float64, error) {
// 	value, err := cart[item]
// 	if !err {
// 		fmt.Println("item does not exist")
// 		return value, errors.New("item does not exist")
// 	}

// 	if quantity <= 0 {
// 		fmt.Println("invalid quantity, please provide a valid quantity")
// 		return value, errors.New("invalid quantity, please provide a valid quantity")
// 	}

// 	cart[item][quantity] -= float64(quantity)
// 	fmt.Printf("The store now has %v of %v in inventory.\n", cart[item], item)

// 	return cart[item], nil

// }

// func UpdateItem(item string, newQuantity int, quantity int, price float64) (map[int]float64, error) {
// 	fmt.Println("cart before updating", cart)

// 	_, ok := cart[item][quantity]
// 	if !ok {
// 		return cart[item], errors.New("oldItem does not exist")
// 	}

// 	if _, ok = cart[item][newQuantity]; ok {
// 		return cart[item], errors.New("NewItem does not exist")
// 	}

// 	delete(cart[item], quantity)
// 	fmt.Println("the store now has %v of %v in cart", cart[item], item)

// 	fmt.Println("ENTER THE NEW QUANTITY")
// 	fmt.Scan(&newQuantity)

// 	cart[item][quantity] = float64(newQuantity)

// 	fmt.Printf("The store now has %v of %v in inventory.\n", newQuantity, item)

// 	return cart[item], nil

// }
