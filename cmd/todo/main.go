package main

import (
	"fmt"
	"os"

	"github.com/jerslf/todo/internal/task"
	"github.com/jerslf/todo/internal/view"
)

func main() {
	var tasks task.Tasks

	err := tasks.Add("Buy milk!")
	if err != nil {
		fmt.Printf("Error adding task: %v\n", err)
	}

	err = tasks.Add("Buy soy!")
	if err != nil {
		fmt.Printf("Error adding task: %v\n", err)
	}

	view.PrintTasks(os.Stdout, tasks)
}
