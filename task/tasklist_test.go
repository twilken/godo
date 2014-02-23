package task

import (
	"fmt"
	"os"
	"testing"
)

// createTasklist returns a Tasklist with numOfTasks tasks.
func createTasklist(numOfTasks int) *Tasklist {
	tasks := []Task{}
	for i := 0; i < numOfTasks; i++ {
		text := fmt.Sprint("Task No", i)
		tasks = append(tasks, Task{text})
	}
	return &Tasklist{Title: "Tasks", Tasks: tasks}
}

func TestLoadNonExistingFile(t *testing.T) {
	path := ".godo_does_not_exist"
	list := createTasklist(0)

	// Delete file after test
	defer func() {
		os.Remove(path)
	}()

	// Return an error if Load function does not automatically create the missing file
	if err := list.Load(path); err != nil {
		t.Error(err)
	}
}

func TestAdd(t *testing.T) {
	list := createTasklist(0)
	list.Add([]string{"T1", "T2", "T3"})
	if list.Len() != 3 {
		t.Error("Expected 3, got ", list.Len())
	}
}

func TestDel(t *testing.T) {
	list := createTasklist(5)
	list.Del([]int{0, 2, 4})
	numOfTasks := list.Len()
	if numOfTasks != 2 {
		t.Error("Expected 2, got ", numOfTasks)
	}
}

func TestDelOutOfBounds(t *testing.T) {
	list := createTasklist(5)
	// catch out of bounds panic
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	list.Del([]int{-1, 100000})
}

func TestLen(t *testing.T) {
	list := createTasklist(0)
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
