package workers

import "log"

func Worker(queueName string) {
	for {
		task, err := PopTask(queueName)
		if err != nil {
			log.Println("Failed to pop task:", err)
			continue
		}

		// Process the task
		err = ProcessTask(task)
		if err != nil {
			log.Println("Task failed:", err)
			// Optionally requeue or retry task
		}
	}
}

func ProcessTask(task Task) error {
	// Define how to handle each task
	// Example: Call an API or process a file
	return nil
}
