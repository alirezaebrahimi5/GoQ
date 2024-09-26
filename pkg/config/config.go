package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

func LoadConfig(filePath string) (*Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type TaskConfig struct {
	QueueName         string `yaml:"queue_name"`
	ScheduledTasksSet string `yaml:"scheduled_tasks_set"`
}

type Config struct {
	Redis RedisConfig  `yaml:"redis"`
	Tasks []TaskConfig `yaml:"tasks"` // Change to a slice of TaskConfig
}
