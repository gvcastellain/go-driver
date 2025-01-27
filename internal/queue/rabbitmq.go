package queue

import (
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitConnection struct {
	conn *ampq.Connection
}

func (rc *RabbitConnection) Publish(msg []byte) error {
	return nil
}

func (rc *RabbitConnection) Consume() error {
	return nil
}
