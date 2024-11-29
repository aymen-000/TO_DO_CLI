package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Edit   string
	Del    int
	Toggle int
	List   bool
	Help   bool
}

func help() {
	fmt.Println("Todo CLI Application Usage:")
	fmt.Println("\nAvailable Commands:")
	fmt.Println("  -add       Add a new todo")
	fmt.Println("            Example: ./todocli -add \"Buy groceries\"")
	fmt.Println("\n  -edit      Edit an existing todo by its index")
	fmt.Println("            Example: ./todocli -edit \"New task title\" -index 2")
	fmt.Println("\n  -delete    Delete a todo by its index")
	fmt.Println("            Example: ./todocli -delete 3")
	fmt.Println("\n  -toggle    Mark a todo as completed or uncompleted")
	fmt.Println("            Example: ./todocli -toggle 1")
	fmt.Println("\n  -list      List all todos")
	fmt.Println("            Example: ./todocli -list")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nNote: Index starts from 1 for todo items.")
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.BoolVar(&cf.Help, "help", false, "Show help information")
	flag.StringVar(&cf.Add, "add", "", "Add a new todo title ...")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title ...")
	flag.IntVar(&cf.Del, "delete", -1, "Delete a specific todo by index ...")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo ...")
	flag.BoolVar(&cf.List, "list", false, "List all todos ...")
	flag.Parse()

	// Show help if -help flag is used
	if cf.Help {
		help()
		os.Exit(0)
	}

	return &cf
}

func (cf *cmdFlags) Execute(totdos *Todos, db *sql.DB) error {
	switch {
	case cf.List:
		totdos.print()

	case cf.Add != "":
		error := totdos.add(cf.Add, db)
		if error != nil {
			return error
		}
		return nil
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
		err = totdos.edit(index, parts[1], db)

		return err

	case cf.Toggle != -1:
		err := totdos.toggle(cf.Toggle, db)

		if err != nil {
			return err
		}

	case cf.Del != -1:
		err := totdos.delete(cf.Del, db)
		return err

	default:
		help()
	}

	return nil

}
