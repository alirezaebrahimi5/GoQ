package workers

import (
	"GoQ/pkg/task"
	"log"
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
