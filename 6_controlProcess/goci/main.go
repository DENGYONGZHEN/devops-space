package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type executer interface {
	execute() (string, error)
}

func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}
	// args := []string{"build", ".", "errors"}
	// cmd := exec.Command("go", args...)
	// cmd.Dir = proj
	// if err := cmd.Run(); err != nil {
	// 	return &stepErr{step: "go build", msg: "go build failed", cause: err}
	// }
	// _, err := fmt.Fprintln(out, "GO Build: SUCCESS")

	// pipeline := make([]step, 2)
	pipeline := make([]executer, 4)
	pipeline[0] = newStep("go build", "go", "GO Build: SUCCESS", proj, []string{"build", ".", "errors"})
	pipeline[1] = newStep("go test", "go", "GO Test: SUCCESS", proj, []string{"test", "-v"})
	pipeline[2] = newExceptionStep("go fmt", "gofmt", "Gofmt: SUCCESS", proj, []string{"-l", "."})
	pipeline[3] = newTimeoutStep("git push", "git", "Git Push: SUCCESS", proj, []string{"push", "origin", "main"}, 10*time.Second)

	sig := make(chan os.Signal, 1)
	errCh := make(chan error)
	done := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for _, s := range pipeline {
			msg, err := s.execute()
			if err != nil {
				errCh <- err
				return
			}
			_, err = fmt.Fprintln(out, msg)
			if err != nil {
				errCh <- err
				return
			}
		}
		close(done)
	}()

	for {
		select {
		case rec := <-sig:
			signal.Stop(sig)
			return fmt.Errorf("%s: existing: %w", rec, ErrSignal)
		case err := <-errCh:
			return err
		case <-done:
			return nil
		}
	}
}


//1 Add another step to the pipeline: code linting using golangci-lint. For more
// information consult its home page.8
//2 Add gocyclo to the pipeline. Capture its output and return an error if gocyclo
// returns any functions with a complexity score of 10 or greater. For more
// information about this tool, consult its GitHub page.9
//3 Add environment variables to handle Git authentication with remote
// repositories that require it.
//4 Add another command-line flag to your tool asking for the Git branch to
// push. Update the Git step to accept a configurable branch instead of
// master.
//5 Get the Pipeline configuration from a file instead of hard-coding it in the
// run() function.