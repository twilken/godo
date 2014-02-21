package task

import (
	"testing"
)

func TestDel(t *testing.T) {
	ts := []Task{
		Task{"Task No. 1"},
		Task{"Task No. 2"},
		Task{"Task No. 3"},
		Task{"Task No. 4"},
		Task{"Task No. 5"},
	}
	list := Tasklist{Title: "Tasks", Tasks: ts}
	list.Del([]int{0, 2, 4})
	numOfTasks := list.Len()
	if numOfTasks != 2 {
		t.Error("Expected 2, got ", numOfTasks)
	}
}
