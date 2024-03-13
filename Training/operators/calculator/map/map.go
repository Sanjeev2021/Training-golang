package main

import "fmt"

func main() {
	// maps := make([]int, 5) // this is a slice actually

	// maps[0] = 1
	// maps[1] = 4
	// maps[2] = 6
	// maps[3] = 9
	// maps[4] = 2

	// fmt.Println("The maps", maps)

	// // for loop in go we use range
	// for key, value := range maps {
	// 	fmt.Println("for key %v , the value is %v\n", key, value)
	// }

	myMap := make(map[string]*int, 5)

	value1 := 1
	value2 := 2
	myMap["sanjeev"] = &value1
	myMap["max"] = &value2

	fmt.Println("the value inside value1", *myMap["sanjeev"])
	fmt.Println("the value inside value2", myMap["max"])

	for key, value := range myMap {
		fmt.Println("for key %v , the value is %v\n", key, value)
	}
}

// %d: decimal integer
// %f: floating-point decimal
// %s: string
// %t: boolean
// %b: base 2 (binary)
// %o: base 8 (octal)
// %x: base 16 (hexadecimal)
// %c: character
// %p: pointer
