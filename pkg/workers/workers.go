package workers

import (
	"GoQ/pkg/task"
	"log"
	"time"
)

func Worker(queueName string, process func(task.Task) error) {
	for {
		t, err := task.PopTask(queueName)
		if err != nil {
			log.Println("Failed to pop task:", err)
			continue
		}

		err = process(t)
		if err != nil {
			log.Println("Task processing failed:", err)
		}
	}
}

func ScheduledWorker(queueName string) {
	for {
		// Check for any tasks that are ready to be moved from the scheduled set to the queue
		err := task.MoveScheduledTasks(queueName)
		if err != nil {
			log.Println("Error moving scheduled tasks:", err)
		}
		time.Sleep(10 * time.Second) // Sleep for a bit before checking again
	}
}
