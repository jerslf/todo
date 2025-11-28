package task

import (
	"strings"
	"testing"
	"time"
)

// --- Add ---

func TestTasksAdd(t *testing.T) {
	var tasks Tasks

	err := tasks.Add("Buy milk!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks.items) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks.items))
	}

	if tasks.items[0].Title != "Buy milk!" {
		t.Errorf("expected title 'Buy milk', got '%s'", tasks.items[0].Title)
	}

	if tasks.items[0].ID != 1 {
		t.Errorf("expected ID 1, got '%d'", tasks.items[0].ID)
	}

	err = tasks.Add("Buy soy!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tasks.items[1].ID != 2 {
		t.Errorf("expected ID 2, got '%d'", tasks.items[1].ID)
	}
}

func TestTasksAddEmptyTitle(t *testing.T) {
	var tasks Tasks

	err := tasks.Add("")
	if err == nil {
		t.Fatalf("expected error for empty title, got nil")
	}
}

func TestTasksAddTooLongTitle(t *testing.T) {
	var tasks Tasks
	longTitle := strings.Repeat("x", 201)

	err := tasks.Add(longTitle)
	if err == nil {
		t.Fatal("expected error for long title, got nil")
	}
}

// --- List ---

func TestList_EmptyTasks(t *testing.T) {
	var tasks Tasks

	result := tasks.List(true)
	if result != nil {
		t.Fatalf("expected nil for empty task list, got %v", result)
	}

	result = tasks.List(false)
	if result != nil {
		t.Fatalf("expected nil for empty task list, got %v", result)
	}
}

func TestList_ListAll(t *testing.T) {
	now := time.Now()
	tasks := Tasks{items: []Task{
		{ID: 1, Title: "A", Done: false, TimeCreated: now},
		{ID: 2, Title: "B", Done: true, TimeCreated: now},
	},
	}

	result := tasks.List(true)

	if len(result) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(result))
	}

	if result[0].Title != "A" || result[1].Title != "B" {
		t.Fatalf("listAll returned wrong tasks: %v", result)
	}
}

func TestList_OnlyDone(t *testing.T) {
	now := time.Now()
	t1 := Task{ID: 1, Title: "A", Done: false, TimeCreated: now}
	t2 := Task{ID: 2, Title: "B", Done: true, TimeCreated: now}
	t3 := Task{ID: 3, Title: "C", Done: false, TimeCreated: now}

	tasks := Tasks{items: []Task{t1, t2, t3}}
	result := tasks.List(false)

	if len(result) != 2 {
		t.Fatalf("expected 2 not done task, got %d", len(result))
	}

	if result[0].ID != 1 || result[1].ID != 3 {
		t.Fatalf("list(false) returned wrong tasks: %v", result)
	}
}

// --- MarkDone ---

func TestMarkDone(t *testing.T) {
	task := Task{
		ID:          1,
		Title:       "Test",
		Done:        false,
		TimeCreated: time.Now(),
		TimeDone:    nil,
	}

	task.MarkDone()

	if !task.Done {
		t.Fatalf("expected task to be marked done")
	}

	if task.TimeDone == nil {
		t.Fatalf("expected TimeDone to be set, got nil")
	}

	// Check that TimeDone is recent (within the last second)
	if time.Since(*task.TimeDone) > time.Second {
		t.Fatalf("expected TimeDone to be recent, got %v", *task.TimeDone)
	}
}

func TestMarkDoneWhenAlreadyDone(t *testing.T) {
	initialTime := time.Now().Add(-1 * time.Hour)
	task := Task{
		ID:          1,
		Title:       "Test",
		Done:        true,
		TimeCreated: time.Now(),
		TimeDone:    &initialTime,
	}

	task.MarkDone()

	if !task.Done {
		t.Fatalf("expected task to stay done")
	}

	if *task.TimeDone != initialTime {
		t.Fatalf("expected TimeDone not to change")
	}
}

// --- MarkUndone ---

func TestMarkUndone(t *testing.T) {
	initialTime := time.Now()

	task := Task{
		ID:          1,
		Title:       "Test",
		Done:        true,
		TimeCreated: time.Now(),
		TimeDone:    &initialTime,
	}

	task.MarkUndone()

	if task.Done {
		t.Fatalf("expected task to be marked undone")
	}

	if task.TimeDone != nil {
		t.Fatalf("expected TimeDone to be nil after MarkUndone, got %v", *task.TimeDone)
	}
}

func TestMarkUndoneWhenAlreadyUndone(t *testing.T) {
	task := Task{
		ID:          1,
		Title:       "Test",
		Done:        false,
		TimeCreated: time.Now(),
		TimeDone:    nil,
	}

	task.MarkUndone()

	if task.Done {
		t.Fatalf("expected task to remain undone")
	}

	if task.TimeDone != nil {
		t.Fatalf("expected TimeDone to stay nil")
	}
}

func TestDeleteExisting(t *testing.T) {
	now := time.Now()
	tasks := Tasks{items: []Task{
		{ID: 1, Title: "A", Done: false, TimeCreated: now},
		{ID: 2, Title: "B", Done: true, TimeCreated: now},
		{ID: 3, Title: "C", Done: false, TimeCreated: now},
	}}

	err := tasks.Delete(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks.items) != 2 {
		t.Fatalf("expected 2 tasks after deletion, got %d", len(tasks.items))
	}
}
