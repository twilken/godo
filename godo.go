package main

import (
	"flag"
	"fmt"
	"github.com/tenpeoplemeet/godo/task"
	"log"
	"strconv"
)

const usage string = `
godo [subcommand] [arguments to subcommand]
subcommands:
	add, a		add one or more tasks
	del, d		delete one or more tasks by their number
	list, l		show a list of all tasks
`
const path string = ".godo" // TODO Use $HOME and use system variable instead
var tasks *task.Tasklist

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal(usage)
	}
	tasks = &task.Tasklist{Title: "Tasks"}
	tasks.Load(path)
	switch args[0] {
	case "add", "a":
		add(args[1:])
		list()
	case "list", "l":
		list()
	case "del", "d":
		del(args[1:])
		list()
	}
	tasks.Save(path)
}

func add(args []string) {
	tasks.Add(args)
}

func list() {
	fmt.Print(tasks)
}

func del(args []string) {
	ids := make([]int, len(args))
	for i := range ids {
		id, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatal(err)
		}
		ids[i] = id
	}
	tasks.Del(ids)
}
