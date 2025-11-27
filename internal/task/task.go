package task

import (
	"errors"
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

type Tasks []Task

func (t *Tasks) Add(title string) error {
	// Check empty string
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}
	// Check length
	if len(title) > 200 {
		return errors.New("title is too long")
	}
	// Normalise
	title = strings.TrimSpace(title)

	// Create and append task
	id := len(*t) + 1
	task := Task{
		ID:          id,
		Title:       title,
		Done:        false,
		TimeCreated: time.Now(),
		TimeDone:    nil,
	}
	*t = append(*t, task)
	return nil
}

func (t *Tasks) List(listAll bool) []Task {
	if len(*t) == 0 {
		return nil
	}

	if listAll {
		return *t
	}

	// only not done tasks
	var undone []Task
	for _, task := range *t {
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
