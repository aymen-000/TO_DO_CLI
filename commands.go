package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo title ...")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title ...")
	flag.IntVar(&cf.Del, "delete", -1, "Delete a specific todo by index ...")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo ...")
	flag.BoolVar(&cf.List, "list", false, "List all todos ...")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(totdos *Todos) {
	switch {
	case cf.List:
		totdos.print()

	case cf.Add != "":
		totdos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use id:title format.")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index")
			os.Exit(1)
		}
		totdos.edit(index, parts[1])

	case cf.Toggle != -1:
		totdos.toggle(cf.Toggle)

	case cf.Del != -1:
		totdos.delete(cf.Del)

	default:
		fmt.Println("Invalid command.")
	}
}
