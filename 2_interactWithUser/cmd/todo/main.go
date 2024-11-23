package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/devops-space/powerfulCommandLineApplicationInGO/2_interactWithUser"
)

const todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for deng\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright forever\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	listFlag := flag.Bool("list", false, "List all tasks")
	taskFlag := flag.String("task", "", "Task to be included in the ToDO list")
	completeFlag := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listFlag:
		for _, task := range *l {
			if !task.Done {
				fmt.Println(task.Task)
			}
		}
	case *completeFlag > 0:
		if err := l.Complete(*completeFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *taskFlag != "":
		l.Add(*taskFlag)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	// switch {
	// case len(os.Args) == 1:
	// 	for _, item := range *l {
	// 		fmt.Println(item.Task)
	// 	}
	// // Concatenate all provided arguments with a space and
	// // add to the list as an item
	// default:
	// 	item := strings.Join(os.Args[1:], " ")
	// 	l.Add(item)

	// 	if err := l.Save(todoFileName); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// }
}
