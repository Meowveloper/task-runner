package main

import (
	"flag"
	"fmt"
	"os"
	"task-runner/runner"
)

func main() {
	file_flag := flag.String("file", "tasks.json", "Path to tasks file (default: tasks.json)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go [--file=tasks.json] <task-name>")
		os.Exit(1)
	}
	task_name := args[0]

	tasks, err := runner.Load_Tasks(*file_flag)
	if err != nil {
		fmt.Printf("error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if err := runner.Run_Task(task_name, tasks); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

}
