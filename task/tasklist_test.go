package task

import (
	"testing"
)

func getFilledTasklist() *Tasklist {
	tasks := []Task{
		Task{"Task No. 1"},
		Task{"Task No. 2"},
		Task{"Task No. 3"},
		Task{"Task No. 4"},
		Task{"Task No. 5"},
	}
	return &Tasklist{Title: "Tasks", Tasks: tasks}
}

func TestDel(t *testing.T) {
	list := getFilledTasklist()
	list.Del([]int{0, 2, 4})
	numOfTasks := list.Len()
	if numOfTasks != 2 {
		t.Error("Expected 2, got ", numOfTasks)
	}
}

func TestDelOutOfBounds(t *testing.T) {
	list := getFilledTasklist()
	// catch out of bounds panic
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	list.Del([]int{-1, 100000})
}

func TestLen(t *testing.T) {
	list := &Tasklist{Title: "Tasks"}
	if list.Len() != 0 {
		t.Error("Expected 0, got ", list.Len())
	}
	list.Add([]string{"T1", "T2", "T3"})
	if list.Len() != 3 {
		t.Error("Expected 3, got ", list.Len())
	}
	list.Del([]int{0, 1})
	if list.Len() != 1 {
		t.Error("Expected 1, got ", list.Len())
	}
	list.Del([]int{0})
	if list.Len() != 0 {
		t.Error("Expected 0, got ", list.Len())
	}
}
