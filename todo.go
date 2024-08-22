/*
Package todo содержит бизнес-логику для работы с to-do.
*/
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents a to-do item
type item struct {
	Task        string
	IsDone      bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of to-do items
type List []item

// Add creates a new to-do item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task: task,
		IsDone: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks to-do item as completed
// by setting `Done = true` and `CompletedAt` to the
// current time.
func (l *List) Complete(i int) error {
	list := *l

	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d does not exist", i)
	}

	list[i-1].IsDone = true
	list[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a to-do item from the list
func (l *List) Delete(i int) error {
	list := *l

	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(list[:i-1], list[i:]... )

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	jsonList, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonList, 0644)
}

// Load method opens the provided file name, decodes
// the JSON data and parses it into a List. Returns 
// nil if file does not exist or is empty.
func (l *List) Load(filename string) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(fileContent) == 0 {
		return nil
	}

	return json.Unmarshal(fileContent, l)
}

// String returns a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	output := ""

	for index, item := range *l {
		prefix := "[ ] "
		if item.IsDone {
			prefix = "[X] "
		}
		output += fmt.Sprintf("%s%d: %s\n", prefix, index+1, item.Task)
	}

	return output
}