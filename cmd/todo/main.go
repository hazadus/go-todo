/*
CLI tool для работы со списком задач. Сохраняет список в файле в формате JSON.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/hazadus/go-todo"
)

// Default file name
// Override via GO_TODO_FILENAME env variable
var todoFileName = ".todo.json"

// getTask decides where to get the description for
// a new task from: arguments or stdin.
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", nil
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task description can't be empty")
	}
	return s.Text(), nil
}

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
	addFlag := flag.Bool("add", false, "Добавить задачу в список")
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
	case *addFlag:
		// flag.Args() returns all the remaining non-flag
		// arguments provided by the user:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		taskList.Add(task)

		// Save updated task list
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
