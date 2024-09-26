
```markdown
# GoQ - Distributed Task Queue System

GoQ is a distributed task queue system implemented in Go with support for Redis. It allows background job processing, task scheduling, and worker management, making it easy to process tasks asynchronously in your Go applications.

## Features
- Simple API for adding tasks to the queue
- Task scheduling with Redis sorted sets
- Workers that continuously pop and process tasks
- Retry mechanism for task processing
- Easily extendable to different task processing strategies

## Installation

First, make sure you have [Go](https://golang.org/dl/) installed, and Redis running on your machine.

```bash
go get github.com/alireza/GoQ
```

Next, import the package in your project:

```go
import "github.com/alireza/GoQ/pkg/task"
import "github.com/alireza/GoQ/pkg/workers"
```

## Usage

### 1. Define Your Task Structure

Tasks can have any structure, but a simple task contains an ID, name, payload, and retry count.

```go
type Task struct {
    ID      string
    Name    string
    Payload interface{}
    Retry   int
}
```

### 2. Adding a Task to the Queue

You can add a task to the queue using the `PushTask` function.

```go
newTask := task.Task{
    ID:    "12345",
    Name:  "ExampleTask",
    Retry: 3,
}

err := task.PushTask("task_queue", newTask)
if err != nil {
    log.Println("Error adding task:", err)
}
```

### 3. Running Workers

To process tasks, you need to start workers that continuously pop tasks from the queue and process them. The worker takes a queue name and a function to handle task processing.

```go
func processTask(t task.Task) error {
    log.Println("Processing task:", t.Name)
    // Task-specific processing logic goes here
    return nil
}

func main() {
    // Start a worker that processes tasks from the "task_queue"
    go workers.Worker("task_queue", processTask)

    // Keep the main function alive
    select {}
}
```

### 4. Scheduling Tasks

To schedule a task for future execution, use `ScheduleTask` to add it to a Redis sorted set.

```go
delay := time.Minute * 5
err := task.ScheduleTask(newTask, delay)
if err != nil {
    log.Println("Error scheduling task:", err)
}
```

### 5. Processing Scheduled Tasks

You can periodically check the scheduled tasks set and push tasks to the queue when their time comes. This can be done using a custom worker or cron job.

```go
func ProcessScheduledTasks() {
    for {
        // Logic to pop scheduled tasks and push them to the task queue
        // based on the task's scheduled time
        time.Sleep(time.Second * 10) // Example of checking every 10 seconds
    }
}
```

## Configuration

By default, Redis is configured to run on `localhost:6379`. If you need to change the Redis server address or add authentication, you can modify the Redis client initialization in `task.go`.

```go
client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    // Password: "", // Set password if required
    DB:       0,    // Use default DB
})
```

## Example Application

```go
package main

import (
    "GoQ/pkg/task"
    "GoQ/pkg/workers"
    "log"
    "time"
)

func main() {
    // Define task processing logic
    processTask := func(t task.Task) error {
        log.Println("Processing task:", t.Name)
        return nil
    }

    // Start workers to process tasks
    go workers.Worker("task_queue", processTask)

    // Add a task to the queue
    newTask := task.Task{
        ID:    "12345",
        Name:  "ExampleTask",
        Retry: 3,
    }

    err := task.PushTask("task_queue", newTask)
    if err != nil {
        log.Println("Error adding task:", err)
    }

    // Schedule a task for 1 minute later
    err = task.ScheduleTask(newTask, time.Minute*1)
    if err != nil {
        log.Println("Error scheduling task:", err)
    }

    // Keep the application running
    select {}
}
```

## Redis Dependency

GoQ requires a running instance of Redis. You can start a local instance of Redis using Docker for quick setup:

```bash
docker run --name some-redis -d redis
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contribution

Feel free to open issues and submit pull requests to improve GoQ! Contributions are welcome.

---

Happy coding! 🚀
```

### Key Points in the `README.md`:

1. **Installation**: Instructions to install the package using `go get`.
2. **Usage**: Step-by-step guide for adding tasks, running workers, and scheduling tasks.
3. **Example Application**: A simple example to demonstrate how to use the task queue in a real-world scenario.
4. **Redis Dependency**: Information on the requirement for Redis and how to quickly set it up.
5. **Configuration**: Brief information on how to configure the Redis client if needed.
6. **License and Contribution**: Licensing information and a call for contributions.

This document will serve as a guide for developers using your task queue package in their projects.