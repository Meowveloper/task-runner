# 🛠️ Go Task Runner CLI

A simple task runner written in Go — similar to `make`, `npm run`, or `just`. It loads tasks from a JSON file and executes them via the terminal.

## 📦 Features

- Execute shell commands defined in `tasks.json`
- Support for multiple commands per task
- Custom task file support via `--file=...`
- Easy to extend and beginner-friendly Go code


## 🚀 Getting Started

### 1. Clone the project

```bash
git clone https://github.com/Meowveloper/task-runner.git
cd task-runner
go run main.go hello
go run main.go --file=custom_tasks.json build
```