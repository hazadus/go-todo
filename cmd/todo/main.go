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

// Default file name
// Override via GO_TODO_FILENAME env variable
var todoFileName = ".todo.json"

func main() {
	if envTodoFileName := os.Getenv("GO_TODO_FILENAME"); envTodoFileName != "" {
		todoFileName = envTodoFileName
	}
	
	// Define custom output for "-h" flag
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s CLI tool. Разработано для изучения языка Go.\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "(c) amgold.ru 2024\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Поддерживаемые параметры:")
		flag.PrintDefaults()
	}

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
		fmt.Print(taskList)
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