package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <task-name>")
		os.Exit(1)
	}
	task_name := os.Args[1]

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Printf("Error reading tasks.json: %v\n", err)
		os.Exit(1)
	}

	tasks := make(map[string]string)

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Printf("Error parsing json file %v\n", err)
	}

	command, ok := tasks[task_name]
	if !ok {
		fmt.Printf("Task '%s' not found.\n", task_name)
		os.Exit(1)
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running task: %s -> %s\n\n", task_name, command)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running task: %v\n", err)
		os.Exit(1)
	}
}
