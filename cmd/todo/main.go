package main

import (
	"github.com/jerslf/todo/internal/cli"
	"github.com/jerslf/todo/internal/task"
)

func main() {
	tasks := &task.Tasks{}

	cli.StartRepl(tasks)
}
