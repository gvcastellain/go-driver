package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

func newRabbitConn(cfg RabbitMQConfig) (rc *RabbitConnection, err error) {
	rc.cfg = cfg

	rc.conn, err = amqp.Dial(cfg.URL)

	return rc, err
}

type RabbitConnection struct {
	cfg  RabbitMQConfig
	conn *amqp.Connection
}

func (rc *RabbitConnection) Publish(msg []byte) error {
	c, err := rc.conn.Channel()
	if err != nil {
		return err
	}

	mp := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         msg,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return c.PublishWithContext(ctx, "", rc.cfg.TopicName, false, false, mp)
}

func (rc *RabbitConnection) Consume(cdto chan<- QueueDto) error {
	ch, err := rc.conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(rc.cfg.TopicName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for d := range msgs {
		dto := QueueDto{}
		dto.Unmarshal(d.Body)

		cdto <- dto
	}

	return nil
}
