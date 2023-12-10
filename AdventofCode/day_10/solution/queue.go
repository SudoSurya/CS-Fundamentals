
package main


// Queue represents a basic queue data structure.
type Queue struct {
	items [][]int
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(item []int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue) Dequeue() []int {
	if len(q.items) == 0 {
		return nil // Queue is empty
	}

	// Get the front item
	front := q.items[0]

	// Remove the front item by slicing the queue
	q.items = q.items[1:]

	return front
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}
