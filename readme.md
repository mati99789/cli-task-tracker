# CLI Task Tracker

**Note: This project is built for learning purposes.**  
If you're a recruiter or someone reviewing this code, please be aware that this project was created to help practice Go programming skills, working with filesystems, handling user inputs, and building a basic CLI application. While functional, it is a simple project meant for educational growth.

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as "in-progress" or "done"
- List all tasks or filter tasks by their status (todo, in-progress, done)
- Store tasks in a JSON file, which is created if it doesn't exist
- Handle errors and edge cases gracefully

## Task Properties

Each task contains the following properties:

- `id`: A unique identifier for the task
- `description`: A short description of the task
- `status`: The current status of the task (can be `todo`, `in-progress`, or `done`)
- `createdAt`: The date and time when the task was created
- `updatedAt`: The date and time when the task was last updated

## Installation

1. Clone this repository:
    ```bash
    git clone <repository-url>
    cd cli-task-tracker
    ```

2. Ensure you're using Go version `1.22.2` or later:
    ```bash
    go version
    # Output should be go version go1.22.2 <platform>
    ```

3. Build the project:
    ```bash
    go build -o task-cli
    ```

## Usage

The task tracker accepts commands as arguments in the following format:

### Adding a new task
```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
