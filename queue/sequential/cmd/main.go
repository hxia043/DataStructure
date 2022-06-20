package main

import (
	"squeue/queue"
)

var queueLength int = 5

func main() {
	queue := queue.New(queueLength, queue.CircularQueueType)
	queue.InitQueue()

	queue.EnQueue(10)
	queue.EnQueue(20)
	queue.EnQueue(30)
	queue.EnQueue(40)
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
}
