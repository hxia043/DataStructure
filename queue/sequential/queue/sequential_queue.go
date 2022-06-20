package queue

import (
	"fmt"
	"os"
)

type SequentialQueue struct {
	Head   int
	Tail   int
	Size   int
	Length int
	queue  []int
}

func (sq *SequentialQueue) GetQueueSize() int {
	return sq.Tail - sq.Head
}

func (sq *SequentialQueue) DeQueue() {
	if sq.Head == sq.Tail {
		fmt.Println("empty queue")
		os.Exit(3)
	}

	data := sq.queue[sq.Head]
	fmt.Println("dequeue:", data)

	sq.Head += 1
}

func (sq *SequentialQueue) EnQueue(data int) {
	if sq.Tail == (sq.Length - 1) {
		fmt.Println("full queue: fake overflow")
		os.Exit(2)
	}

	sq.queue[sq.Tail] = data
	fmt.Println("enqueue", sq.queue[sq.Tail])

	sq.Tail += 1
}

func (sq *SequentialQueue) InitQueue() {
	if sq.Length < 0 {
		fmt.Println("Error: invaliad queue length", sq.Length)
		os.Exit(1)
	}

	sq.queue = make([]int, sq.Length)
	sq.Head, sq.Tail = 0, 0
	sq.Size = 0
}

func NewSequentialQueue(len int) *SequentialQueue {
	return &SequentialQueue{
		Length: len,
	}
}
