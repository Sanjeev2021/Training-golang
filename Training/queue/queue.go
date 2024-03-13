package main

import "fmt"

var arr []int
var front int
var rear int
var size int

func main() {
	size = 10000
	front = 0
	rear = 0
	arr = make([]int, size)

	fmt.Println("Is the queue empty ?", isEmpty())
	enqueue(1)
	enqueue(2)
	enqueue(3)
	fmt.Println("THE FRONT OF THE QUEUE IS ", isFront())
	resetQueue()
	dequeue()
	dequeue()
	dequeue()
	fmt.Println("THE FRONT OF THE QUEUE IS ", isFront())
	fmt.Println("Is the queue empty", isEmpty())
	fmt.Println("The queue has been reset", resetQueue())
}

func enqueue(data int) {
	if rear == size {
		fmt.Println("The array is full")
	} else {
		arr[rear] = data
		rear++
	}
}

func dequeue() int {

	if front == rear {
		return -1
	} else {
		ans := arr[front]
		arr[front] = -1
		front++
		if rear == front {
			front = 0
			rear = 0
		}
		return ans
	}
}

func isEmpty() bool {
	if front == rear {
		fmt.Println("The queue is empty")
		return true
	} else {
		fmt.Println("The queue has some value")
		return false
	}

}

func isFront() int {
	if front == rear {
		return -1
	} else {
		return arr[front]
	}
}

func resetQueue() int {
	if isEmpty() {
		front = 0
		rear = 0
	}
	fmt.Println("There are some elements")
	return -1
}
