package queue

import (
	"fmt"
)

const (
	SequentialQueueType string = "sequential"
	CircularQueueType   string = "circular"
)

type Queue interface {
	InitQueue()
	EnQueue(data int)
	DeQueue()
	GetQueueSize() int
}

func New(len int, queueType string) Queue {
	switch queueType {
	case "sequential":
		return NewSequentialQueue(len)
	case "circular":
		return NewCircularQueue(len)
	default:
		fmt.Println("unknown queue type")
		return nil
	}
}
