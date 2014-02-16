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

func (t *tasklist) load(path string) *[]string {
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
	lines := strings.Split(text, "\n")
	return &lines
}

func (t *tasklist) add(args []string) {
	for id, text := range args {
		t.tasks = append(t.tasks, task{id, text})
	}
}

func (t *tasklist) print() {
	fmt.Println(t.title + ":")
	for _, task := range t.tasks {
		fmt.Printf("%3v %v\n", task.id, task.text)
	}
}
