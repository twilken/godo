/*
Yet another simple command line todo list.

Usage

	godo [subcommand] [arguments to subcommand]

The following subcommands are available:
	add, a	   Add one or more tasks
	del, d	   Delete one or more tasks by their number
	list, l	   Show a list of all tasks
	help, h	   Show help text

There is a short version for every subcommand.

Examples

Show all tasks:
	godo list

Add three tasks:
	godo add "Get milk" Workout "Dump TV"

Delete tasks at position 1 and 2:
	godo del 1 2

Print help information:
	godo help

Save file location

By default the godo save file is in your `HOME` directory. You can change that by creating a
`GODOPATH` environment variable and setting it to another path. Note that, for now, the directory must
exist before godo can use it. The actual file is always named `.godo`.
*/
package main
