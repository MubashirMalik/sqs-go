package queue

import "fmt"

type SQSParams struct {
	visibilityTimeout int // 0sec-12hr
}

type Node struct {
	message string
	next *Node
}

type Queue struct {
	sqlParams SQSParams
	front *Node // pointer to the first node
	rear  *Node // pointer to the last node
}

func CreateQueue() Queue {
	queue := Queue{}
	return queue
}

func (queue *Queue) IsEmpty() bool {
	return queue.front == nil && queue.rear == nil
}

func (queue *Queue) Enqueue(message string) {
	newNode := &Node{ message: message }

	if queue.IsEmpty() {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}
}

func (queue *Queue) Dequeue() (string, bool) {
	if queue.IsEmpty() {
		return "", false
	} else {
		data := queue.front.message
		queue.front = queue.front.next
		if queue.front == nil {
			queue.rear = nil
		}
		return data, true
	}
}

func (queue *Queue) Peek() (string, bool) {
	if queue.IsEmpty() {
		return "", false
	} else {
		return queue.front.message, true
	}
}

func (queue *Queue) Print() {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty.")
		return
	} else {
		fmt.Print("Printing queue: ")
	}

	for temp := queue.front; temp != nil; temp = temp.next {
		fmt.Printf("%v ", temp.message)
	}
	fmt.Println()
}