package main

import (
	"GoQ/pkg/config"
	"GoQ/pkg/task"
	"GoQ/pkg/workers"
	"log"
	"time"
)

func processTask(t task.Task) error {
	log.Println("Processing task:", t.Name)
	return nil
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize Redis client
	task.InitRedis(cfg)

	// Start workers for each queue defined in the configuration
	for _, taskConfig := range cfg.Tasks {
		go workers.Worker(taskConfig.QueueName, processTask)
	}

	// Example of adding tasks to the queues
	newTask1 := task.Task{
		ID:    "12345",
		Name:  "ExampleTask1",
		Retry: 3,
	}

	err = task.PushTask(cfg.Tasks[0].QueueName, newTask1) // Push to the first queue
	if err != nil {
		log.Println("Error adding task to queue 1:", err)
	}

	newTask2 := task.Task{
		ID:    "67890",
		Name:  "ExampleTask2",
		Retry: 3,
	}

	err = task.PushTask(cfg.Tasks[1].QueueName, newTask2) // Push to the second queue
	if err != nil {
		log.Println("Error adding task to queue 2:", err)
	}

	// Schedule a task for 1 minute later
	err = task.ScheduleTask(cfg.Tasks[0].QueueName, newTask1, time.Minute*1) // Schedule for the first queue
	if err != nil {
		log.Println("Error scheduling task for queue 1:", err)
	}

	// Keep the application running
	select {}
}
