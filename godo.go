package main

import (
	"flag"
	"fmt"
	"github.com/tenpeoplemeet/godo/task"
	"log"
	"os"
	"strconv"
)

const path string = ".godo" // TODO Use $HOME and use system variable instead
var tasks *task.Tasklist

func usage() {
	fmt.Println("godo [add|list|del] [subcommands]")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
		os.Exit(0)
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
