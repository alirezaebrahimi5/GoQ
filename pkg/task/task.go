package task

import (
	"GoQ/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var client *redis.Client

func InitRedis(cfg *config.Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}

type Task struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
	Retry   int         `json:"retry"`
}

func PushTask(queueName string, task Task) error {
	// Serialize the task to JSON
	data, err := json.Marshal(task)
	if err != nil {
		log.Println("Error marshaling task:", err)
		return err
	}
	_, err = client.LPush(queueName, data).Result()
	return err
}

func PopTask(queueName string) (Task, error) {
	result, err := client.BRPop(0, queueName).Result()
	if err != nil {
		return Task{}, err
	}

	var task Task
	// Deserialize the task from JSON
	err = json.Unmarshal([]byte(result[1]), &task)
	if err != nil {
		log.Println("Error unmarshaling task:", err)
		return Task{}, err
	}
	task.ID = result[1] // Set the ID to the popped task identifier (if needed)
	return task, nil
}

func ScheduleTask(queueName string, task Task, delay time.Duration) error {
	// Serialize the task to JSON
	data, err := json.Marshal(task)
	if err != nil {
		log.Println("Error marshaling task for scheduling:", err)
		return err
	}
	score := time.Now().Add(delay).Unix()
	err = client.ZAdd(queueName+"-scheduled", redis.Z{Score: float64(score), Member: data}).Err() // Append '-scheduled' to differentiate
	return err
}

func MoveScheduledTasks(queueName string) error {
	now := float64(time.Now().Unix())
	k := queueName + "-scheduled"
	log.Println(k)
	// Get tasks from the scheduled set that are ready to run
	tasks, err := client.ZRangeByScoreWithScores(k, redis.ZRangeBy{
		Min: "-inf",
		Max: fmt.Sprintf("%f", now),
	}).Result()
	if err != nil {
		log.Println("Error retrieving scheduled tasks:", err)
		return err
	}

	for _, taskWithScore := range tasks {
		// Deserialize the task from JSON
		var task Task
		err := json.Unmarshal([]byte(taskWithScore.Member.(string)), &task)
		if err != nil {
			log.Println("Error unmarshaling scheduled task:", err)
			continue
		}

		// Push the task to the main queue
		err = PushTask(queueName, task)
		if err != nil {
			log.Println("Error pushing scheduled task to queue:", err)
			continue
		}

		// Remove the task from the scheduled set after processing
		_, err = client.ZRem(queueName+"-scheduled", taskWithScore.Member).Result()
		if err != nil {
			log.Println("Error removing task from scheduled set:", err)
		}
	}

	return nil
}
