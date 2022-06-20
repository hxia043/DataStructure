package main

import (
	"fmt"
	"linked_queue/queue"
)

func main() {
	linkedQueue := queue.NewLinkedQueue()

	linkedQueue.EnQueue(10)
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.EnQueue("hxia")
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.EnQueue("Troy")
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.EnQueue(40)
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.DeQueue()
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.DeQueue()
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.DeQueue()
	fmt.Println("size:", linkedQueue.GetQueueSize())

	linkedQueue.DeQueue()
	fmt.Println("size:", linkedQueue.GetQueueSize())
}
