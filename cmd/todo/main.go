/*
CLI tool для работы со списком задач. Сохраняет список в файле в формате JSON.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/hazadus/go-todo"
)

const todoFileName = ".todo.json"

func main() {
	// Parse command line flags
	taskFlag := flag.String("task", "", "Задача для добавления в список")
	listFlag := flag.Bool("list", false, "Вывести список задач")
	completeFlag := flag.Int("complete", 0, "Завершить задачу с указанным номером")
	flag.Parse()

	// Define and load task taskList
	taskList := &todo.List{}

	if err := taskList.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decided what to do depending on the number of
	// arguments provided
	switch {
	case *listFlag:
		for _, item := range *taskList {
			if !item.IsDone {
				fmt.Println(item.Task)
			}
		}
	case *completeFlag > 0:
		if err := taskList.Complete(*completeFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := taskList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *taskFlag != "":
		taskList.Add(*taskFlag)
		if err := taskList.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flags, or none at all
		fmt.Fprintln(os.Stderr, "Неверные параметры")
		os.Exit(1)
}
}