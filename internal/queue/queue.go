package queue

import (
	"fmt"
	"log"
	"reflect"
)

const (
	RabbitMQ QueueType = iota
)

type QueueType int

type QueueConnection interface {
	Publish([]byte) error
	Consume(chan<- QueueDto) error
}

type Queue struct {
	cfg any
	qc  QueueConnection
}

func New(qt QueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("cfg need to be a RabbitMQConfig")
		}
		conn, err := newRabbitConn(cfg.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}

		q.qc = conn
	default:
		log.Fatal("type not implemented")
	}

	return
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume(cdto chan<- QueueDto) error {
	return q.qc.Consume(cdto)
}
