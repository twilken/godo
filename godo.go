package main

import (
	"flag"
	"fmt"
	"github.com/tenpeoplemeet/godo/task"
	"log"
	"os"
	"strconv"
)

const dataPath string = ".godo" // TODO Use $HOME and use system variable instead
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
	tasks.Load(dataPath)
	switch args[0] {
	case "add", "a":
		tasks.Add(args[1:])
		tasks.List()
	case "list", "l":
		tasks.List()
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
		tasks.List()
	}
	tasks.Save(dataPath)
}
