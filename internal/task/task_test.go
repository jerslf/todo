package task

import (
	"strings"
	"testing"
	"time"
)

func TestTasksAdd(t *testing.T) {
	var tasks Tasks

	err := tasks.Add("Buy milk!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Buy milk!" {
		t.Errorf("expected title 'Buy milk', got '%s'", tasks[0].Title)
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
	tasks := Tasks{
		{Title: "A", Done: false, TimeCreated: now},
		{Title: "B", Done: true, TimeCreated: now},
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
	t1 := Task{Title: "A", Done: false, TimeCreated: now}
	t2 := Task{Title: "B", Done: true, TimeCreated: now}
	t3 := Task{Title: "C", Done: false, TimeCreated: now}

	tasks := Tasks{t1, t2, t3}

	result := tasks.List(false)

	if len(result) != 2 {
		t.Fatalf("expected 2 not done task, got %d", len(result))
	}

	if result[0].Title != "A" || result[1].Title != "C" {
		t.Fatalf("list(false) returned wrong tasks: %v", result)
	}
}
