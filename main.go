package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	Name        string
	Description string
	Done        bool
}

var allTasks []task

func main() {
	menu()
}

func menu() {
	for {
		var option int
		fmt.Printf("1 - Add Task\n2 - List tasks\n3 - Finish a task\n4 - delete a  task\n5 - Exit\n")
		fmt.Println("What you want to do?")
		fmt.Scan(&option)
		switch option {
		case 1:
			allTasks = append(allTasks, addTask())
		case 2:
			list()
		case 3:
			fmt.Println("you typed 3")
		case 4:
			fmt.Println("you typed 4")
		case 5:
			exit()
		default:
			fmt.Println("you typed it wrong")
		}
	}
}

func addTask() task {
	task := task{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Task Name: ")
	scanner.Scan()
	name := scanner.Text()
	task.Name = name
	fmt.Println("Task description: ")
	scanner.Scan()
	description := scanner.Text()
	task.Description = description
	return task
}

func list() {
	fmt.Println(allTasks)
}

func exit() {
	os.Exit(0)
}
