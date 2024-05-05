package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (todo *Todo) Print() {
	fmt.Println("Title: ", todo.Title)
	fmt.Println("Desc: ", todo.Desc)
	fmt.Println("Done: ", todo.Done)
	fmt.Println("CreatedAt: ", todo.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("------------------")
	// fmt.Println("UpdatedAt: ", todo.UpdatedAt)
}

func (todo *Todo) ValidateWithError() error {
	if len(todo.Title) == 0 {
		return fmt.Errorf("title is empty")
	}

	if len(todo.Desc) > 100 {
		return fmt.Errorf("desc is too long")
	}
	return nil
}

func (todo *Todo) Validate() bool {
	if len(todo.Title) == 0 {
		return false
	}

	if len(todo.Desc) > 1000 {
		return false
	}
	return true
}
