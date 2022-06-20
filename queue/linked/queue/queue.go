package queue

import (
	"fmt"
	"os"
)

type Node struct {
	data    interface{}
	pointer *Node
}

type LinkedQueue struct {
	Head *Node
	Rear *Node
	Size int
}

func (lq *LinkedQueue) GetQueueSize() int {
	lq.Size = 0
	counter := lq.Head

	for {
		if counter != nil {
			lq.Size += 1
			counter = counter.pointer
		} else {
			break
		}
	}

	return lq.Size
}

func (lq *LinkedQueue) DeQueue() {
	if lq.Head == nil {
		fmt.Println("empty queue")
		os.Exit(1)
	}

	fmt.Println("dequeue:", lq.Head.data)

	if lq.Head.pointer != nil {
		lq.Head = lq.Head.pointer
	} else {
		lq.Head = nil
	}

}

func (lq *LinkedQueue) EnQueue(data interface{}) {
	node := &Node{
		data:    data,
		pointer: nil,
	}

	if lq.Rear != nil {
		lq.Rear.pointer = node
		lq.Rear = node
	} else {
		lq.Rear = node
	}

	// parse the interface{} by Println
	fmt.Println("enqueue:", lq.Rear.data)

	if lq.Head == nil {
		lq.Head = lq.Rear
	}
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		Head: nil,
		Rear: nil,
		Size: 0,
	}
}
