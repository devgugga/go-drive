package queue

import (
	"fmt"
	"log"
	"reflect"
)

type QueueType int

const (
	RabbitMQ QueueType = iota
)

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
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
			return nil, fmt.Errorf("config need's to be of type RabbitMQConfig")
		}
		fmt.Println("Não implementado!")
	default:
		log.Fatal("type not implemented")
	}

	return
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume() error {
	return q.qc.Consume()
}
