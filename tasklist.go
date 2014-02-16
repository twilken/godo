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
	t.add(lines)
}

func (t *tasklist) add(args []string) {
	n := len(t.tasks)
	for i, text := range args {
		t.tasks = append(t.tasks, task{n + i, text})
	}
}

func (t *tasklist) print() {
	fmt.Println(t.title + ":")
	for _, task := range t.tasks {
		fmt.Printf("%3v %v\n", task.id, task.text)
	}
}
