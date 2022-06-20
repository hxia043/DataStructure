package queue

import (
	"fmt"
	"os"
)

type CircularQueue struct {
	Head   int
	Tail   int
	Size   int
	Length int
	queue  []int
}

func (cq *CircularQueue) GetQueueSize() int {
	return (cq.Tail - cq.Head + cq.Length) % cq.Length
}

func (cq *CircularQueue) DeQueue() {
	if cq.Head == cq.Tail {
		fmt.Println("empty queue")
		os.Exit(3)
	}

	data := cq.queue[cq.Head]
	fmt.Println(data)

	if cq.Head == (cq.Length - 1) {
		cq.Head = 0
	} else {
		cq.Head += 1
	}
}

func (cq *CircularQueue) EnQueue(data int) {
	if ((cq.Tail-cq.Head+cq.Length)%cq.Length)+1 == cq.Length {
		fmt.Println("full queue: true overflow")
		os.Exit(2)
	}

	cq.queue[cq.Tail] = data
	fmt.Println("enqueue:", cq.queue[cq.Tail])

	if cq.Tail == (cq.Length - 1) {
		cq.Tail = 0
	} else {
		cq.Tail += 1
	}
}

func (cq *CircularQueue) InitQueue() {
	if cq.Length < 0 {
		fmt.Println("Error: invaliad queue length", cq.Length)
		os.Exit(1)
	}

	cq.queue = make([]int, cq.Length)
	cq.Head, cq.Tail = 0, 0
	cq.Size = 0
}

func NewCircularQueue(len int) *CircularQueue {
	return &CircularQueue{
		Length: len,
	}
}
