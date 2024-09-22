package redis

func PushTask(queueName string, task Task) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.LPush(queueName, task).Result()
	return err
}

func PopTask(queueName string) (Task, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	task, err := client.BRPop(0, queueName).Result()
	return task, err
}
