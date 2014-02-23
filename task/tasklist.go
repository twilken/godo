package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Tasklist struct {
	Title string
	Tasks []Task
}

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
	lines := strings.Split(text, "\n") // TODO Check for other new line chars
	t.Add(lines[:len(lines)-1])
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

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

func (t *Tasklist) Add(texts []string) {
	for _, text := range texts {
		trimmed := strings.Replace(text, "\n", " ", -1)
		t.Tasks = append(t.Tasks, Task{trimmed})
	}
}

func (t *Tasklist) Del(ids []int) {
	len := t.Len()
	sort.Sort(sort.IntSlice(ids))
	for i, id := range ids {
		if id >= 0 && id < len {
			t.Tasks = append(t.Tasks[:id-i], t.Tasks[id-i+1:]...)
		}
	}
}

func (t *Tasklist) Len() int {
	return len(t.Tasks)
}

func (t Tasklist) String() string {
	s := fmt.Sprint(t.Title, ":\n")
	for i, task := range t.Tasks {
		s += fmt.Sprintf("%3v %v\n", i, task.Text)
	}
	return s
}
