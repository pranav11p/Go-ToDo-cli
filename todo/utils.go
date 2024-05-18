package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readTodoId() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n Enter -1 to exit")
	fmt.Printf("\n Please enter the Todo ID: \t")
	id_str, err := reader.ReadString('\n')
	id_str = strings.Trim(id_str, "\n")
	if err != nil {
		return -1, fmt.Errorf("error reading Todo ID: %v", err)
	}

	if id_str == "-1" {
		return -1, nil
	}

	id, err := strconv.Atoi(id_str)
	return id, err
}

func ViewTodos(todos []Todo) {
	for i, t := range todos {
		fmt.Printf("ID: %d\n", i)
		t.Print()
	}
}

func CreateATodo() (Todo, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return Todo{}, fmt.Errorf("error: %v", err)
	}

	fmt.Printf("Enter the description: ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		return Todo{}, fmt.Errorf("error: %v", err)
	}

	return Todo{
		Title:     strings.Trim(title, "\n"),
		Desc:      strings.Trim(desc, "\n"),
		CreatedAt: time.Now(),
	}, nil
}

func MarkTodoDone(todos []Todo) {

start:
	{
		id, err := readTodoId()
		if id == -1 {
			return
		}

		if err != nil || id >= len(todos) || id < 0 {
			fmt.Printf("Invalid Todo ID %d: %v\n", id, err)
			goto start
		}

		t := todos[id-1]
		t.Done = true
		todos[id-1] = t
	}
}

func DeleteTodo(todos []Todo) []Todo {

start:
	{
		id, err := readTodoId()
		if id == -1 {
			return todos
		}

		if err != nil || id >= len(todos) || id < 0 {
			fmt.Printf("Invalid Todo ID %d: %v\n", id, err)
			goto start
		}
		return append(todos[:id], todos[id+1:]...)
	}

}
