package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents a Todo item
type item struct {
	Task        string
	Done        bool
	CreateAt    time.Time
	CompletedAt time.Time
}

type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreateAt:    time.Now(),
		CompletedAt: time.Time{},
	}
	//Note that you need to dereference the pointer to the List type with *l
	//in the append call to access the underlying slice
	//append的第一个参数必须是slice，返回值是一个新切片，赋值给指针类型不匹配
	*l = append(*l, t)

}

// Complete method marks a Todo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exit", i)
	}
	//Adjusting index for 0 based indexing
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a Todo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Itme %d does not exit", i)
	}

	//Adjusting index for 0 based index

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

//Save method encodes the List as JSON and saves it
//using the provided file name

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name,decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		//Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
