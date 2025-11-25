package task

import (
	"strings"
	"testing"
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
