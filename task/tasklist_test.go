package task

import (
	"fmt"
	"os"
	"testing"
)

// createTasklist returns a Tasklist with num tasks.
func createTasklist(num int) *Tasklist {
	tasks := []Task{}
	for i := 0; i < num; i++ {
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

func TestLoadExisingFile(t *testing.T) {
	path := ".godo_file_exists"
	list := createTasklist(0)

	// Delete file after test
	defer func() {
		os.Remove(path)
	}()

	// Setup a file with some tasks
	file, err := os.Create(path)
	if err != nil {
		t.Error("Test setup is not able to create file. ", err)
	}
	defer file.Close()
	if _, err := file.WriteString("Task1\nTask2\nTask3\n"); err != nil {
		t.Error("Test setup can not write to file. ", err)
	}

	// Actual test
	if err := list.Load(path); err != nil {
		t.Error("Load returned error: ", err)
	}
	numOfTasks := list.Len()
	if numOfTasks != 3 {
		t.Error("Expected 3, got ", numOfTasks)
	}
	if list.Tasks[0].Text != "Task1" ||
		list.Tasks[1].Text != "Task2" ||
		list.Tasks[2].Text != "Task3" {
		t.Error("Task content not loaded properly")
	}
}

func TestSaveNotPreviouslyExisingFile(t *testing.T) {
	path := ".godo_file_does_not_exist"
	list := createTasklist(3)

	// Delete file after test
	defer func() {
		os.Remove(path)
	}()

	list.Save(path)
	loaded := createTasklist(0)
	loaded.Load(path)
	numOfTasks := loaded.Len()
	if numOfTasks != 3 {
		t.Error("Expected 3, got ", numOfTasks)
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

	list = createTasklist(4)
	list.Del([]int{3, 0, 1, 2})
	numOfTasks = list.Len()
	if numOfTasks != 0 {
		t.Error("Expected 0, got ", numOfTasks)
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
