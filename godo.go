package main

import (
	"flag"
	"fmt"
	"github.com/tenpeoplemeet/godo/task"
	"log"
	"os/user"
	"strconv"
)

const usage string = `
godo [subcommand] [arguments to subcommand]
subcommands:
	add, a		Add one or more tasks
	del, d		Delete one or more tasks by their number
	list, l		Show a list of all tasks
	help, h		Show help text
`

const saveFileName string = ".godo"

var tasks *task.Tasklist

func main() {
	checkNumOfArgs()
	path := getSaveFilePath()
	tasks = &task.Tasklist{Title: "Tasks"}
	if err := tasks.Load(path); err != nil {
		log.Fatal(err)
	}
	processSubcommands()
	tasks.Save(path)
}

func processSubcommands() {
	args := flag.Args()
	switch args[0] {
	case "add", "a":
		add(args[1:])
		list()
	case "del", "d":
		del(args[1:])
		list()
	case "list", "l":
		list()
	case "help", "h":
		fmt.Print(usage)
	}
}

func checkNumOfArgs() []string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal(usage)
	}
	return args
}

func getSaveFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/" + saveFileName
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
