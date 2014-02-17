package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type tasklist struct {
	title string
	tasks []task
}

func (t *tasklist) load(path string) {
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
	t.add(lines[:len(lines)-1])
}

func (t *tasklist) save(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, task := range t.tasks {
		_, err := file.WriteString(task.text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (t *tasklist) add(args []string) {
	n := len(t.tasks)
	for i, text := range args {
		t.tasks = append(t.tasks, task{n + i, text})
	}
}

func (t *tasklist) list() {
	fmt.Println(t.title + ":")
	for _, task := range t.tasks {
		fmt.Printf("%3v %v\n", task.id, task.text)
	}
}

func (t *tasklist) del(ids []int) {
	for i, id := range ids {
		if intInSlice(id, ids) {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
		}
	}
}

func intInSlice(a int, slice []int) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
