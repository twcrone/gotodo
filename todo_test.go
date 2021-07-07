package todo_test

import (
	"pragprog.com/rggo/interacting/todo"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, go %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	err := l.Complete(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}
