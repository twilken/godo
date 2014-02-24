package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// Tasklist is a list of Tasks with some meta information.
type Tasklist struct {
	Title string
	Tasks []Task
}

// Load loads the Tasklist t with the Tasks contained in the file at path.
// If the file does not exist, it will be created. In case the file can not
// be created or can not be read properly, an error is returned.
func (t *Tasklist) Load(path string) error {
	if !fileExists(path) {
		if _, err := os.Create(path); err != nil {
			return err
		}
	}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(raw)
	lines := strings.Split(text, "\n")
	t.Add(lines[:len(lines)-1])
	return nil
}

// fileExists returns true if the file specified by path exists. Otherwise
// it returns false.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Save saves all tasks in t to the file at path. Tasks are stored in plain text
// and each tasks is separated by a new line.
func (t *Tasklist) Save(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, task := range t.Tasks {
		_, err := file.WriteString(task.Text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Add creates a task for each string in texts and adds them to t.
func (t *Tasklist) Add(texts []string) {
	for _, text := range texts {
		trimmed := strings.Replace(text, "\n", " ", -1)
		t.Tasks = append(t.Tasks, Task{trimmed})
	}
}

// Del deletes every task in t with an index contained in ids. If the index is out of
// bounds, it will be ignored.
func (t *Tasklist) Del(ids []int) {
	len := t.Len()
	sort.Sort(sort.IntSlice(ids))
	for i, id := range ids {
		if id >= 0 && id < len {
			t.Tasks = append(t.Tasks[:id-i], t.Tasks[id-i+1:]...)
		}
	}
}

// Len returns the number of Tasks contained in t.
func (t *Tasklist) Len() int {
	return len(t.Tasks)
}

// String returns a pretty representation of t.
func (t Tasklist) String() string {
	s := fmt.Sprint(t.Title, ":\n")
	for i, task := range t.Tasks {
		s += fmt.Sprintf("%3v %v\n", i, task.Text)
	}
	return s
}
