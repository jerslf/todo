package main

import (
	"log"

	"github.com/jerslf/todo/internal/cli"
	"github.com/jerslf/todo/internal/task"
)

func main() {
	db, err := task.NewDB("tasks.db")
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	cli.StartRepl(db)
}
