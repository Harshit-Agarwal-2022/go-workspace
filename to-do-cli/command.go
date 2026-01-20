package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Drop   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title and estimated time of completion. title:time_in_hours")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Drop, "drop", false, "deletes all the current elements in to do list")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {

	case cf.List:
		todos.print()

	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for add. Please use title:time_in_hours")
			os.Exit(1)
		}
		eT, err := strconv.ParseFloat(parts[1], 32)

		if err != nil {
			fmt.Println("Error , invalid format for add. Please use title:time_in_hours")
			os.Exit(1)
		}

		todos.add(parts[0], float32(eT))

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	case cf.Drop != false:
		todos.DeleteAll()

	default:
		fmt.Println("Invalid command")

	}

}
