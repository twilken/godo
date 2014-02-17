package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const dataPath string = ".godo" // TODO Use $HOME and use system variable instead
var tasks []task = make([]task, 0, 50)

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
	tasks := tasklist{title: "Tasks"}
	tasks.load(dataPath)
	switch args[0] {
	case "add":
		tasks.add(args[1:])
		tasks.list()
	case "list":
		tasks.list()
	case "del":
		delArgs := args[1:]
		ids := make([]int, len(delArgs))
		for i := range ids {
			id, err := strconv.Atoi(delArgs[i])
			if err != nil {
				log.Fatal(err)
			}
			ids[i] = id
		}
		tasks.del(ids)
	}
	tasks.save(dataPath)
}
