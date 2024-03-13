package main

import (
	"errors"
	"fmt"
)

var arr [3]int
var top = -1
var size = len(arr)

func main() {
	fmt.Println("Is the stack empty?", isEmpty())
	fmt.Println("the first element is :", peak())
	push(10)
	push(20)
	push(30)
	fmt.Println("the first element is :", peak())
	fmt.Println("The element inside array is:")
	ArrElement(arr[:])
	poppedElement, err := pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped element:", poppedElement)
	}
	push(2)
	fmt.Println("The element inside array is:")
	ArrElement(arr[:])

}

func push(element int) {
	if top < size-1 {
		top++
		arr[top] = element
	} else {
		fmt.Println("The stack is full , Stack overflow")
	}
}

func pop() (int, error) {
	if top == -1 {
		return -1, errors.New("THE stack is empty")
	} else {
		var res = arr[top]
		top--
		return res, nil
	}
}

func peak() int {
	if top >= 0 && top < size {
		return arr[top]
	} else {
		fmt.Println("The stack is empty")
		return -1
	}
}

func isEmpty() bool {
	return top == -1
}

func ArrElement(arr []int) {
	for _, val := range arr {
		fmt.Println("THE VALUE INSIDE IS", val)
	}
}
