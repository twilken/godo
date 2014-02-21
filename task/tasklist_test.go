package task

import (
	"testing"
)

func TestDel(t *testing.T) {
	/*
		tasks := []Task{
			Task{"Task No. 1"},
			Task{"Task No. 2"},
			Task{"Task No. 3"},
			Task{"Task No. 4"},
			Task{"Task No. 5"},
		} */
	var tasks *Tasklist
	tasks = &Tasklist{Title: "Tasks"}
	tasks.Add([]string{"Task1", "Task2", "Task3", "Task4", "Task5"})
	tasks.Del([]int{0, 2, 4})
	/*
		list := Tasklist{}
		list.Add([]string{"Task1", "Task2", "Task3", "Task4", "Task5"})
		list.Del([]int{1, 3, 5})
		if list.Len() != 2 {
			t.Fail()
		}
	*/
	if tasks.Len() != 2 {
		t.Fail()
	}
}
