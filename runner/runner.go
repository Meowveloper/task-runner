package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func Load_Tasks(file_path string) (map[string]string, error) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}

	tasks := make(map[string]string)

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func Run_Task(task_name string, tasks map[string]string) error {
	command, ok := tasks[task_name]
	if !ok {
		return fmt.Errorf("task '%s' not found", task_name)
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running command %s -> %s\n\n", command, task_name)

	return cmd.Run()
}
