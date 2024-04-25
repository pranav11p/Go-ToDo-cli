package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pranavpatel3012/go-todo-cli/todo"
)

func main() {
	fmt.Printf("Hello\n")

	t := todo.Todo{
		Title:     "Title of Todo",
		Desc:      "Description of Todo",
		Done:      false,
		CreatedAt: time.Now(),
	}
	t.Print()
	fmt.Println(t.Validate())

	err := t.ValidateWithError()
	if err != nil {
		fmt.Printf("\nTodo is not valid. Error: %v\n", err)
	} else {
		fmt.Printf("\nTodo is valid")
	}

	todoBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("\nSomething went wrong while encoding to JSON: %v\n", err)
	}
	// fmt.Printf("\n JSON Bytes: %v\n", todoBytes)
	todoJSONString := string(todoBytes)
	fmt.Printf("\n JSON String: %v\n", todoJSONString)
	// fmt.Printf("\n JSON String to Bytes: %v\n", []byte(todoJSONString))

	// t.Print()

	// Struct from JSON
	jsonString := `
	{
		"Title": "Title of the todo",
		"desc": "Desc of the todo"
	}
	`

	newTodo := todo.Todo{}
	err = json.Unmarshal([]byte(jsonString), &newTodo)
	fmt.Printf("\nJSON string: %v\n Converting error: %v\n", jsonString, err)
	newTodo.Print()

}
