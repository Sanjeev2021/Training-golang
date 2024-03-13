package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{10, 20, 30, 40}
	// a3 := [4]int{100, 200, 300, 400}
	a1 = a2
	// copy(a1, a2)
	a1[0] = 1000
	fmt.Println("a1[0]= ", a1) //1000
	fmt.Println("a2[0]= ", a2) //1000

}

// func add(a []int) { // pass by reference
// 	for i := 1; i < 3; i++ {
// 		a[0] += a[i]
// 	}
// 	fmt.Println("inside add function a[0]= ", a[0]) //12
// }

// func main() {
// 	a1 := []int{3, 4, 5}
// 	add(a1)
// 	fmt.Println("From main a1[0]= ", a1[0]) //12
// }

// func add(a [3]int) { // pass by value is happening over here
// 	for i := 1; i < 3; i++ {
// 		a[0] += a[i]
// 	}
// 	fmt.Println("inside add function a[0]= ", a[0]) //12
// }

// func main() {
// 	a1 := [3]int{3, 4, 5}
// 	add(a1)
// 	fmt.Println("From main a1[0]= ", a1[0]) //3
// }

// func main() {
// 	a := []int{1, 2, 3, 6, 7, 7}
// 	result := add(a...)
// 	fmt.Println("the result is", result)
// }

func add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	ntype := reflect.TypeOf(nums)
	fmt.Println("THE TYPE OF nums is", ntype)
	return total
}
