package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func Load_Tasks(file_path string) (map[string][]string, error) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}

	raw := make(map[string]any)

	tasks := make(map[string][]string)

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	for name, val := range raw {
		switch v := val.(type) {
		case string:
			tasks[name] = []string{v}
		case []any:
			for _, item := range v {
				if cmd, ok := item.(string); ok {
					tasks[name] = append(tasks[name], cmd)
				}
			}
		default:
			return nil, fmt.Errorf("invalid format for task '%s'", name)
		}
	}

	return tasks, nil
}

func Run_Task(task_name string, tasks map[string][]string) error {
	commands, ok := tasks[task_name]
	if !ok {
		return fmt.Errorf("task '%s' not found", task_name)
	}
	for i, command := range commands {
		fmt.Printf("(%d/%d) running %s\n", i+1, len(commands), command)
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("command failed %s\n%v", command, err)
		}
	}
	return nil

}
