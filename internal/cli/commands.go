package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jerslf/todo/internal/task"
	"github.com/jerslf/todo/internal/view"
)

func commandExit(ts *task.Tasks, args ...string) error {
	fmt.Println("Closing Todo app...")
	os.Exit(0)
	return nil
}

func commandHelp(ts *task.Tasks, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Todo app!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	fmt.Println()
	return nil
}

func commandAdd(ts *task.Tasks, args ...string) error {
	title := strings.Join(args, " ")
	err := ts.Add(title)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return err
	}
	fmt.Printf("Task added successfully. Title: %s\n", title)
	return nil
}

func commandList(ts *task.Tasks, args ...string) error {
	if len(args) > 0 && args[0] == "-a" {
		view.PrintTasks(os.Stdout, ts.List(true))
	} else {
		view.PrintTasks(os.Stdout, ts.List(false))
	}
	return nil
}

func commandComplete(ts *task.Tasks, args ...string) error {
	id := args[0]
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", id)
	}

	if intID <= 0 || intID > len(ts.Items) {
		return fmt.Errorf("task ID %d does not exist", intID)
	}

	for i := range ts.Items {
		if ts.Items[i].ID == intID {
			ts.Items[i].MarkDone()
			fmt.Printf("Task ID %d marked as done.\n", intID)
			return nil
		}
	}
	return nil
}

func commandUnComplete(ts *task.Tasks, args ...string) error {
	id := args[0]
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", id)
	}

	if intID <= 0 || intID > len(ts.Items) {
		return fmt.Errorf("task ID %d does not exist", intID)
	}

	for i := range ts.Items {
		if ts.Items[i].ID == intID {
			ts.Items[i].MarkUndone()
			fmt.Printf("Task ID %d marked as not done.\n", intID)
			return nil
		}
	}
	return nil
}

func commandDelete(ts *task.Tasks, args ...string) error {
	id := args[0]
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", id)
	}

	if intID <= 0 || intID > len(ts.Items) {
		return fmt.Errorf("task ID %d does not exist", intID)
	}

	err = ts.Delete(intID)
	if err != nil {
		return fmt.Errorf("error deleting task ID %d: %v", intID, err)
	}

	fmt.Printf("Task ID %d deleted successfully.\n", intID)
	return nil
}
