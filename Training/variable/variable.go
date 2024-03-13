// ways to declare variables in go

package main

import "fmt"

func main() {
	// ways to declare variables

	// so basically there are 3 ways to declare variable
	// what is a variable ? variable is something that holds a value

	//2 ways to declare variable

	//1 using var keyword

	var sanjeev int
	sanjeev = 1
	fmt.Println("variable declaration", sanjeev)

	//2 shorthand declaration

	sahil := 2
	fmt.Println("short-hand-declaration", sahil)
	fmt.Println("TYPE OF VARIABLE %T\n", sahil)

	// cannot declare a variable again of the same name ?
	var a string
	a = "max"
	fmt.Println(a)

	var c bool
	c = true
	fmt.Println(c)

	var d float32
	d = 10.5
	fmt.Println(d)

	var e float64
	e = 11.3
	fmt.Println(e)

	var f bool
	f = true
	fmt.Println(f)

	type g struct {
		x int
		y int
	}

	sam := g{
		x: 1,
		y: 1,
	}
	fmt.Println("the value of sam", sam.x)
	fmt.Println("the value of y sam", sam.y)

	var ptr *int
	q := 12
	ptr = &q
	fmt.Println("THE address of q", &q)
	fmt.Println("the value of ptr", ptr)
	fmt.Println("the address of pointer", &ptr)
	fmt.Println("the value inside pointer", *ptr)

	array := [3]int{10, 20, 30}
	fmt.Println("the value inside array", array)
	fmt.Println("the address of array", &array[0])

	slice := []int{10, 20, 30}
	fmt.Println("the value inside of slice", slice)
	slice = append(slice, 70)
	fmt.Println("the value inside slice", slice)

	// var complex complex64 = complex(1, 2)
	// fmt.Println("the complex number is", complex)

	// you cant do var complex complex64 := complex(1,2) because that is short hand declaration
}

// the list of data types in go
// slice
// array
// pointer
// int
// string
// bool
// float32
// float64
// int32
//int64
// struct
