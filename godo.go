package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tenp/godo/task"
)

// usage contains the godo help text.
const usage string = `
godo [subcommand] [arguments to subcommand]

subcommands:
	add, a      Add one or more tasks
	del, d      Delete one or more tasks by their number
	list, l     Show a list of all tasks
	help, h     Show help text

examples:
	godo add "Buy car" Milk "Run away"      Add three tasks
	godo del 0 1                            Delete two tasks
	godo list                               Show all tasks
	godo help                               Show help
`

// saveFileName is the name that is used for the godo save file. Note that this
// solely describes the file name and not the path.
const saveFileName string = ".godo"

// tasks is the actual list of tasks used by godo.
var tasks *task.Tasklist

func main() {
	checkNumOfArgs()
	path := getSaveFilePath()
	tasks = &task.Tasklist{Title: "Tasks"}
	if err := tasks.Load(path); err != nil {
		log.Fatal(err)
	}
	processSubcommands()
	if err := tasks.Save(path); err != nil {
		log.Fatal(err)
	}
}

// processSubcommands evaluates the second cli argument.
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
	default:
		fmt.Println("Unknown subcommand")
		fmt.Print(usage)
	}
}

// checkNumOfArgs exits the program if no cli arguments are supplied.
func checkNumOfArgs() []string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal(usage)
	}
	return args
}

// getSaveFilePath either returns the value of the environment variable GODOPATH,
// or the HOME env var. The save file name is always determined by the saveFileName
// const.
func getSaveFilePath() string {
	path := os.Getenv("GODOPATH")
	if path == "" {
		path = os.Getenv("HOME")
	}
	return path + "/" + saveFileName
}

// add processes the add subcommand.
func add(args []string) {
	tasks.Add(args)
}

// list processes the list subcommand.
func list() {
	fmt.Print(tasks)
}

// del processes the del subcommand.
func del(args []string) {
	ids := make([]int, len(args))
	for i := range ids {
		id, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatal("Arguments to del subcommand must be of type integer\n" + usage)
		}
		ids[i] = id
	}
	tasks.Del(ids)
}
