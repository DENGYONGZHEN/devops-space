package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/devops-space/powerfulCommandLineApplicationInGO/2_interactWithUser"
)

var todoFileName = ".todo.json"

func main() {

	//flag.Usage: this varaible is pointing to a function,
	// change it will display custome message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for deng\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright forever\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	listFlag := flag.Bool("list", false, "List all tasks")
	// taskFlag := flag.String("task", "", "Task to be included in the ToDO list")
	addFalg := flag.Bool("add", false, "Add task to the ToDo list")
	completeFlag := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	//Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listFlag:
		//List类型实现了Stringer接口，拥有自定义的String方法，fmt.Print(l)会调用对象的String方法
		fmt.Print(l)
		// for _, task := range *l {
		// 	if !task.Done {
		// 		fmt.Println(task.Task)
		// 	}
		// }
	case *completeFlag > 0:
		if err := l.Complete(*completeFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// case *taskFlag != "":
	case *addFalg:
		// When any arguments (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}

// getTask function decides where to get the description for a new
// task from: arguments of STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}

//exercise
//1  Implement the flag -del to delete an item from the list. Use the Delete()
//method from the API to perform the action.
//2 Add another flag to enable verbose output, showing information like
//date/time.
//3 Add another flag to prevent displaying completed items.
//4 Update the custom usage function to include additional instructions on
//how to provide new tasks to the tool.
//5 Include test cases for the remaining options, such as -complete.
//6 Update the tests to use the TODO_FILENAME environment variable instead of
//hard-coding the test file name so that it doesn’t cause conflicts with an
// existing file.
//7 Update the getTask() function allowing it to handle multiline input from
// STDIN. Each line should be a new task in the list.
