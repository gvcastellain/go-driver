package main

import (
	"os"
	"time"

	"github.com/gvcastellain/go-driver/internal/queue"
)

func main() {
	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	qc, err := queue.New(queue.RabbitMQ, qcfg)

	if err != nil {
		panic(err) //todo
	}
}
