package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vishnu-g-k/to-do-list-application/model"
)

var todo []model.Todo

func addTask(task string) {
	newtask := model.Todo{
		Task:      task,
		Completed: false,
	}
	todo = append(todo, newtask)
	fmt.Println("Tasks added Successfully")
}

func listTasks() {
	for i, task := range todo {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("\n\t%d. %s \t [%s]", i+1, task.Task, status)
	}
}

func markAsCompleted(index int) {
	if index > 0 && index <= len(todo) {
		todo[index-1].Completed = true
	} else {
		fmt.Println("Invalid input...!!!")
	}

}

func removeTask(index int) {
	if index > 0 && index <= len(todo) {
		todo = append(todo[:index-1], todo[index:]...)
		fmt.Println("Task removed successfully")
	} else {
		fmt.Println("Invalid input...!!!")
	}
}

func TodoListHandler() {
	for {
		fmt.Println("\n----------------\nTODO List\n\t1. Add a task\n\t2. List all tasks\n\t3. Mark task as completed\n\t4. Deleted from list\n\t5. Exit")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter your choice: ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input..!!")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Enter the task: ")
			scanner.Scan()
			task := scanner.Text()
			addTask(task)
		case 2:
			listTasks()
		case 3:
			fmt.Println("Enter the index of task that needs to be marked as completed: ")
			scanner.Scan()
			value := scanner.Text()
			index, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Invalid Choice..!!!")
			}
			markAsCompleted(index)
		case 4:
			fmt.Println("Enter the index of task that needs to be removed: ")
			scanner.Scan()
			value := scanner.Text()
			index, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Invalid Choice..!!!")
			}
			removeTask(index)
		case 5:
			fmt.Println("Exiting...!!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Input. Try again...!!")
		}
	}

}
