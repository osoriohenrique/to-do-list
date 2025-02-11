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
		fmt.Printf("1 - Adicionar Tarefa\n2 - Listar tarefa\n3 - Finalizar tarefa\n4 - Excluir tarefa\n5 - Sair\n")
		fmt.Println("O que vc deseja fazer? ")
		fmt.Scan(&option)
		switch option {
		case 1:
			allTasks = append(allTasks, addTask())
		case 2:
			list()
		case 3:
			fmt.Println("Digitou 3")
		default:
			fmt.Println("Digitou errado")
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
