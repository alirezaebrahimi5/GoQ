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

First, make sure you have [Go](https://golang.org/dl/) installed and Redis running on your machine.

```bash
go get github.com/alirezaebrahimi5/GoQ
```

Next, import the package in your project:

```go
import "github.com/alirezaebrahimi5/GoQ/pkg/task"
import "github.com/alirezaebrahimi5/GoQ/pkg/workers"
```

## Configuration

Create a `config.yaml` file in the root directory of the project to configure the Redis connection and define your task queues:

```yaml
# config.yaml

redis:
  addr: "localhost:6379"   # Redis server address
  password: ""             # Redis password (leave empty if not used)
  db: 0                    # Redis database number

tasks:
  - queue_name: "task_queue_1" 
    scheduled_tasks_set: "scheduled_tasks_1"
  - queue_name: "task_queue_2" 
    scheduled_tasks_set: "scheduled_tasks_2"
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

err := task.PushTask("task_queue_1", newTask) // Use the appropriate queue name
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
    // Start a worker that processes tasks from "task_queue_1"
    go workers.Worker("task_queue_1", processTask)

    // Start a worker that processes tasks from "task_queue_2"
    go workers.Worker("task_queue_2", processTask)

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
    go workers.Worker("task_queue_1", processTask)
    go workers.Worker("task_queue_2", processTask)

    // Add a task to the first queue
    newTask := task.Task{
        ID:    "12345",
        Name:  "ExampleTask",
        Retry: 3,
    }

    err := task.PushTask("task_queue_1", newTask)
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

### Key Updates

1. **Configuration Section**: Added a new section to specify how to set up the `config.yaml` file for multiple task queues and Redis connection settings.
2. **Usage Examples**: Updated examples to show how to push tasks to different queues based on the configuration.
3. **Worker Management**: Included how to start multiple workers for processing tasks from different queues.

This updated document should guide users on how to utilize your task queue system effectively while configuring multiple task queues. Let me know if you need further adjustments!
