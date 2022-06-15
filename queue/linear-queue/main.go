package main

import "linear-queue/linearqueue"

func main() {
	q := linearqueue.NewQueue(6)
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(8)
	q.Enqueue(6)
	q.Enqueue(4)
	q.Enqueue(7)
	q.Display()

	q.Dequeue()
	q.Display()
}
