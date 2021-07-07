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
