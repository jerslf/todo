package task

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Done        bool
	TimeCreated time.Time
	TimeDone    *time.Time
}

type Tasks struct {
	Items  []Task
	NextID int
}

func (ts *Tasks) Add(title string) error {
	// Check empty string
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}
	// Check length
	if len(title) > 100 {
		return errors.New("title is too long")
	}
	// Normalise
	title = strings.TrimSpace(title)

	// Create and append task
	id := ts.NextID + 1
	ts.NextID = id
	task := Task{
		ID:          id,
		Title:       title,
		Done:        false,
		TimeCreated: time.Now(),
		TimeDone:    nil,
	}
	ts.Items = append(ts.Items, task)
	return nil
}

func (ts *Tasks) List(listAll bool) []Task {
	if len(ts.Items) == 0 {
		return nil
	}

	if listAll {
		return ts.Items
	}

	// only not done tasks
	var undone []Task
	for _, task := range ts.Items {
		if !task.Done {
			undone = append(undone, task)
		}
	}
	return undone
}

func (t *Task) MarkDone() {
	if t.Done {
		return
	}
	t.Done = true
	now := time.Now()
	t.TimeDone = &now
}

func (t *Task) MarkUndone() {
	t.Done = false
	t.TimeDone = nil
}

func (ts *Tasks) Delete(id int) error {
	for i, task := range ts.Items {
		if task.ID == id {
			ts.Items = slices.Delete(ts.Items, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
