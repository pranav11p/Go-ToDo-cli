package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func OpenFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return f, fmt.Errorf("\nerror opening file %v", err)
	}
	return f, nil
}

func ParseFileToTodo(path string) ([]Todo, error) {
	var todos []Todo

	// Read the file contents
	data, err := os.ReadFile(path)
	if err != nil {
		return todos, fmt.Errorf("\nError reading file %v", err)
	}

	// Cast it to the slice
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return todos, fmt.Errorf("\nError decoding the data from the file %v", err)
	}

	fmt.Printf("The existing data in the file is: \n")
	fmt.Println(todos)

	return todos, nil
}

func WriteTodosToFile(f *os.File, todos []Todo) error {
	// Marshal(Convert) the Todo objects into a JSON string
	todosJSONBytes, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		return fmt.Errorf("\nError marshaling Todo %v", err)
	}
	todosJSONContents := string(todosJSONBytes)

	// Truncate the file before writing
	err = f.Truncate(0)
	if err != nil {
		return fmt.Errorf("\nError truncating Todo file %v", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("\nError seeking Todo file %v", err)
	}

	_, err = f.WriteString(todosJSONContents)
	if err != nil {
		return fmt.Errorf("\nError writing to file %v", err)
	}
	return nil
}
