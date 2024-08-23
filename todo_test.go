package todo_test

import (
	"os"
	"testing"

	todo "github.com/hazadus/go-todo"
)

func TestAdd(t *testing.T) {
	list := todo.List{}
	taskName := "New to-do"

	list.Add(taskName)

	if list[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, list[0].Task)
	}

	if list[0].IsDone {
		t.Errorf("Task should not be completed by default.")
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}
	list.Add("New to-do")

	list.Complete(1)

	if !list[0].IsDone {
		t.Errorf("Task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	list := todo.List{}
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
	for _, task := range tasks {
		list.Add(task)
	}

	list.Delete(2)

	expLen := 2
	if len(list) != expLen {
		t.Errorf("Expected list length %d, got %d instead.", expLen, len(list))
	}
	if list[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], list[1].Task)
	}
}

func TestSaveLoad(t *testing.T) {
	list1 := todo.List{}
	list2 := todo.List{}

	taskName := "New to-do"
	list1.Add(taskName)

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s.", err)
	}
	defer os.Remove(tempFile.Name())

	if err := list1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s.", err)
	}
	if err := list2.Load(tempFile.Name()); err != nil {
		t.Fatalf("Error loading list from file: %s.", err)
	}
	if list1[0].Task != list2[0].Task {
		t.Errorf("Task %q should be equal to %q.", list1[0].Task, list1[1].Task)
	}
}
