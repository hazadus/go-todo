/*
CLI tool для работы со списком задач. Сохраняет список в файле в формате JSON.
*/
package main

import (
	"fmt"
	"os"
	"strings"

	todo "github.com/hazadus/go-todo"
)

const todoFileName = ".todo.json"

func main() {
	list := &todo.List{}

	if err := list.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decided what to do depending on the number of
	// arguments provided
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		for _, item := range *list {
			fmt.Println(item.Task)
		}
	// Concatenate all provided arguments with a space
	// and add to the list as an item
	default:
		item := strings.Join(os.Args[1:], " ")
		list.Add(item)

		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}