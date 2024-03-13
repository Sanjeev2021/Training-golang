package main

import (
	"errors"
	"fmt"
	"sort"
)

var task = make(map[string]int)
var taskStatus = make(map[string]bool) //true : completed false - pending
var pendingTask = make(map[string]int)
var completedTasks = make(map[string]bool)

func main() {
	fmt.Println(task)
	fmt.Println(pendingTask)
	fmt.Println(completedTasks)

	var taskName string
	var priority int

	for true {
		var choice int
		fmt.Println("\nChoice Menu:")
		fmt.Println("1. add task")
		fmt.Println("2. get task")
		fmt.Println("3. remove task")
		fmt.Println("4. mark as completed ")
		fmt.Println("5. All pending task.")
		fmt.Println("Enter your choice")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			fmt.Scan(&taskName)

			fmt.Print("Enter task priority (integer): ")
			fmt.Scan(&priority)
			err := addTask(taskName, priority)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			result, err := getTask(taskName, priority)
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Printf("the task in the list is of name %v with priority %v", taskName, result)
		case 3:
			fmt.Print("Enter task name: ")
			fmt.Scan(&taskName)

			fmt.Print("Enter task priority (integer): ")
			fmt.Scan(&priority)
			err := removeTask(taskName, priority)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			err := markComplete(taskName, priority)
			if err != nil {
				fmt.Println("Error", err)
			}
		case 5:
			err := AllpendingTask(taskName, priority)
			if err != nil {
				fmt.Println("Error", err)
			}
		}

	}
}

func addTask(taskName string, priority int) error {
	if _, ok := task[taskName]; ok {
		return errors.New("Task already exists")
	}

	task[taskName] = priority
	fmt.Println("pendingTask", pendingTask)
	// taskStatus[taskName] = false
	fmt.Printf("Added task: %v with priority %d\n", taskName, priority)

	return nil
}

func getTask(taskName string, priority int) (int, error) {
	value, err := task[taskName]
	if !err {
		return value, errors.New("Task does not exists")
	}

	return value, nil
}

func removeTask(taskName string, priority int) error {
	value, err := task[taskName]
	if !err {
		return errors.New("Task does not exists")
	}

	newPriority := value - priority
	task[taskName] = newPriority

	if newPriority <= 0 {
		delete(task, taskName)
		fmt.Printf("\nRemoved the task successfully of name %v\n", taskName)
	} else {
		task[taskName] = newPriority
		fmt.Printf("\nUpdated priority for task %v to %v\n", taskName, newPriority)
	}

	return nil

}

//

func markComplete(taskName string, priority int) error {
	fmt.Printf("THE PREVIOUS TASK IS %v", taskName)

	if _, ok := task[taskName]; !ok {
		return errors.New("Task does not exists")
	}
	if completedTasks[taskName] {
		return errors.New("\nTask is already completed , there are no pending tasks")
	}

	// delete(task, taskName)
	completedTasks[taskName] = true

	fmt.Println("pendingTask", pendingTask)
	fmt.Printf("\nThe task is completed %v", taskName)

	return nil
}

func AllpendingTask(taskName string, priority int) error {

	if _, ok := task[taskName]; !ok {
		return errors.New("Task does not exists")
	}

	if completedTasks[taskName] {
		return errors.New("\nTask is already completed , there are no pending tasks")
	}

	pendingTask[taskName] = priority

	//this is the sorting part .. basically i copied from google need to understand this !
	// i will write this on my own once i understand it
	keys := make([]string, 0, len(pendingTask))
	for k := range pendingTask {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Print sorted map
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, pendingTask[k])
	}

	fmt.Printf("the pending task are %v with priority %v", taskName, priority)

	return nil
}
