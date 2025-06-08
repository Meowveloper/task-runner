package main

import (
	"fmt"
	"os"
	"task-runner/runner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <task-name>")
		os.Exit(1)
	}
	task_name := os.Args[1]

	tasks, err := runner.Load_Tasks("tasks.json")
	if err != nil {
		fmt.Printf("error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if err := runner.Run_Task(task_name, tasks); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

}
