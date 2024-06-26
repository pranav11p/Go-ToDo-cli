package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func InputLoop(todos []Todo, file *os.File) ([]Todo, error) {
loop:
	for {
		go WriteTodosToFile(file, todos)
		fmt.Println("")
		fmt.Println("Please choose an option:")
		fmt.Println("1. View todos")
		fmt.Println("2. Create a new Todo")
		fmt.Println("3. Toggle todo as done")
		fmt.Println("4. Delete a Todo")
		fmt.Println("9. Save to file")
		fmt.Println("0. Exit")
		fmt.Println("Enter your Choice: ")

		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("E#1VUOMM - %v\n", err)
		}

		switch input {
		case "1":
			{
				ViewTodos(todos)
			}
		case "2":
			{
				t, err := CreateATodo()
				if err != nil {
					fmt.Println(err)
				} else {
					todos = append(todos, t)
				}
			}
		case "3":
			{
				ToggleTodoDone(todos)
			}
		case "4":
			{
				todos = DeleteTodo(todos)
			}
		case "9":
			{
				err = WriteTodosToFile(file, todos)
				if err != nil {
					println(err)
				} else {
					println("Wrote to the file!!")
				}
			}
		case "0":
			{
				break loop
			}
		default:
			{
				fmt.Println("Invalid Options is selected")
			}
		}
	}
	return todos, nil
}

func ReadATodo() Todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Creating new ToDo")

	fmt.Printf("Enter the Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Enter the description: ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return Todo{
		Title:     strings.Trim(title, "\n"),
		Desc:      strings.Trim(desc, "\n"),
		CreatedAt: time.Now(),
	}
}

func ReadTodosInLoop() []Todo {
	todos := make([]Todo, 0)
	for {
		fmt.Printf("\nPlease choose an option:\n")
		fmt.Printf("1. Create a new Todo:\n")
		fmt.Printf("0. Exit:\n")
		fmt.Printf("Enter your Choice: ")

		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("E#1VUOMM - %v\n", err)
		}

		if input == "0" {
			break
		} else if input == "1" {
			t := ReadATodo()
			// t.Print()
			todos = append(todos, t)
		} else {
			fmt.Printf("Invalid option selected, try again")
		}
	}
	return todos
}
