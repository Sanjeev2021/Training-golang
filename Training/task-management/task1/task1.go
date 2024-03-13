package main

import (
	"errors"
	"fmt"
	"sort"
)

var tasklist = make(map[string]int)
var taskStatus = make(map[string]bool) //true : completed false - pending

func main() {
	fmt.Println(tasklist)
	var taskName string
	var priority int
	var oldtaskName string
	var newtaskName string

	for true {
		var choice int
		fmt.Println("\nChoice Menu:")
		fmt.Println("1. add task")
		fmt.Println("2. get task")
		fmt.Println("3. get All task")
		fmt.Println("4. remove task")
		fmt.Println("5. mark as completed ")
		fmt.Println("6. All pending task.")
		fmt.Println("7. All completed task.")
		fmt.Println("8. update task Name.")
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
			result, err := getTask(taskName)
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Printf("the task in the list is of name %v with priority %v", taskName, result)
		case 3:
			err := getAllTask()
			if err != nil {
				fmt.Println("Error", err)
			}
		case 4:
			fmt.Print("Enter task name: ")
			fmt.Scan(&taskName)

			fmt.Print("Enter task priority (integer): ")
			fmt.Scan(&priority)
			err := removeTask(taskName, priority)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 5:
			fmt.Println("Enter the task you want to marks as complete")
			fmt.Scan(&taskName)
			err := markComplete(taskName, priority)
			if err != nil {
				fmt.Println("Error", err)
			}
		case 6:
			err := getAllpendingTask()
			if err != nil {
				fmt.Println("Error", err)
			}
		case 7:
			err := getAllCompletedTask()
			if err != nil {
				fmt.Println("Error", err)
			}
		case 8:
			fmt.Println("Enter the old task name you want to update :")
			fmt.Scan(&oldtaskName)

			fmt.Println("Enter the new name :")
			fmt.Scan(&newtaskName)

			err := updateTaskName(oldtaskName, newtaskName)
			if err != nil {
				fmt.Println("Error", err)
			}
		}

	}
}

func addTask(taskName string, priority int) error {
	if _, ok := tasklist[taskName]; ok {
		return errors.New("Task already exists")
	}

	tasklist[taskName] = priority
	taskStatus[taskName] = false
	fmt.Printf("Added task: %v with priority %d\n", taskName, priority)

	return nil
}

func getTask(taskName string) (int, error) {
	value, err := tasklist[taskName]
	if !err {
		return value, errors.New("Task does not exists")
	}

	return value, nil
}

func getAllTask() error {
	for taskName, value := range tasklist {
		fmt.Println(taskName, value)
		fmt.Printf("the task is of priority  %v and name %v", tasklist[taskName], taskName)
	}

	return nil
}

func removeTask(taskName string, priority int) error {
	_, err := tasklist[taskName]
	if !err {
		return errors.New("Task does not exists.")
	}

	_, err = taskStatus[taskName]
	if !err {
		return errors.New("Task does not exists.")
	}

	delete(tasklist, taskName)
	delete(taskStatus, taskName)

	fmt.Printf("Removed the task successfully of name %v\n", taskName)
	fmt.Printf("The values in map %v\n and %v\n", tasklist[taskName], taskStatus[taskName])

	return nil

}

func markComplete(taskName string, priority int) error {

	_, ok := tasklist[taskName]
	if !ok {
		return errors.New("Task does not exists.")
	}

	_, ok = taskStatus[taskName]
	if !ok {
		return errors.New("Task does not exists.")
	}

	taskStatus[taskName] = true

	fmt.Println("completed task", taskStatus)
	fmt.Printf("\n The task is completed %v", taskName)

	return nil

}

func getAllpendingTask() error {

	for taskName, status := range taskStatus {
		fmt.Println(taskName, status)
		if status == false {
			fmt.Printf("\n The task is pending of name %v and priority %v ", taskName, tasklist[taskName])
		}
	}

	keys := make([]string, 0, len(tasklist))

	for k := range tasklist {
		fmt.Println("TASKLIST", k)
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return tasklist[keys[i]] <= tasklist[keys[j]]
	})

	//  fmt.Println("they key %v with priority %v", keys, tasklist[keys])

	fmt.Print("the name of the task and priority: ")
	for _, key := range keys {
		fmt.Printf("%s (%v) ", key, tasklist[key])
	}
	fmt.Println()

	return nil

}

func getAllCompletedTask() error {

	for taskName, status := range taskStatus {
		fmt.Println(taskName, status)
		if status == true {
			fmt.Printf("\n The tasklist is completed of name %v and priority %v ", taskName, tasklist[taskName])
		}
	}

	return nil
}

func updateTaskName(oldTaskname, newTaskname string) error {
	_, ok := tasklist[oldTaskname]
	if !ok {
		return errors.New("The old task name does not exist in task map!")
	}

	_, ok = taskStatus[oldTaskname]
	if !ok {
		return errors.New("The old task name does not exist in taskStatus map!")
	}

	tasklist[newTaskname] = tasklist[oldTaskname]
	taskStatus[newTaskname] = taskStatus[oldTaskname]

	delete(tasklist, oldTaskname)
	delete(taskStatus, oldTaskname)

	fmt.Printf("Updated item name from %v to %v\n", oldTaskname, newTaskname)

	return nil
}
