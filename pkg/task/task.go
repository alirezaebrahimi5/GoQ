package task

import (
	"GoQ/pkg/config" // Import the config package
	"github.com/go-redis/redis"
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
	ID      string
	Name    string
	Payload interface{}
	Retry   int
}

func PushTask(queueName string, task Task) error {
	_, err := client.LPush(queueName, task).Result()
	return err
}

func PopTask(queueName string) (Task, error) {
	result, err := client.BRPop(0, queueName).Result()
	if err != nil {
		return Task{}, err
	}

	var task Task
	// Deserialize the task here (you might need to use json.Unmarshal based on how you store tasks)
	task.ID = result[1] // Adjust according to how you store tasks
	return task, nil
}

func ScheduleTask(queueName string, task Task, delay time.Duration) error {
	score := time.Now().Add(delay).Unix()
	err := client.ZAdd(queueName+"-scheduled", redis.Z{Score: float64(score), Member: task}).Err() // Append '-scheduled' to differentiate
	return err
}
