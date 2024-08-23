package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName      = "todo_test"
	todoFileName = ".todo_test.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	envTodoFileName := os.Getenv("GO_TODO_FILENAME")
	os.Setenv("GO_TODO_FILENAME", todoFileName)

	buildCmd := exec.Command("go", "build", "-o", binName)
	if err := buildCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't build binary %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(todoFileName)

	os.Setenv("GO_TODO_FILENAME", envTodoFileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "Test task 1"
	task2 := "Test task 2"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTaskFromStdin", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("[ ] 1: %s\n[ ] 2: %s\n", task, task2)
		if expected != string(output) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(output))
		}
	})
}
