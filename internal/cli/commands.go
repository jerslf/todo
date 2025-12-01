package cli

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jerslf/todo/internal/task"
	"github.com/jerslf/todo/internal/view"
)

func commandExit(db *task.DB, args ...string) error {
	fmt.Println("Closing Todo app...")
	os.Exit(0)
	return nil
}

func commandHelp(db *task.DB, args ...string) error {
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

func commandAdd(db *task.DB, args ...string) error {
	title := strings.Join(args, " ")
	now := time.Now()

	result, err := db.Conn.Exec("INSERT INTO tasks(title, done, time_created) VALUES (?, ?, ?)", title, false, now)
	if err != nil {
		return fmt.Errorf("insert task: %w", err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("Task added successfully. ID: %d, Title: %s\n", id, title)
	return nil
}

func commandList(db *task.DB, args ...string) error {
	listAll := len(args) > 0 && args[0] == "-a"

	query := "SELECT id, title, done, time_created, time_done FROM tasks"
	if !listAll {
		query += " WHERE done = 0"
	}

	rows, err := db.Conn.Query(query)
	if err != nil {
		return fmt.Errorf("query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		var doneTime sql.NullTime
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.TimeCreated, &doneTime); err != nil {
			return fmt.Errorf("scan row: %w", err)
		}
		if doneTime.Valid {
			t.TimeDone = &doneTime.Time
		}
		tasks = append(tasks, t)
	}

	view.PrintTasks(os.Stdout, tasks)
	return nil
}

func commandComplete(db *task.DB, args ...string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	now := time.Now()
	_, err = db.Conn.Exec("UPDATE tasks SET done = 1, time_done = ? WHERE id = ?", now, id)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	fmt.Printf("Task ID %d marked as done.\n", id)
	return nil
}

func commandUnComplete(db *task.DB, args ...string) error {
	id, _ := strconv.Atoi(args[0])
	_, err := db.Conn.Exec("UPDATE tasks SET done = 0, time_done = NULL WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	fmt.Printf("Task ID %d marked as not done.\n", id)
	return nil
}

func commandDelete(db *task.DB, args ...string) error {
	id, _ := strconv.Atoi(args[0])
	_, err := db.Conn.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	fmt.Printf("Task ID %d deleted successfully.\n", id)
	return nil
}
