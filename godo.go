package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var dataPath string = ".godo" // Use system variable instead

func load() *[]string {
	if _, err := os.Stat(dataPath); err != nil {
		if _, err := os.Create(dataPath); err != nil {
			log.Fatal(err, "\nCould not create file at "+dataPath)
		}
	}
	raw, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatal(err, "\nCould not read data from "+dataPath)
	}
	text := string(raw)
	lines := strings.Split(text, "\n")
	return &lines
}

func add(args []string) {
	for _, task := range args {
		fmt.Println(task)
	}
}

func list(args []string) {
	fmt.Println("list", args)
}

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
	data := load()
	fmt.Println(*data)
	switch args[0] {
	case "add":
		add(args[1:])
	case "list":
		list(args[1:])
	}
}
