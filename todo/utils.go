package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ViewTodos(todos []Todo) {
	for i, t := range todos {
		fmt.Printf("ID: %d\n", i+1)
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
	reader := bufio.NewReader(os.Stdin)

start:
	{
		fmt.Printf("\n Enter -1 to exit")
		fmt.Printf("\n Please enter the Todo ID: \t")
		id_str, err := reader.ReadString('\n')
		id_str = strings.Trim(id_str, "\n")
		if err != nil {
			fmt.Printf("Error reading Todo ID: %v\n", err)
			return
		}

		if id_str == "-1" {
			return
		}
		id, err := strconv.Atoi(id_str)

		if err != nil || id >= len(todos) || id < 0 {
			fmt.Printf("Invalid Todo ID %d: %v\n", id, err)
			goto start
		}

		t := todos[id-1]
		t.Done = true
		todos[id-1] = t
	}
	// return todos
}
