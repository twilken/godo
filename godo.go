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
var tasks []task.Task = make([]task.Task, 0, 50)

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
	tasks := task.Tasklist{Title: "Tasks"}
	tasks.Load(path)
	switch args[0] {
	case "add", "a":
		tasks.Add(args[1:])
		fmt.Print(tasks)
	case "list", "l":
		fmt.Print(tasks)
	case "del", "d":
		delArgs := args[1:]
		ids := make([]int, len(delArgs))
		for i := range ids {
			id, err := strconv.Atoi(delArgs[i])
			if err != nil {
				log.Fatal(err)
			}
			ids[i] = id
		}
		tasks.Del(ids)
		fmt.Print(tasks)
	}
	tasks.Save(path)
}
