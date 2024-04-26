package main

import (
	"fmt"

	"github.com/pranavpatel3012/go-todo-cli/todo"
)

const TODOFile = "todo/todos.json"

func main() {
	fmt.Printf("Welcome to ToDo cli made in GoLang!!\n")

	// Read the file and get the todos
	f, err := todo.OpenFile(TODOFile)
	if err != nil {
		println(err)
		return
	}
	defer f.Close()

	todos, err := todo.ParseFileToTodo(TODOFile)
	if err != nil {
		println(err)
		return
	}

	todos, err = todo.InputLoop(todos)
	if err != nil {
		println(err)
	}

	err = todo.WriteTodosToFile(f, todos)
	if err != nil {
		println(err)
	} else {
		println("Wrote to the file!!")
	}

}
