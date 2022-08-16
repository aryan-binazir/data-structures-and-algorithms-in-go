package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue adds a value at end
func (q *Queue) Enqueue(toAdd int) {
	q.items = append(q.items, toAdd)
}

// Dequeue removes value front of queue and returns it
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)
	myQueue.Enqueue(15)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)
}
