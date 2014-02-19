package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Tasklist struct {
	Title string
	Tasks []Task
}

func (t *Tasklist) Load(path string) {
	if _, err := os.Stat(path); err != nil {
		if _, err := os.Create(path); err != nil {
			log.Fatal(err, "\nCould not create file at "+path)
		}
	}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err, "\nCould not read data from "+path)
	}
	text := string(raw)
	lines := strings.Split(text, "\n") // TODO Check for other new line chars
	t.Add(lines[:len(lines)-1])
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

func (t *Tasklist) Add(args []string) {
	for _, text := range args {
		t.Tasks = append(t.Tasks, Task{text})
	}
}

func (t *Tasklist) List() {
	fmt.Println(t.Title + ":")
	for i, task := range t.Tasks {
		fmt.Printf("%3v %v\n", i, task.Text)
	}
}

func (t *Tasklist) Del(ids []int) {
	for i, id := range ids {
		t.Tasks = append(t.Tasks[:id-i], t.Tasks[id-i+1:]...)
	}
}
