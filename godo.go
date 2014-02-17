package main

import (
	"flag"
	"fmt"
	"os"
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
	}
	tasks.save(dataPath)
}
