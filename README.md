# Go Task Queue System

A distributed task queue system in Go with Redis/RabbitMQ support. Perfect for background job processing and task scheduling.

## Features
- Task queues using Redis/RabbitMQ
- Worker management for distributed task execution
- Task retries and failure handling
- HTTP API and CLI for task management

## Installation
```bash
go get github.com/alirezaebrahimi5/GoQ
```
## Usage Example
```bash
task := Task{Type: "email", Payload: "Send welcome email"}
PushTask(queueName, task)
Worker(queueName)
```
## Roadmap

- [ ] Support for RabbitMQ
- [ ] Task scheduling with cron jobs
